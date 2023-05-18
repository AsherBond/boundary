// Code generated by "make api"; DO NOT EDIT.
// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sessionrecordings

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/hashicorp/boundary/api"
	"github.com/hashicorp/boundary/api/scopes"
)

type SessionRecording struct {
	Id                   string                 `json:"id,omitempty"`
	Scope                *scopes.ScopeInfo      `json:"scope,omitempty"`
	SessionId            string                 `json:"session_id,omitempty"`
	StorageBucketId      string                 `json:"storage_bucket_id,omitempty"`
	BytesUp              uint64                 `json:"bytes_up,string,omitempty"`
	BytesDown            uint64                 `json:"bytes_down,string,omitempty"`
	CreatedTime          time.Time              `json:"created_time,omitempty"`
	UpdatedTime          time.Time              `json:"updated_time,omitempty"`
	StartTime            time.Time              `json:"start_time,omitempty"`
	EndTime              time.Time              `json:"end_time,omitempty"`
	Duration             api.Duration           `json:"duration,omitempty"`
	Type                 string                 `json:"type,omitempty"`
	State                string                 `json:"state,omitempty"`
	ErrorDetails         string                 `json:"error_details,omitempty"`
	MimeTypes            []string               `json:"mime_types,omitempty"`
	Endpoint             string                 `json:"endpoint,omitempty"`
	ConnectionRecordings []*ConnectionRecording `json:"connection_recordings,omitempty"`
	CreateTimeValues     *ValuesAtTime          `json:"create_time_values,omitempty"`
	AuthorizedActions    []string               `json:"authorized_actions,omitempty"`

	response *api.Response
}

type SessionRecordingReadResult struct {
	Item     *SessionRecording
	response *api.Response
}

func (n SessionRecordingReadResult) GetItem() *SessionRecording {
	return n.Item
}

func (n SessionRecordingReadResult) GetResponse() *api.Response {
	return n.response
}

type SessionRecordingListResult struct {
	Items    []*SessionRecording
	response *api.Response
}

func (n SessionRecordingListResult) GetItems() []*SessionRecording {
	return n.Items
}

func (n SessionRecordingListResult) GetResponse() *api.Response {
	return n.response
}

// Client is a client for this collection
type Client struct {
	client *api.Client
}

// Creates a new client for this collection. The submitted API client is cloned;
// modifications to it after generating this client will not have effect. If you
// need to make changes to the underlying API client, use ApiClient() to access
// it.
func NewClient(c *api.Client) *Client {
	return &Client{client: c.Clone()}
}

// ApiClient returns the underlying API client
func (c *Client) ApiClient() *api.Client {
	return c.client
}

func (c *Client) Read(ctx context.Context, id string, opt ...Option) (*SessionRecordingReadResult, error) {
	if id == "" {
		return nil, fmt.Errorf("empty id value passed into Read request")
	}
	if c.client == nil {
		return nil, fmt.Errorf("nil client")
	}

	opts, apiOpts := getOpts(opt...)

	req, err := c.client.NewRequest(ctx, "GET", fmt.Sprintf("session-recordings/%s", url.PathEscape(id)), nil, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error creating Read request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error performing client request during Read call: %w", err)
	}

	target := new(SessionRecordingReadResult)
	target.Item = new(SessionRecording)
	apiErr, err := resp.Decode(target.Item)
	if err != nil {
		return nil, fmt.Errorf("error decoding Read response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr
	}
	target.response = resp
	return target, nil
}

func (c *Client) List(ctx context.Context, scopeId string, opt ...Option) (*SessionRecordingListResult, error) {
	if scopeId == "" {
		return nil, fmt.Errorf("empty scopeId value passed into List request")
	}
	if c.client == nil {
		return nil, fmt.Errorf("nil client")
	}

	opts, apiOpts := getOpts(opt...)
	opts.queryMap["scope_id"] = scopeId

	req, err := c.client.NewRequest(ctx, "GET", "session-recordings", nil, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error creating List request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing client request during List call: %w", err)
	}

	target := new(SessionRecordingListResult)
	apiErr, err := resp.Decode(target)
	if err != nil {
		return nil, fmt.Errorf("error decoding List response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr
	}
	target.response = resp
	return target, nil
}
