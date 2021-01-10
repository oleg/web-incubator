package main

func main() {
	e := encryption{
		subjectNumber: 7,
		cardPublicKey: 2069194,
		doorPublicKey: 16426071,
	}
	e.findLoopSizes()

	println(e.key())
}

type encryption struct {
	subjectNumber               int
	cardPublicKey, cardLoopSize int
	doorPublicKey, doorLoopSize int
}

func (e *encryption) findLoopSizes() {
	e.cardLoopSize = findLoopSize(e.subjectNumber, e.cardPublicKey)
	e.doorLoopSize = findLoopSize(e.subjectNumber, e.doorPublicKey)
}

func (e *encryption) key() int {
	value := 1
	for i := 0; i < e.cardLoopSize; i++ {
		value *= e.doorPublicKey
		value %= 20201227
	}
	return value
}

func findLoopSize(subjectNumber, publicKey int) int {
	for value, i := 1, 1; ; i++ {
		value *= subjectNumber
		value %= 20201227
		if value == publicKey {
			return i
		}
	}
}
