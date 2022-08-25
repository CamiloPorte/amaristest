package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var datapool = []struct {
	inputnames string
	names      []string
	len        int
}{
	{"Luis,Camilo,Andres,Laura", []string{"Andres", "Camilo", "Laura", "Luis"}, 4},
	{"Luis", []string{"Luis"}, 1},
	{"Luis,luis,Andres,Angelo,angelo", []string{"Andres", "Angelo", "Luis", "angelo", "luis"}, 5},
	{"", nil, 0},
}

func TestGetDataByPoint(t *testing.T) {
	for _, dp := range datapool {
		answ, lenAnsw := GetDataByPoint(dp.inputnames)
		assert.Equal(t, dp.names, answ)
		assert.Equal(t, dp.len, lenAnsw)
	}
}
