package main

import (
	"blion-worker-consenso/internal/dbx"
	"blion-worker-consenso/pkg/bk"
	"blion-worker-consenso/worker"
	"github.com/fatih/color"
	"github.com/google/uuid"
	"time"
)

func main() {
	color.Blue("Worker MTI V1.0.9 ", time.Now())
	db := dbx.GetConnection()
	srv := bk.NewServerBk(db, nil, uuid.New().String())
	wk := worker.NewWorker(srv)
	wk.Execute()
}
