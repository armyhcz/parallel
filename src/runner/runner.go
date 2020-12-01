package runner

import (
	"context"
	"modules/parallel/src/worker"
	"sync"
	"sync/atomic"
)

type Runner struct {
	Workers []worker.Worker
}

func (r *Runner) Start(ctx context.Context) error {
	var closed int32
	errs := make(chan error, len(r.Workers))

	defer func() {
		close(errs)
		atomic.StoreInt32(&closed, 1)
	}()

	wg := sync.WaitGroup{}
	for _, w := range r.Workers {
		wg.Add(1)
		go func(worker worker.Worker) {
			defer wg.Done()
			if err := worker.Do(ctx); err != nil {
				if closed == 0 {
					errs <- err
				}
			}
		}(w)
	}

	go func() {
		wg.Wait()
		if closed == 0 {
			errs <- nil
		}
	}()

	return <-errs
}
