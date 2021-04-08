package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strings"
)

type Word struct {
	Text string `json:"text"`
	Freq int    `json:"freq"`
}

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3001"
	}
	http.HandleFunc("/words", words)
	http.HandleFunc("/upload", upload)
	http.HandleFunc("/lists", lists)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func words(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	x := []Word{
		{"hello", 30},
		{"my", 20},
		{"friend", 10},
	}

	var t ListCreate
	err := load(&t, "default")
	if err == nil {
		x = t.Words
	}
	encoder := json.NewEncoder(w)
	err = encoder.Encode(x)
	if err != nil {
		log.Print(err)
	}
	fmt.Printf("%v\n", "words")
}

type ListCreate struct {
	Name  string `json:"name"`
	Words []Word `json:"words"`
}

func lists(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Access-Control-Allow-Origin", "*")
	//w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	decoder := json.NewDecoder(r.Body)
	var t ListCreate
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}

	err = store(t, t.Name)
	if err != nil {
		panic(err)
	}

	fmt.Fprintln(w, `{"hello": "world"}`)

	fmt.Printf("%v\n", "lists")
}

func upload(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	file, _, err := r.FormFile("subfile")
	if err != nil {
		log.Print(err)
	}

	all, err := io.ReadAll(file)
	if err != nil {
		log.Print(err)
	}

	s := string(all)

	s = regexp.MustCompile(`([0-9]|>|-|<|,|!|\.|\?|:|\\)`).ReplaceAllString(s, " ")
	allString := regexp.MustCompile(`\s`).Split(s, -1)
	m := make(map[string]int, len(allString))
	for _, v := range allString {
		m[v]++
	}

	delete(m, "")

	for k, v := range m {
		kLower := strings.ToLower(k)
		if _, found := m[kLower]; found {
			delete(m, k)
			m[kLower] += v
		}
	}

	x := make([]Word, 0, len(m))
	for k, v := range m {
		x = append(x, Word{k, v})
	}

	sort.Slice(x, func(i, j int) bool {
		return x[i].Freq > x[j].Freq
	})

	encoder := json.NewEncoder(w)
	err = encoder.Encode(x)
	if err != nil {
		log.Print(err)
	}
	fmt.Printf("%v\n", "upload")
}
func store(data interface{}, filename string) error {
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(data)
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, buffer.Bytes(), 0600)
	if err != nil {
		return err
	}
	return nil
}

func load(data interface{}, filename string) error {
	raw, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	buffer := bytes.NewBuffer(raw)
	dec := gob.NewDecoder(buffer)
	err = dec.Decode(data)
	if err != nil {
		return err
	}
	return nil
}
