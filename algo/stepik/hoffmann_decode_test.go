package stepik

import (
	"strings"
	"testing"
)

func TestReadHoffmannDecodeInput(t *testing.T) {
	enc := readHoffmannDecodeInput(strings.NewReader(`4 14
a: 0
b: 10
c: 110
d: 111
01001100100111
`))
	expectedCode := "01001100100111"
	if enc.code != expectedCode {
		t.Errorf("Wrong code %s, expected %s", enc.code, expectedCode)
	}
	expectedDictLen := 4
	if len(enc.dict) != expectedDictLen {
		t.Errorf("Wrong length %d expected %v", len(enc.dict), expectedDictLen)
	}
}

func TestDecodeHoffmann(t *testing.T) {
	decoding := hoffmanDecoding{
		n:    4,
		l:    14,
		code: "01001100100111",
		dict: map[string]rune{
			"0":   'a',
			"10":  'b',
			"110": 'c',
			"111": 'd',
		},
	}

	decoded := decodeHoffman(decoding)

	expected := "abacabad"
	if decoded != expected {
		t.Errorf("Wrong decoded string %v expected %v", decoded, expected)
	}
}
