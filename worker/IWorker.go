package worker

import "blion-worker-consenso/pkg/bk/lottery"

type IWorker interface {
	Execute()
	doWork(workItem *lottery.Lottery)
}
