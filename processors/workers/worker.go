package workers

import (
	"context"
)

type WorkerContext struct{}

type Worker struct {
	context WorkerContext
}

func InitWorkers() *Worker {
	return &Worker{
		context: WorkerContext{},
	}
}

func (w *Worker) StartWorkers() {
	w.InitEventWorkers().Start(context.Background())
}
