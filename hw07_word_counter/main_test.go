package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountWords(t *testing.T) {
	text := "five, Five,  one.. two: three/// one   five"
	wantMap := make(map[string]int)
	wantMap["five"] = 3
	wantMap["one"] = 2
	wantMap["two"] = 1
	wantMap["three"] = 1

	gotMap := countWords(text)
	assert.Equal(t, wantMap, gotMap)
}

func TestCountWordsEmpty(t *testing.T) {
	text := ""
	wantMap := make(map[string]int)

	gotMap := countWords(text)
	assert.Equal(t, wantMap, gotMap)
}
