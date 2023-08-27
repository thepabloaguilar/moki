package projects

import "errors"

var (
	ErrEmptyProjectName = errors.New("project name cannot be empty")
	ErrProjectNotFound  = errors.New("project not found")
)
