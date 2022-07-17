package stepik

import (
	"algo/assert"
	"bytes"
	"strings"
	"testing"
)

func TestWritePrice(t *testing.T) {
	b := new(bytes.Buffer)

	writePrice(b, 0.123456)

	assert.Equal(t, b.String(), "0.123")
}

func TestReadBackpack(t *testing.T) {
	input := strings.NewReader(
		`3 50
60 20
100 50
120 30`)

	b := readBackpack(input)

	assert.Equal(t, b.volume, 50.0)
	assert.Equal(t, b.items[0], item{price: 60, volume: 20})
	assert.Equal(t, b.items[1], item{price: 100, volume: 50})
	assert.Equal(t, b.items[2], item{price: 120, volume: 30})
}

func TestCalculatePrice(t *testing.T) {
	bp := backpack{
		volume: 50,
		items: []item{
			{price: 60, volume: 20},
			{price: 100, volume: 50},
			{price: 120, volume: 30},
		},
	}

	price := calculatePrice(bp)

	if price != 180.0 {
		t.Errorf("Wrong price %v", price)
	}

}
