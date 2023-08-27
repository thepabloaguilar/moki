package http_operations_test

import (
	"context"
	"testing"

	"github.com/google/uuid"

	"github.com/thepabloaguilar/moki/core/http_operations/mocks"
	"github.com/thepabloaguilar/moki/core/projects"

	"github.com/stretchr/testify/require"

	"github.com/thepabloaguilar/moki/core/entities"
	"github.com/thepabloaguilar/moki/core/http_operations"
	"github.com/thepabloaguilar/moki/extensions/test_resources"
	"github.com/thepabloaguilar/moki/gateways/postgres"
)

func TestCreateHTTPOperation_Failure(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name        string
		input       http_operations.CreateHTTPOperationInput
		expectedErr error
	}{
		{
			name: "should return an error when method is wrong",
			input: http_operations.CreateHTTPOperationInput{
				Method: "WRONG",
			},
			expectedErr: entities.ErrInvalidHTTPMethod,
		},
		{
			name: "should return an error when mime type is wrong",
			input: http_operations.CreateHTTPOperationInput{
				Method:   "POST",
				MIMEType: "text/css",
			},
			expectedErr: entities.ErrUnsupportedMIMEType,
		},
		{
			name: "should return an error when response status is wrong",
			input: http_operations.CreateHTTPOperationInput{
				Method:         "POST",
				MIMEType:       "application/json",
				ResponseStatus: 1000,
			},
			expectedErr: entities.ErrInvalidHTTPStatusValue,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			persistence := &mocks.CreateHTTPOperationPersistenceGatewayMock{}
			uc := http_operations.NewCreateHTTPOperation(test_resources.Now(), test_resources.UUID(), persistence)

			_, err := uc.CreateHTTPOperation(ctx, tc.input)

			require.ErrorIs(t, err, tc.expectedErr)
		})
	}
}

func TestCreateHTTPOperation_Integration_Success(t *testing.T) {
	t.Parallel()

	// SETUP
	ctx := context.Background()
	now := test_resources.Now()
	uuidProvider := test_resources.UUID()
	db := postgresResource.NewDB(t, ctx)
	persistenceGtw := &struct {
		*postgres.Projects
		*postgres.HTTPOperations
	}{
		Projects:       postgres.NewProjects(db),
		HTTPOperations: postgres.NewHTTPOperations(db),
	}
	uc := http_operations.NewCreateHTTPOperation(now, uuidProvider, persistenceGtw)

	project, err := persistenceGtw.Projects.CreateProject(ctx, entities.Project{
		ID:          uuidProvider(),
		Name:        "Test Project",
		Description: "Description",
		CreatedAt:   now(),
		UpdatedAt:   now(),
	})
	require.NoError(t, err)

	expectedOperation := entities.HTTPOperation{
		ID:             uuidProvider(),
		ProjectID:      project.ID,
		Method:         entities.HTTPMethodGet,
		MIMEType:       entities.MIMETypeJSON,
		Route:          "/my-route",
		ResponseStatus: entities.HTTPStatusOK,
		ResponseBody:   `{"response": "body"}`,
		CreateAt:       now(),
		UpdatedAt:      now(),
	}

	// TEST
	output, err := uc.CreateHTTPOperation(ctx, http_operations.CreateHTTPOperationInput{
		ProjectID:      project.ID,
		Method:         "GET",
		MIMEType:       "application/json",
		Route:          "/my-route",
		ResponseStatus: 200,
		ResponseBody:   `{"response": "body"}`,
	})

	// ASSERT
	require.NoError(t, err)
	require.Equal(t, expectedOperation, output.CreatedOperation)
}

func TestCreateHTTPOperation_Integration_Failure(t *testing.T) {
	t.Parallel()

	t.Run("should return an error when project does not exist", func(t *testing.T) {
		// SETUP
		ctx := context.Background()
		now := test_resources.Now()
		uuidProvider := test_resources.UUID()
		db := postgresResource.NewDB(t, ctx)
		persistenceGtw := &struct {
			*postgres.Projects
			*postgres.HTTPOperations
		}{
			Projects:       postgres.NewProjects(db),
			HTTPOperations: postgres.NewHTTPOperations(db),
		}
		uc := http_operations.NewCreateHTTPOperation(now, uuidProvider, persistenceGtw)

		// TEST
		_, err := uc.CreateHTTPOperation(ctx, http_operations.CreateHTTPOperationInput{
			ProjectID:      uuid.New(),
			Method:         "GET",
			MIMEType:       "application/json",
			Route:          "/my-route",
			ResponseStatus: 200,
			ResponseBody:   `{"response": "body"}`,
		})

		// ASSERT
		require.ErrorIs(t, err, projects.ErrProjectNotFound)
	})
}
