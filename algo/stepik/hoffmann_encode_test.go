package stepik

import (
	"bytes"
	"strings"
	"testing"
)

func TestReadHoffmannEncodeInput(t *testing.T) {
	expected := "abacabad"
	input := readHoffmannEncodeInput(strings.NewReader(expected))

	if input != expected {
		t.Errorf("Wrong input read %s, expected %s", input, expected)
	}
}

func TestWriteHoffmannEncodeOutput(t *testing.T) {
	buf := new(bytes.Buffer)
	encoding := hoffmanEncoding{
		dict: map[rune]string{
			'a': "0",
			'b': "10",
			'c': "110",
			'd': "111",
		},
		code: "01001100100111",
	}

	writeHoffmannEncodeOutput(buf, encoding)

	expected := `4 14
a: 0
b: 10
c: 110
d: 111
01001100100111`
	if buf.String() != expected {
		t.Errorf("Wrong output \n%s expected\n%s", expected, buf.String())
	}
}

func TestEncodeHoffman(t *testing.T) {
	enc := encodeHoffman("abacabad")

	expectedCode := "01001100100111"
	if enc.code != expectedCode {
		t.Errorf("Wrong result code %s, expected %s", enc.code, expectedCode)
	}
	expectedDictLen := 4
	if len(enc.dict) != expectedDictLen {
		t.Errorf("Wrong dict len %v expected %v", len(enc.dict), expectedDictLen)
	}
}

func TestSolveSingleLetter(t *testing.T) {
	buf := new(bytes.Buffer)
	solveHoffmanEncode(strings.NewReader(`a 

`), buf)

	expectedOutput := `1 1
a: 0
0`
	if buf.String() != expectedOutput {
		t.Errorf("Wrong result code %s, expected %s", buf.String(), expectedOutput)
	}
}

func TestSolve2(t *testing.T) {
	buf := new(bytes.Buffer)
	solveHoffmanEncode(strings.NewReader("abacabad"), buf)

	expectedOutput := `4 14
a: 0
b: 10
c: 110
d: 111
01001100100111`
	if buf.String() != expectedOutput {
		t.Errorf("Wrong result code %s, expected %s", buf.String(), expectedOutput)
	}
}

func TestSolveLong(t *testing.T) {
	buf := new(bytes.Buffer)
	solveHoffmanEncode(strings.NewReader("LoremipsumdolorsitametconsecteturadipisicingelitDelectuseligendiillumisteomnisvoluptatumAspernaturatqueautautemculpadoloremeveniethicillumnequenesciuntofficiispossimusrepudiandaevoluptatevoluptatem"), buf)
	expectedPrefix := "22 794"
	if !strings.HasPrefix(buf.String(), expectedPrefix) {
		t.Errorf("Wrong result code \n%s, expected prefix \n%s", buf.String(), expectedPrefix)
	}
}

func TestSolveSingleLetterManyTimes(t *testing.T) {
	buf := new(bytes.Buffer)
	solveHoffmanEncode(strings.NewReader("aaaaaaaa"), buf)

	expectedOutput := `1 8
a: 0
00000000`
	if buf.String() != expectedOutput {
		t.Errorf("Wrong result code %s, expected %s", buf.String(), expectedOutput)
	}
}
