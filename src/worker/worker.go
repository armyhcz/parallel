package worker

import "context"

type Worker interface {
	Do(ctx context.Context) error
}

type FuncWorker func(ctx context.Context, args ...interface{}) error
type funcWorker struct {
	worker FuncWorker
	args   []interface{}
}

func (w *funcWorker) Do(ctx context.Context) error {
	return w.worker(ctx, w.args...)
}

func NewFuncWorker(worker FuncWorker, args ...interface{}) *funcWorker {
	return &funcWorker{
		worker: worker,
		args:   args,
	}
}
