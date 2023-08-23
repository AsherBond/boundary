// Code generated by "make cli"; DO NOT EDIT.
// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package targetscmd

import (
	"errors"
	"fmt"
	"os"
	"sync"

	"github.com/hashicorp/boundary/api"
	"github.com/hashicorp/boundary/api/targets"
	"github.com/hashicorp/boundary/internal/cmd/base"
	"github.com/hashicorp/boundary/internal/cmd/common"
	"github.com/hashicorp/go-secure-stdlib/strutil"
	"github.com/mitchellh/cli"
	"github.com/posener/complete"
)

func initFlags() {
	flagsOnce.Do(func() {
		extraFlags := extraActionsFlagsMapFunc()
		for k, v := range extraFlags {
			flagsMap[k] = append(flagsMap[k], v...)
		}
	})
}

var (
	_ cli.Command             = (*Command)(nil)
	_ cli.CommandAutocomplete = (*Command)(nil)
)

type Command struct {
	*base.Command

	Func string

	plural string

	extraCmdVars
}

func (c *Command) AutocompleteArgs() complete.Predictor {
	initFlags()
	return complete.PredictAnything
}

func (c *Command) AutocompleteFlags() complete.Flags {
	initFlags()
	return c.Flags().Completions()
}

func (c *Command) Synopsis() string {
	if extra := extraSynopsisFunc(c); extra != "" {
		return extra
	}

	synopsisStr := "target"

	return common.SynopsisFunc(c.Func, synopsisStr)
}

func (c *Command) Help() string {
	initFlags()

	var helpStr string
	helpMap := common.HelpMap("target")

	switch c.Func {

	case "read":
		helpStr = helpMap[c.Func]() + c.Flags().Help()

	case "delete":
		helpStr = helpMap[c.Func]() + c.Flags().Help()

	case "list":
		helpStr = helpMap[c.Func]() + c.Flags().Help()

	default:

		helpStr = c.extraHelpFunc(helpMap)

	}

	// Keep linter from complaining if we don't actually generate code using it
	_ = helpMap
	return helpStr
}

var flagsMap = map[string][]string{

	"read": {"id"},

	"delete": {"id"},

	"list": {"scope-id", "filter", "recursive"},
}

func (c *Command) Flags() *base.FlagSets {
	if len(flagsMap[c.Func]) == 0 {
		return c.FlagSet(base.FlagSetNone)
	}

	set := c.FlagSet(base.FlagSetHTTP | base.FlagSetClient | base.FlagSetOutputFormat)
	f := set.NewFlagSet("Command Options")
	common.PopulateCommonFlags(c.Command, f, "target", flagsMap, c.Func)

	extraFlagsFunc(c, set, f)

	return set
}

