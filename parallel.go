package parallel

import (
	"context"
	"modules/parallel/src/runner"
	"modules/parallel/src/worker"
)

type Parallel struct {
	workers []worker.Worker
}

func (p *Parallel) Start(ctx context.Context) error {
	cancelCtx, cancel := context.WithCancel(ctx)

	r := runner.Runner{Workers: p.workers[:]}
	if err := r.Start(cancelCtx); err != nil {
		cancel()
		return err
	}
	return nil
}

func (p *Parallel) Add(worker worker.Worker)  {
	p.workers = append(p.workers, worker)
}

func New() (*Parallel, error)  {
	return &Parallel{}, nil
}
