package stepik

import (
	"fmt"
	"io"
	"strings"
)

func solveHoffmannDecode(r io.Reader, w io.Writer) {
	writeHoffmannDecodeOutput(w, decodeHoffman(readHoffmannDecodeInput(r)))
}

func readHoffmannDecodeInput(r io.Reader) hoffmanDecoding {
	var n, l int
	_, _ = fmt.Fscan(r, &n, &l)

	dict := make(map[string]rune)
	for i := 0; i < n; i++ {
		var letter, letterCode string
		_, _ = fmt.Fscan(r, &letter, &letterCode)
		dict[letterCode] = rune(letter[0])
	}

	var code string
	_, _ = fmt.Fscan(r, &code)

	return hoffmanDecoding{
		n:    n,
		l:    l,
		code: code,
		dict: dict,
	}
}

func writeHoffmannDecodeOutput(w io.Writer, code string) {
	_, _ = fmt.Fprintf(w, "%s\n", code)
}

func decodeHoffman(input hoffmanDecoding) string {
	var b strings.Builder
	for len(input.code) > 0 {
		for k, v := range input.dict {
			if strings.HasPrefix(input.code, k) {
				input.code = strings.TrimPrefix(input.code, k)
				b.WriteRune(v)
				break
			}
		}
	}
	return b.String()
}

type hoffmanDecoding struct {
	n, l int
	code string
	dict map[string]rune
}