func (c *Command) Run(args []string) int {
	initFlags()

	if os.Getenv("BOUNDARY_EXAMPLE_CLI_OUTPUT") != "" {
		c.UI.Output(exampleOutput())
		return base.CommandSuccess
	}

	switch c.Func {
	case "":
		return cli.RunResultHelp

	case "create":
		return cli.RunResultHelp

	case "update":
		return cli.RunResultHelp

	}

	c.plural = "target"
	switch c.Func {
	case "list":
		c.plural = "targets"
	}

	f := c.Flags()

	if err := f.Parse(args); err != nil {
		c.PrintCliError(err)
		return base.CommandUserError
	}

	var opts []targets.Option

	if strutil.StrListContains(flagsMap[c.Func], "scope-id") {
		switch c.Func {

		case "list":
			if c.FlagScopeId == "" {
				c.PrintCliError(errors.New("Scope ID must be passed in via -scope-id or BOUNDARY_SCOPE_ID"))
				return base.CommandUserError
			}

		}
	}

	client, err := c.Client()
	if c.WrapperCleanupFunc != nil {
		defer func() {
			if err := c.WrapperCleanupFunc(); err != nil {
				c.PrintCliError(fmt.Errorf("Error cleaning kms wrapper: %w", err))
			}
		}()
	}
	if err != nil {
		c.PrintCliError(fmt.Errorf("Error creating API client: %w", err))
		return base.CommandCliError
	}
	targetsClient := targets.NewClient(client)

	switch c.FlagName {
	case "":
	case "null":
		opts = append(opts, targets.DefaultName())
	default:
		opts = append(opts, targets.WithName(c.FlagName))
	}

	switch c.FlagDescription {
	case "":
	case "null":
		opts = append(opts, targets.DefaultDescription())
	default:
		opts = append(opts, targets.WithDescription(c.FlagDescription))
	}

	switch c.FlagRecursive {
	case true:
		opts = append(opts, targets.WithRecursive(true))
	}

	if c.FlagFilter != "" {
		opts = append(opts, targets.WithFilter(c.FlagFilter))
	}

	var version uint32

	switch c.Func {

	case "add-host-sources":
		switch c.FlagVersion {
		case 0:
			opts = append(opts, targets.WithAutomaticVersioning(true))
		default:
			version = uint32(c.FlagVersion)
		}

	case "remove-host-sources":
		switch c.FlagVersion {
		case 0:
			opts = append(opts, targets.WithAutomaticVersioning(true))
		default:
			version = uint32(c.FlagVersion)
		}

	case "set-host-sources":
		switch c.FlagVersion {
		case 0:
			opts = append(opts, targets.WithAutomaticVersioning(true))
		default:
			version = uint32(c.FlagVersion)
		}

	case "add-credential-sources":
		switch c.FlagVersion {
		case 0:
			opts = append(opts, targets.WithAutomaticVersioning(true))
		default:
			version = uint32(c.FlagVersion)
		}

	case "remove-credential-sources":
		switch c.FlagVersion {
		case 0:
			opts = append(opts, targets.WithAutomaticVersioning(true))
		default:
			version = uint32(c.FlagVersion)
		}

	case "set-credential-sources":
		switch c.FlagVersion {
		case 0:
			opts = append(opts, targets.WithAutomaticVersioning(true))
		default:
			version = uint32(c.FlagVersion)
		}

	}

	if ok := extraFlagsHandlingFunc(c, f, &opts); !ok {
		return base.CommandUserError
	}

	var resp *api.Response
	var item *targets.Target

	var items []*targets.Target

	var readResult *targets.TargetReadResult

	var deleteResult *targets.TargetDeleteResult

	var listResult *targets.TargetListResult

	switch c.Func {

	case "read":
		readResult, err = targetsClient.Read(c.Context, c.FlagId, opts...)
		if exitCode := c.checkFuncError(err); exitCode > 0 {
			return exitCode
		}
		resp = readResult.GetResponse()
		item = readResult.GetItem()

	case "delete":
		deleteResult, err = targetsClient.Delete(c.Context, c.FlagId, opts...)
		if exitCode := c.checkFuncError(err); exitCode > 0 {
			return exitCode
		}
		resp = deleteResult.GetResponse()

	case "list":
		listResult, err = targetsClient.List(c.Context, c.FlagScopeId, opts...)
		if exitCode := c.checkFuncError(err); exitCode > 0 {
			return exitCode
		}
		resp = listResult.GetResponse()
		items = listResult.GetItems()

	}

	resp, item, items, err = executeExtraActions(c, resp, item, items, err, targetsClient, version, opts)
	if exitCode := c.checkFuncError(err); exitCode > 0 {
		return exitCode
	}

	output, err := printCustomActionOutput(c)
	if err != nil {
		c.PrintCliError(err)
		return base.CommandUserError
	}
	if output {
		return base.CommandSuccess
	}

	switch c.Func {

	case "delete":
		switch base.Format(c.UI) {
		case "json":
			if ok := c.PrintJsonItem(resp); !ok {
				return base.CommandCliError
			}

		case "table":
			c.UI.Output("The delete operation completed successfully.")
		}

		return base.CommandSuccess

	case "list":
		switch base.Format(c.UI) {
		case "json":
			if ok := c.PrintJsonItems(resp); !ok {
				return base.CommandCliError
			}

		case "table":
			c.UI.Output(c.printListTable(items))
		}

		return base.CommandSuccess

	}

	switch base.Format(c.UI) {
	case "table":
		c.UI.Output(printItemTable(item, resp))

	case "json":
		if ok := c.PrintJsonItem(resp); !ok {
			return base.CommandCliError
		}
	}

	return base.CommandSuccess
}

func (c *Command) checkFuncError(err error) int {
	if err == nil {
		return 0
	}
	if apiErr := api.AsServerError(err); apiErr != nil {
		c.PrintApiError(apiErr, fmt.Sprintf("Error from controller when performing %s on %s", c.Func, c.plural))
		return base.CommandApiError
	}
	c.PrintCliError(fmt.Errorf("Error trying to %s %s: %s", c.Func, c.plural, err.Error()))
	return base.CommandCliError
}

var (
	flagsOnce = new(sync.Once)

	extraActionsFlagsMapFunc = func() map[string][]string { return nil }
	extraSynopsisFunc        = func(*Command) string { return "" }
	extraFlagsFunc           = func(*Command, *base.FlagSets, *base.FlagSet) {}
	extraFlagsHandlingFunc   = func(*Command, *base.FlagSets, *[]targets.Option) bool { return true }
	executeExtraActions      = func(_ *Command, inResp *api.Response, inItem *targets.Target, inItems []*targets.Target, inErr error, _ *targets.Client, _ uint32, _ []targets.Option) (*api.Response, *targets.Target, []*targets.Target, error) {
		return inResp, inItem, inItems, inErr
	}
	printCustomActionOutput = func(*Command) (bool, error) { return false, nil }
)
