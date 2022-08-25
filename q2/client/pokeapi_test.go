package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var datapool = []struct {
	id        int
	answName  string
	answError error
}{
	{1, "bulbasaur", nil},
	{2, "ivysaur", nil},
	{182, "bellossom", nil},
	{18239, "", nil},
	{1829, "", nil},
	{555, "darmanitan-standard", nil},
	{448, "lucario", nil},
}

func TestGetDataByPoint(t *testing.T) {
	PA := NewPokeApi("https://pokeapi.co/api/v2/pokemon-form/")
	for _, dp := range datapool {
		answ, err := PA.Get(dp.id)
		assert.Equal(t, dp.answName, answ)
		assert.Equal(t, dp.answError, err)
	}
}
