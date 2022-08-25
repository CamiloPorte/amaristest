package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var datapool = []struct {
	inputv   string
	inputu   string
	isfriend bool
}{
	{"tokyo", "otoky", true},
	{"tokyo", "yotok", true},
	{"tokyo", "kyoto", true},
	{"tokyo", "yotok", true},
	{"tokyo", "otoky", true},
	{"tokyo", "tokyo", true},
	{"tokyo", "toky", false},
	{"hi", "ih", true},
	{"hi", "ihi", false},
}

func TestGetDataByPoint(t *testing.T) {
	for _, dp := range datapool {
		answ := StringsFriends(dp.inputu, dp.inputv)
		assert.Equal(t, dp.isfriend, answ)
	}
}
