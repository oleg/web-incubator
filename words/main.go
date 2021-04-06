package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"sort"
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
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func words(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	encoder := json.NewEncoder(w)
	x := []Word{
		{"hello", 30},
		{"my", 20},
		{"friend", 10},
	}
	err := encoder.Encode(x)
	if err != nil {
		log.Print(err)
	}
	//fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	fmt.Printf("%v\n", "words")
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

	//fmt.Printf("%v\n", s)

	reg := regexp.MustCompile(`\p{L}+`)
	allString := reg.FindAllString(s, -1)
	m := make(map[string]int, len(allString))
	for _, v := range allString {
		m[v]++
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
