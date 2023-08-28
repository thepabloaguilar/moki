package mock_test

import (
	"context"
	"testing"

	"github.com/thepabloaguilar/moki/core/mock/mocks"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"github.com/thepabloaguilar/moki/core/entities"
	"github.com/thepabloaguilar/moki/core/mock"
	"github.com/thepabloaguilar/moki/gateways/postgres"
)

func TestExecuteMock_Failure(t *testing.T) {
	t.Parallel()

	t.Run("should return an error when http method is wrong", func(t *testing.T) {
		// SETUP
		ctx := context.Background()
		persistence := &mocks.ExecuteMockPersistenceGatewayMock{}
		uc := mock.NewExecuteMock(persistence)

		// TEST
		_, err := uc.ExecuteMock(ctx, mock.ExecuteMockInput{
			ProjectID:  uuid.New(),
			HTTPMethod: "WRONG",
			Route:      "",
		})

		// ASSERT
		require.ErrorIs(t, err, entities.ErrInvalidHTTPMethod)
	})
}

func TestExecuteMock_Integration_Success(t *testing.T) {
	t.Parallel()

	// SETUP
	ctx := context.Background()
	db := postgresResource.NewDB(t, ctx)
	persistence := &struct {
		*postgres.Projects
		*postgres.HTTPOperations
	}{
		Projects:       postgres.NewProjects(db),
		HTTPOperations: postgres.NewHTTPOperations(db),
	}
	uc := mock.NewExecuteMock(persistence)
	expectedOutput := mock.ExecuteMockOutput{
		ResponseBody:     `{"my-body": "message"}`,
		ResponseMimeType: entities.MIMETypeJSON,
		ResponseStatus:   entities.HTTPStatusBadRequest,
	}

	project, err := persistence.Projects.CreateProject(ctx, entities.Project{
		ID:          uuid.New(),
		Name:        "Test Project",
		Description: "Description",
	})
	require.NoError(t, err)

	_, err = persistence.HTTPOperations.CreateHTTPOperation(ctx, entities.HTTPOperation{
		ProjectID:      project.ID,
		Method:         entities.HTTPMethodPost,
		MIMEType:       entities.MIMETypeJSON,
		Route:          "/my-test/route",
		ResponseStatus: entities.HTTPStatusBadRequest,
		ResponseBody:   `{"my-body": "message"}`,
	})
	require.NoError(t, err)

	// TEST
	output, err := uc.ExecuteMock(ctx, mock.ExecuteMockInput{
		ProjectID:  project.ID,
		HTTPMethod: "POST",
		Route:      "/my-test/route",
	})

	// ASSERT
	require.NoError(t, err)
	require.Equal(t, expectedOutput, output)
}
