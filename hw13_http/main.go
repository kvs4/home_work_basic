package main

import (
	"strings"
	"time"

	"github.com/kvs4/home_work_basic/hw13_http/client"
	"github.com/kvs4/home_work_basic/hw13_http/server"
	"github.com/spf13/pflag"
)

func main() {
	url := getURL()

	go func() {
		server.Run()
	}()

	time.Sleep(time.Second * 2)

	client.Run(url)
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
