package main

import (
	"blion-worker-consenso/internal/dbx"
	"blion-worker-consenso/pkg/auth"
	"blion-worker-consenso/pkg/bk"
	"blion-worker-consenso/pkg/cfg"
	"blion-worker-consenso/worker"
	"github.com/fatih/color"
	"github.com/google/uuid"
	"time"
)

func main() {
	color.Blue("Worker BLion V1.0.0 ", time.Now())
	db := dbx.GetConnection()
	srvBk := bk.NewServerBk(db, nil, uuid.New().String())
	srvAuth := auth.NewServerAuth(db, nil, uuid.New().String())
	srvCfg := cfg.NewServerCfg(db, nil, uuid.New().String())
	wk := worker.NewWorker(srvBk, srvAuth, srvCfg)
	wk.Execute()
}
