// Code generated by "make api"; DO NOT EDIT.
// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workers

type RemoteStorageState struct {
	Status      string                    `json:"status,omitempty"`
	Permissions *RemoteStoragePermissions `json:"permissions,omitempty"`
}
