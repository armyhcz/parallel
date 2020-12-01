package parallel

import (
	"context"
	"errors"
	"fmt"
	"modules/parallel/src/worker"
	"testing"
	"time"
)

func TestParallel_Start(t *testing.T) {
	p, _ := New()
	p.Add(worker.NewFuncWorker(func(ctx context.Context, args ...interface{}) error {
		select {
		case <- time.After(time.Second):
			fmt.Println("worker1 start")
		case <- ctx.Done():
			fmt.Println("worker caceled")
		}

		return nil
	}))
	p.Add(worker.NewFuncWorker(func(ctx context.Context, args ...interface{}) error {
		select {
		case <- time.After(time.Second * 3):
			fmt.Println("worker2 start")
		case <- ctx.Done():
			fmt.Println("worker2 canceled")
		}
		return errors.New("err")
	}))

	if err := p.Start(context.Background()); err != nil {
		t.Error(err)
	}
}
