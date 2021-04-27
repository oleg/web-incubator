package sub

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func GetLines(reader io.Reader) []string {
	liens := make([]string, 0)
	scanner := bufio.NewScanner(reader)
	count := 1
	timeStamp := false
	for scanner.Scan() {
		text := strings.Trim(scanner.Text(), "\ufeff")
		itoa := strconv.Itoa(count)
		if text == itoa {
			timeStamp = true
			count++
			continue
		}
		if timeStamp {
			timeStamp = false
			continue
		}
		if text == "" {
			continue
		}
		liens = append(liens, text)
	}
	return liens
}

func GetWords(lines []string) []string {
	words := make([]string, 0)
	for _, line := range lines {
		line = strings.ReplaceAll(line, "?", " ")
		line = strings.ReplaceAll(line, ",", " ")
		line = strings.ReplaceAll(line, ".", " ")
		line = strings.ReplaceAll(line, "!", " ")
		line = strings.ReplaceAll(line, "-", " ")
		line = strings.ReplaceAll(line, "'", " ")
		line = strings.Trim(line, " ")
		split := strings.Split(line, " ")
		for _, s := range split {
			if s != "" {
				words = append(words, s)
			}
		}
	}
	return words
}
