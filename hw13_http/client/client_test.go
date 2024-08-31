package client

import (
	"testing"
	"time"

	"github.com/kvs4/home_work_basic/hw13_http/server"
	"github.com/stretchr/testify/assert"
)

func runServer() {
	go func() {
		server.Run()
	}()

	time.Sleep(time.Second * 2)
}

func TestMakeGetRequest(t *testing.T) {
	runServer()

	wantbody := `"this is method GET"`
	url := "http://0.0.0.0:8080/v1/"

	gotbody, status := MakeGetRequest(url)
	if status == -1 {
		t.Errorf("Result GET body: %s, Status: %v", gotbody, status)
	}

	assert.Equal(t, wantbody, gotbody)

}

func TestMakeGPostRequest(t *testing.T) {
	runServer()

	wantbody := `"this is method POST"`
	url := "http://0.0.0.0:8080/v1/"

	gotbody, status := MakeGPostRequest(url)
	if status == -1 {
		t.Errorf("Result POST body: %s, Status: %v", gotbody, status)
	}

	assert.Equal(t, wantbody, gotbody)

}
