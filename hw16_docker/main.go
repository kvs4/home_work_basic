package main

import (
	"context"
	"log"
	"os"

	"github.com/kvs4/home_work_basic/hw16_docker/internal/config"
	"github.com/kvs4/home_work_basic/hw16_docker/pkg/pgdb"
	"github.com/kvs4/home_work_basic/hw16_docker/server"
)

func init() {
	if err := config.Load(".env"); err != nil {
		log.Fatalf("Didn`t read .env config: %v", err)
		return
	}
	ctx := context.Background()

	if err := pgdb.New(ctx, os.Getenv("DB_DSN")); err != nil {
		log.Fatal("@[main] can't init service s3client: ", err)
		return
	}
}

func main() {
	server.Run()
}
