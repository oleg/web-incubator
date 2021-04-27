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
	"path/filepath"
	"sort"
	"strings"
	"words/sub"
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
	corsHeaders(w)
	x := []Word{
		{"hello", 30},
		{"my", 20},
		{"friend", 10},
	}

	name := r.URL.Query().Get("name")
	var t ListCreate
	err := load(&t, name)
	if err == nil {
		x = t.Words
	}
	encoder := json.NewEncoder(w)
	err = encoder.Encode(x)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", "words")
}

type ListCreate struct {
	Name  string `json:"name"`
	Words []Word `json:"words"`
}
type List struct {
	Name string `json:"name"`
}
type ListsList struct {
	Lists []List `json:"lists"`
}

func lists(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		createList(w, r)
	case http.MethodGet:
		listLists(w, r)
	default:
		corsHeaders(w)
		w.WriteHeader(200)
	}
}

func createList(w http.ResponseWriter, r *http.Request) {
	corsHeaders(w)
	w.Header().Set("Content-Type", "application/json")
	var t ListCreate
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		log.Fatal(err)
	}
	err = store(t, t.Name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", "create list")
	return
}

func listLists(w http.ResponseWriter, r *http.Request) {
	corsHeaders(w)
	w.Header().Set("Content-Type", "application/json")
	names, err := listDir(dataDir())
	if err != nil {
		log.Fatal(err)
	}
	lists := make([]List, 0)
	for _, name := range names {
		lists = append(lists, List{Name: name})
	}
	err = json.NewEncoder(w).Encode(ListsList{Lists: lists})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("lists list")
}

func listDir(dataDir string) ([]string, error) {
	dirs, err := os.ReadDir(dataDir)
	if err != nil {
		return nil, err
	}
	var names []string
	for _, dir := range dirs {
		if !dir.IsDir() {
			names = append(names, dir.Name())
		}
	}
	return names, nil
}

func upload(w http.ResponseWriter, r *http.Request) {
	corsHeaders(w)
	file, _, err := r.FormFile("subfile")
	if err != nil {
		log.Fatal(err)
	}
	all, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	allString := sub.GetWords(sub.GetLines(bytes.NewReader(all)))
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
		log.Fatal(err)
	}
	fmt.Printf("%v\n", "upload")
}

func store(data interface{}, filename string) error {
	buffer := new(bytes.Buffer)
	err := gob.NewEncoder(buffer).Encode(data)
	if err != nil {
		return err
	}
	err = os.WriteFile(filepath.Join(dataDir(), filename), buffer.Bytes(), 0600)
	if err != nil {
		return err
	}
	return nil
}

func load(data interface{}, filename string) error {
	raw, err := os.ReadFile(filepath.Join(dataDir(), filename))
	if err != nil {
		return err
	}
	err = gob.NewDecoder(bytes.NewReader(raw)).Decode(data)
	if err != nil {
		return err
	}
	return nil
}

func dataDir() string {
	path := dataDirPath()
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
	return path
}

func dataDirPath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "words"
	}
	return filepath.Join(homeDir, ".local", "words")
}

func corsHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}
