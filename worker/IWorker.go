package worker

import "blion-worker-consenso/pkg/bk/lotteries"

type IWorker interface {
	Execute()
	doWork(workItem *lotteries.Lottery)
}
