package worker

type IWorker interface {
	Execute()
	doWork()
}
