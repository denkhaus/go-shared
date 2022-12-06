package workflow

import (
	"context"
	"fmt"

	"github.com/denkhaus/go-shared/zaplog"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.uber.org/zap"
	zapadapter "logur.dev/adapter/zap"
	"logur.dev/logur"
)

// ExecuteWorkflow bootstraps a workflow with the params provided by provideParams
func Execute(ctx context.Context,
	workFlow interface{},
	workflowName string,
	taskQueueName string,
	provideParams func(*zap.Logger) (interface{}, error),
) error {
	logger := zaplog.Logger()
	c, err := client.NewClient(client.Options{
		Logger:   logur.LoggerToKV(zapadapter.New(logger)),
		HostPort: client.DefaultHostPort,
	})
	if err != nil {
		return errors.Wrap(err, "NewClient")
	}
	defer c.Close()

	workflowOptions := client.StartWorkflowOptions{
		ID:        fmt.Sprintf("%s_%s", taskQueueName, uuid.New()),
		TaskQueue: taskQueueName,
	}

	params, err := provideParams(logger)
	if err != nil {
		return errors.Wrap(err, "provideParams")
	}

	we, err := c.ExecuteWorkflow(ctx, workflowOptions, workFlow, params)
	if err != nil {
		return errors.Wrap(err, "ExecuteWorkflow")
	}

	logger.Info("started workflow",
		zap.String("name", workflowName),
		zap.String("workflowID", we.GetID()),
		zap.String("runID", we.GetRunID()),
	)

	return nil
}

// blocking
func Run(workFlow interface{}, taskQueueName string, activities ...interface{}) error {
	logger := zaplog.Logger()
	c, err := client.NewLazyClient(client.Options{
		Logger:   logur.LoggerToKV(zapadapter.New(logger)),
		HostPort: client.DefaultHostPort,
	})
	if err != nil {
		return errors.Wrap(err, "NewLazyClient")
	}
	defer c.Close()

	w := worker.New(c, taskQueueName, worker.Options{})
	w.RegisterWorkflow(workFlow)

	for _, act := range activities {
		w.RegisterActivity(act)
	}

	if err := w.Run(worker.InterruptCh()); err != nil {
		return errors.Wrap(err, "Run")
	}

	return nil
}
