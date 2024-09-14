package main

import (
	"context"
	"log"
	"os"
	"strings"
	"time"

	"github.com/kvs4/home_work_basic/hw15_go_sql/client"
	"github.com/kvs4/home_work_basic/hw15_go_sql/internal/config"
	"github.com/kvs4/home_work_basic/hw15_go_sql/pkg/pgdb"
	"github.com/kvs4/home_work_basic/hw15_go_sql/server"
	"github.com/spf13/pflag"
)

func init() {
	if err := config.Load(".env"); err != nil {
		log.Fatal("Didn`t read .env config")
		return
	}
	ctx := context.Background()

	if err := pgdb.New(ctx, os.Getenv("DB_DSN")); err != nil {
		log.Fatal("@[main] can't init service s3client: ", err)
		return
	}
}

func main() {
	url := getURL()

	go func() {
		server.Run()
	}()

	time.Sleep(time.Second * 3)

	client.Run(url)

	/*repo := repository.New(pgdb.DB.Conn())
	ctx := context.Background()
	result, err := repo.GetUsers(ctx)
	if err != nil {
		fmt.Printf("success: false, msg: %v\n", err)
		return
	}
	fmt.Println("result =", result)*/
}

func getURL() string {
	var serverURL, pathRes string
	pflag.StringVarP(&serverURL, "url", "u", "", "server URL")
	pflag.StringVarP(&pathRes, "pathRes", "p", "", "resources address")
	pflag.Parse()

	var urlB strings.Builder
	urlB.WriteString(serverURL)
	urlB.WriteString(pathRes)
	return urlB.String() // "http://0.0.0.0:8080/v1/"
}
