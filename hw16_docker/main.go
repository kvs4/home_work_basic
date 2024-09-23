package main

import (
	"context"
	"log"
	"os"

	//"github.com/kvs4/home_work_basic/hw16_docker/client"
	"github.com/kvs4/home_work_basic/hw16_docker/internal/config"
	"github.com/kvs4/home_work_basic/hw16_docker/pkg/pgdb"
	"github.com/kvs4/home_work_basic/hw16_docker/server"
	//"github.com/spf13/pflag"
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
	//url := getURL()

	//go func() {
	server.Run()
	//}()

	//time.Sleep(time.Second * 3)

	//client.Run(url)

}

/*func getURL() string {
	var serverURL, pathRes string
	pflag.StringVarP(&serverURL, "url", "u", "", "server URL")
	pflag.StringVarP(&pathRes, "pathRes", "p", "", "resources address")
	pflag.Parse()

	var urlB strings.Builder
	urlB.WriteString(serverURL)
	urlB.WriteString(pathRes)
	return urlB.String() // "http://0.0.0.0:8080/v1/"
}*/
