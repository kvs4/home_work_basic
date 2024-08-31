package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeGetRequest(t *testing.T) {
	wantbody := `"this is method GET"`
	url := "http://0.0.0.0:8080/v1/"

	gotbody, status := MakeGetRequest(url)
	if status == -1 {
		t.Errorf("Result GET body: %s, Status: %v", gotbody, status)
	}

	assert.Equal(t, wantbody, gotbody)

}

func TestMakeGPostRequest(t *testing.T) {
	wantbody := `"this is method POST"`
	url := "http://0.0.0.0:8080/v1/"

	gotbody, status := MakeGPostRequest(url)
	if status == -1 {
		t.Errorf("Result POST body: %s, Status: %v", gotbody, status)
	}

	assert.Equal(t, wantbody, gotbody)

}
