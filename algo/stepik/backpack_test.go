package stepik

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestWritePrice(t *testing.T) {
	b := new(bytes.Buffer)

	writePrice(b, 0.123456)

	assert.Equal(t, "0.123", b.String())
}

func TestReadBackpack(t *testing.T) {
	input := strings.NewReader(
		`3 50
60 20
100 50
120 30`)

	b := readBackpack(input)

	assert.Equal(t, 50.0, b.volume)
	assert.Equal(t, item{price: 60, volume: 20}, b.items[0])
	assert.Equal(t, item{price: 100, volume: 50}, b.items[1])
	assert.Equal(t, item{price: 120, volume: 30}, b.items[2])
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
