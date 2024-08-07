// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

package server

import (
	stderrors "errors"
	"fmt"

	"github.com/hashicorp/boundary/internal/daemon/controller/handlers"
	pb "github.com/hashicorp/boundary/sdk/pbs/controller/api/resources/targets"
	"github.com/hashicorp/boundary/sdk/pbs/plugin"
	"github.com/hashicorp/boundary/version"
	"github.com/hashicorp/go-bexpr"
	"github.com/mitchellh/pointerstructure"
	"google.golang.org/grpc/codes"
)

const ManagedWorkerTag = "boundary.cloud.hashicorp.com:managed"

// WorkerList is a helper type to make the selection of workers clearer and more declarative.
type WorkerList []*Worker

// addresses converts the slice of workers to a slice of their addresses
func (w WorkerList) Addresses() []string {
	ret := make([]string, 0, len(w))
	for _, worker := range w {
		ret = append(ret, worker.GetAddress())
	}
	return ret
}

// PublicIds converts the slice of workers to a slice of public ids of those
// workers.
func (w WorkerList) PublicIds() []string {
	ret := make([]string, 0, len(w))
	for _, worker := range w {
		ret = append(ret, worker.GetPublicId())
	}
	return ret
}

// workerInfos converts the slice of workers to a slice of their workerInfo protos
func (w WorkerList) WorkerInfos() []*pb.WorkerInfo {
	ret := make([]*pb.WorkerInfo, 0, len(w))
	for _, worker := range w {
		ret = append(ret, &pb.WorkerInfo{Address: worker.GetAddress()})
	}
	return ret
}

// SupportsFeature returns a new WorkerList composed of all workers in this
// WorkerList which supports the provided feature.
func (w WorkerList) SupportsFeature(f version.Feature) WorkerList {
	var ret []*Worker
	for _, worker := range w {
		sv := version.FromVersionString(worker.GetReleaseVersion()).Semver()
		if version.SupportsFeature(sv, f) {
			ret = append(ret, worker)
		}
	}
	return ret
}

// filtered returns a new workerList where all elements contained in it are the
// ones which from the original workerList that pass the evaluator's evaluation.
func (w WorkerList) Filtered(eval *bexpr.Evaluator) (WorkerList, error) {
	var ret []*Worker
	for _, worker := range w {
		filterInput := map[string]any{
			"name": worker.GetName(),
			"tags": worker.CanonicalTags(),
		}
		ok, err := eval.Evaluate(filterInput)
		if err != nil && !stderrors.Is(err, pointerstructure.ErrNotFound) {
			return nil, handlers.ApiErrorWithCodeAndMessage(
				codes.FailedPrecondition,
				fmt.Sprintf("Worker filter expression evaluation resulted in error: %s", err))
		}
		if ok {
			ret = append(ret, worker)
		}
	}
	return ret, nil
}

// SeparateManagedWorkers divides the incoming workers into managed and
// unmanaged workers, respectively
func SeparateManagedWorkers(workers WorkerList) (managedWorkers, nonManagedWorkers WorkerList) {
	// Build a set of managed and unmanaged workers
	managedWorkers = make([]*Worker, 0, len(workers))
	nonManagedWorkers = make([]*Worker, 0, len(workers))
	for _, worker := range workers {
		if IsManagedWorker(worker) {
			managedWorkers = append(managedWorkers, worker)
		} else {
			nonManagedWorkers = append(nonManagedWorkers, worker)
		}
	}
	return managedWorkers, nonManagedWorkers
}

// IsManagedWorker indicates whether the given worker is managed
func IsManagedWorker(worker *Worker) bool {
	return len(worker.CanonicalTags()[ManagedWorkerTag]) != 0
}

// FilterWorkersByLocalStorageState filters the workers by their local storage state.
// Workers that support local storage state feature will be considered healthy if their
// local storage state is Available.
// If the worker does not have any workers in the Available local storage state, it
// will return workers with Unknown local storage state.
// Workers that do not support local storage state will be considered healthy.
func FilterWorkersByLocalStorageState(workers WorkerList) (healthyWorkers WorkerList) {
	availableWorkers := make([]*Worker, 0, len(workers))
	unknownWorkers := make([]*Worker, 0, len(workers))

	for _, worker := range workers {
		sv := version.FromVersionString(worker.GetReleaseVersion()).Semver()
		if version.SupportsFeature(sv, version.LocalStorageState) {
			ls := worker.GetLocalStorageState()
			if ls == AvailableLocalStorageState.String() {
				availableWorkers = append(availableWorkers, worker)
			} else if ls == UnknownLocalStorageState.String() {
				unknownWorkers = append(unknownWorkers, worker)
			}
		} else {
			availableWorkers = append(availableWorkers, worker)
		}
	}

	if len(availableWorkers) > 0 {
		return availableWorkers
	}
	return unknownWorkers
}

// FilterStorageBucketCredentialStateFn is a function definition that is used to filter
// out workers that are considered to be in a unhealthy state. The function should return
// true for healthy workers.
type FilterStorageBucketCredentialStateFn func(*plugin.StorageBucketCredentialState) bool

// FilterStorageBucketCredentialByWriteAccess will return true if the write state is missing or
// if the state is in an OK or UNKNOWN state.
func FilterStorageBucketCredentialByWriteAccess(sbcState *plugin.StorageBucketCredentialState) bool {
	if sbcState == nil || sbcState.State == nil || sbcState.State.Write == nil {
		return true
	}
	return sbcState.State.Write.State != plugin.StateType_STATE_TYPE_ERROR
}

// FilterStorageBucketCredentialByReadAccess will return true if the read state is missing or
// if the state is in an OK or UNKNOWN state.
func FilterStorageBucketCredentialByReadAccess(sbcState *plugin.StorageBucketCredentialState) bool {
	if sbcState == nil || sbcState.State == nil || sbcState.State.Read == nil {
		return true
	}
	return sbcState.State.Read.State != plugin.StateType_STATE_TYPE_ERROR
}

// FilterStorageBucketCredentialByDeleteAccess will return true if the delete state is missing or
// if the state is in an OK or UNKNOWN state.
func FilterStorageBucketCredentialByDeleteAccess(sbcState *plugin.StorageBucketCredentialState) bool {
	if sbcState == nil || sbcState.State == nil || sbcState.State.Delete == nil {
		return true
	}
	return sbcState.State.Delete.State != plugin.StateType_STATE_TYPE_ERROR
}
