package sub

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var text = `1
00:00:09,660 --> 00:00:12,040
Mi scusi. Puo' lasciarci
una dichiarazione?

2
00:00:17,440 --> 00:00:20,560
Si, signora, rimanga qui dietro.

3
00:00:24,620 --> 00:00:26,080
Per favore, state indietro.

4
00:00:26,090 --> 00:00:28,250
- Capitano.
- Agente Lisbon.

5
00:00:28,580 --> 00:00:30,330
Non crediate che
abbiamo bisogno di voi.

6
00:00:30,340 --> 00:00:32,400
Ci piace il ragazzo dei vicini
che ha trovato il corpo.

7
00:00:32,410 --> 00:00:33,910
Ha confessato?
`

func TestGetLines(t *testing.T) {
	lines := GetLines(strings.NewReader(text))

	expected := []string{
		"Mi scusi. Puo' lasciarci",
		"una dichiarazione?",
		"Si, signora, rimanga qui dietro.",
		"Per favore, state indietro.",
		"- Capitano.",
		"- Agente Lisbon.",
		"Non crediate che",
		"abbiamo bisogno di voi.",
		"Ci piace il ragazzo dei vicini",
		"che ha trovato il corpo.",
		"Ha confessato?",
	}
	assert.Equal(t, expected, lines)
}

func TestGetWords(t *testing.T) {
	lines := GetWords(GetLines(strings.NewReader(text)))

	expected := []string{
		"Mi", "scusi", "Puo", "lasciarci",
		"una", "dichiarazione",
		"Si", "signora", "rimanga", "qui", "dietro",
		"Per", "favore", "state", "indietro",
		"Capitano", "Agente", "Lisbon",
		"Non", "crediate", "che",
		"abbiamo", "bisogno", "di", "voi",
		"Ci", "piace", "il", "ragazzo", "dei", "vicini",
		"che", "ha", "trovato", "il", "corpo",
		"Ha", "confessato",
	}
	assert.Equal(t, expected, lines)
}

