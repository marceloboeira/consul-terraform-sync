package api

import (
	"context"

	"github.com/hashicorp/consul-terraform-sync/config"
)

//go:generate mockery --name=Server --filename=server.go --output=../mocks/server

// Server represents the Controller methods used for the API server
type Server interface {
	Config() config.Config

	Task(ctx context.Context, taskname string) (config.TaskConfig, error)
	TaskCreate(context.Context, config.TaskConfig) (config.TaskConfig, error)
	TaskCreateAndRun(context.Context, config.TaskConfig) (config.TaskConfig, error)
	TaskDelete(ctx context.Context, taskName string) error
	// TODO update signatures to return a new run object
	TaskInspect(context.Context, config.TaskConfig) (bool, string, string, error)
	TaskUpdate(ctx context.Context, conf config.TaskConfig, runOp string) (bool, string, string, error)
}
