package projects_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/thepabloaguilar/moki/core/entities"
	"github.com/thepabloaguilar/moki/core/projects"
	"github.com/thepabloaguilar/moki/extensions/test_resources"
	"github.com/thepabloaguilar/moki/gateways/postgres"
)

func TestCreateProject_Integration_Success(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	now := test_resources.Now()
	uuidProvider := test_resources.UUID()
	pg := postgres.NewProjects(postgresResource.NewDB(t, ctx))
	uc := projects.NewCreateProject(now, uuidProvider, pg)
	expectedProject := entities.Project{
		ID:          uuidProvider(),
		Name:        "Jujutsu Kaisen",
		Description: "One of the best animes",
		CreatedAt:   now(),
		UpdatedAt:   now(),
	}

	output, err := uc.CreateProject(ctx, projects.CreateProjectInput{
		ProjectName:        "Jujutsu Kaisen",
		ProjectDescription: "One of the best animes",
	})

	require.NoError(t, err)
	require.Equal(t, expectedProject, output.CreatedProject)
}

func TestCreateProject_Integration_Failure(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name        string
		input       projects.CreateProjectInput
		expectedErr error
	}{
		{
			name:        "should return an error when name is empty",
			input:       projects.CreateProjectInput{},
			expectedErr: projects.ErrEmptyProjectName,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			now := test_resources.Now()
			uuidProvider := test_resources.UUID()
			pg := postgres.NewProjects(postgresResource.NewDB(t, ctx))
			uc := projects.NewCreateProject(now, uuidProvider, pg)

			_, err := uc.CreateProject(ctx, tc.input)

			require.ErrorIs(t, err, tc.expectedErr)
		})
	}
}
