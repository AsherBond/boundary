// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

package ferry

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/boundary/api"
	"github.com/hashicorp/boundary/internal/cmd/base"
	"github.com/hashicorp/boundary/internal/errors"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/mitchellh/cli"
	"github.com/posener/complete"
)

var (
	_ cli.Command             = (*StatusCommand)(nil)
	_ cli.CommandAutocomplete = (*StatusCommand)(nil)
)

type StatusCommand struct {
	*base.Command
}

func (c *StatusCommand) Synopsis() string {
	return "Get the status information of the running boundary ferry daemon"
}

func (c *StatusCommand) Help() string {
	helpText := `
Usage: boundary ferry status [options]

  Get the status of the boundary ferry daemon:

      $ boundary ferry status

  For a full list of examples, please see the documentation.

` + c.Flags().Help()
	return strings.TrimSpace(helpText)
}

func (c *StatusCommand) Flags() *base.FlagSets {
	set := c.FlagSet(base.FlagSetOutputFormat)
	f := set.NewFlagSet("Client Options")

	f.BoolVar(&base.BoolVar{
		Name:   "output-curl-string",
		Target: &c.FlagOutputCurlString,
		Usage:  "Instead of executing the request, print an equivalent cURL command string and exit.",
	})

	f.UintVar(&base.UintVar{
		Name:    "ferry-port",
		Target:  &c.FlagFerryDaemonPort,
		Default: 9300,
		EnvVar:  base.EnvFerryDaemonPort,
		Usage:   "The port on which the ferry daemon is listening.",
	})

	return set
}

func (c *StatusCommand) AutocompleteArgs() complete.Predictor {
	return complete.PredictNothing
}

func (c *StatusCommand) AutocompleteFlags() complete.Flags {
	return c.Flags().Completions()
}

func (c *StatusCommand) Run(args []string) int {
	ctx := c.Context
	f := c.Flags()
	if err := f.Parse(args); err != nil {
		c.PrintCliError(err)
		return base.CommandUserError
	}

	resp, result, apiErr, err := c.Status(ctx)
	if err != nil {
		c.PrintCliError(err)
		return base.CommandCliError
	}
	if apiErr != nil {
		c.PrintApiError(apiErr, "Error from ferry daemon when getting the status")
		return base.CommandApiError
	}

	switch base.Format(c.UI) {
	case "json":
		if ok := c.PrintJsonItem(resp); !ok {
			return base.CommandCliError
		}
	default:
		c.UI.Output(printStatusTable(result))
	}
	return base.CommandSuccess
}

type GetStatusResponse struct {
	BoundaryAddr    string    `json:"boundary_addr"`
	AuthTokenId     string    `json:"auth_token_id"`
	AuthTokenExpiry time.Time `json:"auth_token_expiry"`
	Version         string    `json:"version"`
	Errors          []string  `json:"errors"`
	Warnings        []string  `json:"warnings"`
}

func (c *StatusCommand) Status(ctx context.Context) (*api.Response, *GetStatusResponse, *api.Error, error) {
	const op = "ferry.(StatusCommand).Status"
	client := retryablehttp.NewClient()
	client.Logger = nil
	client.RetryWaitMin = 100 * time.Millisecond
	client.RetryWaitMax = 1500 * time.Millisecond

	req, err := retryablehttp.NewRequestWithContext(ctx, "GET", ferryUrl(c.FlagFerryDaemonPort, "v1/status"), nil)
	if err != nil {
		return nil, nil, nil, err
	}
	req.Header.Set("content-type", "application/json")

	if c.FlagOutputCurlString {
		api.LastOutputStringError = &api.OutputStringError{Request: req}
		return nil, nil, nil, api.LastOutputStringError
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, nil, errors.Wrap(ctx, err, op, errors.WithMsg("client do failed"))
	}
	apiResp := api.NewResponse(resp)

	res := &GetStatusResponse{}
	apiErr, err := apiResp.Decode(&res)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("Error when sending request to the ferry daemon: %w.", err)
	}
	if apiErr != nil {
		return apiResp, nil, apiErr, nil
	}
	return apiResp, res, nil, nil
}

func printStatusTable(status *GetStatusResponse) string {
	nonAttributeMap := map[string]any{
		"Address":               status.BoundaryAddr,
		"Auth Token Id":         status.AuthTokenId,
		"Auth Token Expiration": time.Until(status.AuthTokenExpiry).Round(time.Second).String(),
		"Version":               status.Version,
	}

	maxLength := base.MaxAttributesLength(nonAttributeMap, nil, nil)

	ret := []string{
		"",
		"Status:",
		base.WrapMap(2, maxLength+2, nonAttributeMap),
	}

	if len(status.Errors) > 0 {
		ret = append(ret, "  Recent errors:")
		ret = append(ret, base.WrapSlice(4, status.Errors))
	}

	if len(status.Warnings) > 0 {
		ret = append(ret, "  Recent warnings:")
		ret = append(ret, base.WrapSlice(4, status.Warnings))
	}
	return base.WrapForHelpText(ret)
}
