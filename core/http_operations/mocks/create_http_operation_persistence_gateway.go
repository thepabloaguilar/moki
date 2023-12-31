// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mocks

import (
	"context"
	"github.com/google/uuid"
	"github.com/thepabloaguilar/moki/core/entities"
	"github.com/thepabloaguilar/moki/core/http_operations"
	"sync"
)

// Ensure, that CreateHTTPOperationPersistenceGatewayMock does implement http_operations.CreateHTTPOperationPersistenceGateway.
// If this is not the case, regenerate this file with moq.
var _ http_operations.CreateHTTPOperationPersistenceGateway = &CreateHTTPOperationPersistenceGatewayMock{}

// CreateHTTPOperationPersistenceGatewayMock is a mock implementation of http_operations.CreateHTTPOperationPersistenceGateway.
//
//	func TestSomethingThatUsesCreateHTTPOperationPersistenceGateway(t *testing.T) {
//
//		// make and configure a mocked http_operations.CreateHTTPOperationPersistenceGateway
//		mockedCreateHTTPOperationPersistenceGateway := &CreateHTTPOperationPersistenceGatewayMock{
//			CreateHTTPOperationFunc: func(ctx context.Context, operation entities.HTTPOperation) (entities.HTTPOperation, error) {
//				panic("mock out the CreateHTTPOperation method")
//			},
//			GetProjectByIDFunc: func(ctx context.Context, projectID uuid.UUID) (entities.Project, error) {
//				panic("mock out the GetProjectByID method")
//			},
//		}
//
//		// use mockedCreateHTTPOperationPersistenceGateway in code that requires http_operations.CreateHTTPOperationPersistenceGateway
//		// and then make assertions.
//
//	}
type CreateHTTPOperationPersistenceGatewayMock struct {
	// CreateHTTPOperationFunc mocks the CreateHTTPOperation method.
	CreateHTTPOperationFunc func(ctx context.Context, operation entities.HTTPOperation) (entities.HTTPOperation, error)

	// GetProjectByIDFunc mocks the GetProjectByID method.
	GetProjectByIDFunc func(ctx context.Context, projectID uuid.UUID) (entities.Project, error)

	// calls tracks calls to the methods.
	calls struct {
		// CreateHTTPOperation holds details about calls to the CreateHTTPOperation method.
		CreateHTTPOperation []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Operation is the operation argument value.
			Operation entities.HTTPOperation
		}
		// GetProjectByID holds details about calls to the GetProjectByID method.
		GetProjectByID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ProjectID is the projectID argument value.
			ProjectID uuid.UUID
		}
	}
	lockCreateHTTPOperation sync.RWMutex
	lockGetProjectByID      sync.RWMutex
}

// CreateHTTPOperation calls CreateHTTPOperationFunc.
func (mock *CreateHTTPOperationPersistenceGatewayMock) CreateHTTPOperation(ctx context.Context, operation entities.HTTPOperation) (entities.HTTPOperation, error) {
	callInfo := struct {
		Ctx       context.Context
		Operation entities.HTTPOperation
	}{
		Ctx:       ctx,
		Operation: operation,
	}
	mock.lockCreateHTTPOperation.Lock()
	mock.calls.CreateHTTPOperation = append(mock.calls.CreateHTTPOperation, callInfo)
	mock.lockCreateHTTPOperation.Unlock()
	if mock.CreateHTTPOperationFunc == nil {
		var (
			hTTPOperationOut entities.HTTPOperation
			errOut           error
		)
		return hTTPOperationOut, errOut
	}
	return mock.CreateHTTPOperationFunc(ctx, operation)
}

// CreateHTTPOperationCalls gets all the calls that were made to CreateHTTPOperation.
// Check the length with:
//
//	len(mockedCreateHTTPOperationPersistenceGateway.CreateHTTPOperationCalls())
func (mock *CreateHTTPOperationPersistenceGatewayMock) CreateHTTPOperationCalls() []struct {
	Ctx       context.Context
	Operation entities.HTTPOperation
} {
	var calls []struct {
		Ctx       context.Context
		Operation entities.HTTPOperation
	}
	mock.lockCreateHTTPOperation.RLock()
	calls = mock.calls.CreateHTTPOperation
	mock.lockCreateHTTPOperation.RUnlock()
	return calls
}

// GetProjectByID calls GetProjectByIDFunc.
func (mock *CreateHTTPOperationPersistenceGatewayMock) GetProjectByID(ctx context.Context, projectID uuid.UUID) (entities.Project, error) {
	callInfo := struct {
		Ctx       context.Context
		ProjectID uuid.UUID
	}{
		Ctx:       ctx,
		ProjectID: projectID,
	}
	mock.lockGetProjectByID.Lock()
	mock.calls.GetProjectByID = append(mock.calls.GetProjectByID, callInfo)
	mock.lockGetProjectByID.Unlock()
	if mock.GetProjectByIDFunc == nil {
		var (
			projectOut entities.Project
			errOut     error
		)
		return projectOut, errOut
	}
	return mock.GetProjectByIDFunc(ctx, projectID)
}

// GetProjectByIDCalls gets all the calls that were made to GetProjectByID.
// Check the length with:
//
//	len(mockedCreateHTTPOperationPersistenceGateway.GetProjectByIDCalls())
func (mock *CreateHTTPOperationPersistenceGatewayMock) GetProjectByIDCalls() []struct {
	Ctx       context.Context
	ProjectID uuid.UUID
} {
	var calls []struct {
		Ctx       context.Context
		ProjectID uuid.UUID
	}
	mock.lockGetProjectByID.RLock()
	calls = mock.calls.GetProjectByID
	mock.lockGetProjectByID.RUnlock()
	return calls
}
