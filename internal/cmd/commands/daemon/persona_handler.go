// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

package daemon

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/hashicorp/boundary/internal/cmd/base"
	"github.com/hashicorp/boundary/internal/daemon/cache"
	"github.com/hashicorp/boundary/internal/errors"
	"github.com/hashicorp/boundary/internal/util"
)

type refresher interface {
	refresh()
}

// upsertPersonaRequest is the payload sent as the body of a post request
// and processed by this handler. All except for the AuthToken are always
// required. When KeyringType is not "none" then AuthToken should not be
// set and all other fields are required.  When it is "none" AuthToken must be
// set and the TokenName should match the AuthTokenId.
type upsertPersonaRequest struct {
	KeyringType  string
	TokenName    string
	BoundaryAddr string
	AuthTokenId  string
	AuthToken    string
}

func newPersonaHandlerFunc(ctx context.Context, repo *cache.Repository, refresher refresher) (http.HandlerFunc, error) {
	const op = "daemon.newPersonaHandlerFunc"
	switch {
	case util.IsNil(repo):
		return nil, errors.New(ctx, errors.InvalidParameter, op, "repository is nil")
	case util.IsNil(refresher):
		return nil, errors.New(ctx, errors.InvalidParameter, op, "refresher is nil")
	}
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		if r.Method != http.MethodPost {
			writeError(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		var perReq upsertPersonaRequest

		data, err := io.ReadAll(r.Body)
		if err != nil {
			writeError(w, "unable to read request body", http.StatusBadRequest)
			return
		}
		if err := json.Unmarshal(data, &perReq); err != nil {
			writeError(w, "unable to parse request body", http.StatusBadRequest)
			return
		}

		switch {
		case perReq.TokenName == "":
			writeError(w, "TokenName is a required field but was empty", http.StatusBadRequest)
			return
		case perReq.BoundaryAddr == "":
			writeError(w, "BoundaryAddr is a required field but was empty", http.StatusBadRequest)
			return
		case perReq.AuthTokenId == "":
			writeError(w, "AuthTokenId is a required field but was empty", http.StatusBadRequest)
			return
		}

		var opts []cache.Option
		switch perReq.KeyringType {
		case base.NoneKeyring:
			switch {
			case perReq.TokenName != perReq.AuthTokenId:
				writeError(w, "When KeyringType is 'none' TokenName must match the AuthTokenId", http.StatusBadRequest)
				return
			case perReq.AuthToken == "":
				writeError(w, "AuthToken is empty when KeyringType is 'none'", http.StatusBadRequest)
				return
			}
			opts = append(opts, cache.WithAuthToken(perReq.AuthToken))
		case "":
			writeError(w, "KeyringType is a required field but was empty", http.StatusBadRequest)
			return
		}

		p, err := repo.LookupPersona(ctx, perReq.TokenName, perReq.KeyringType)
		if err != nil {
			writeError(w, "error performing persona lookup", http.StatusInternalServerError)
			return
		}

		if err = repo.AddPersona(ctx, perReq.BoundaryAddr, perReq.TokenName, perReq.KeyringType, perReq.AuthTokenId, opts...); err != nil {
			writeError(w, "Failed to add a persona", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)

		if p == nil || p.AuthTokenId != perReq.AuthTokenId {
			// If this was a new persona or an updated auth token refresh the cache.
			refresher.refresh()
		}
	}, nil
}
