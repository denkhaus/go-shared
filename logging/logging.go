package logging

import (
	"context"
	"os"

	"golang.org/x/sync/errgroup"
	"gopkg.in/pipe.v2"
)

// blocking
func StreamDockerComposeLogs(ctx context.Context, serviceName string) error {
	line := pipe.Line(
		pipe.Exec("docker-compose", "logs", "-f"),
		pipe.Exec("grep", serviceName),
	)

	eg, _ := errgroup.WithContext(ctx)

	eg.Go(func() error {
		s := pipe.NewState(os.Stdout, os.Stderr)
		err := line(s)
		if err == nil {
			return s.RunTasks()
		}

		return err
	})

	return eg.Wait()
}
