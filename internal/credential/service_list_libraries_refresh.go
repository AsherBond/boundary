// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

package credential

import (
	"context"
	"time"

	"github.com/hashicorp/boundary/globals"
	"github.com/hashicorp/boundary/internal/errors"
	"github.com/hashicorp/boundary/internal/listtoken"
	"github.com/hashicorp/boundary/internal/pagination"
	"github.com/hashicorp/boundary/internal/types/resource"
	"github.com/hashicorp/boundary/internal/util"
)

// ListRefresh lists up to page size credential libraries, filtering out entries that
// do not pass the filter item function. It will automatically request
// more credential libraries from the database, at page size chunks, to fill the page.
// It will start its paging based on the information in the token.
// It returns a new list token used to continue pagination or refresh items.
// Credential libraries are ordered by update time descending (most recently updated first).
// Credential libraries may contain items that were already returned during the initial
// pagination phase. It also returns a list of any credential libraries deleted since the
// start of the initial pagination phase or last response.
func ListLibrariesRefresh(
	ctx context.Context,
	grantsHash []byte,
	pageSize int,
	filterItemFn pagination.ListFilterFunc[Library],
	tok *listtoken.Token,
	service LibraryService,
	credentialStoreId string,
) (*pagination.ListResponse[Library], error) {
	const op = "credential.ListLibrariesRefresh"

	switch {
	case len(grantsHash) == 0:
		return nil, errors.New(ctx, errors.InvalidParameter, op, "missing grants hash")
	case pageSize < 1:
		return nil, errors.New(ctx, errors.InvalidParameter, op, "page size must be at least 1")
	case filterItemFn == nil:
		return nil, errors.New(ctx, errors.InvalidParameter, op, "missing filter item callback")
	case tok == nil:
		return nil, errors.New(ctx, errors.InvalidParameter, op, "missing token")
	case util.IsNil(service):
		return nil, errors.New(ctx, errors.InvalidParameter, op, "missing service")
	case credentialStoreId == "":
		return nil, errors.New(ctx, errors.InvalidParameter, op, "missing credential store ID")
	case tok.ResourceType != resource.CredentialLibrary:
		return nil, errors.New(ctx, errors.InvalidParameter, op, "token did not have an credential library resource type")
	}
	rt, ok := tok.Subtype.(*listtoken.StartRefreshToken)
	if !ok {
		return nil, errors.New(ctx, errors.InvalidParameter, op, "token did not have a start-refresh token component")
	}

	listItemsFn := func(ctx context.Context, lastPageItem Library, limit int) ([]Library, time.Time, error) {
		opts := []Option{
			WithLimit(limit),
		}
		if lastPageItem != nil {
			opts = append(opts, WithStartPageAfterItem(lastPageItem))
		}
		// Add the database read timeout to account for any creations missed due to concurrent
		// transactions in the initial pagination phase.
		return service.ListLibrariesRefresh(ctx, credentialStoreId, rt.PreviousPhaseUpperBound.Add(-globals.RefreshReadLookbackDuration), opts...)
	}
	listDeletedIdsFn := func(ctx context.Context, since time.Time) ([]string, time.Time, error) {
		// Add the database read timeout to account for any deletions missed due to concurrent
		// transactions in previous requests.
		return service.ListDeletedLibraryIds(ctx, since.Add(-globals.RefreshReadLookbackDuration))
	}
	return pagination.ListRefresh(ctx, grantsHash, pageSize, filterItemFn, listItemsFn, service.EstimatedLibraryCount, listDeletedIdsFn, tok)
}
