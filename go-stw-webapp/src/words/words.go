package words

import (
	"net/http"
	tt "text/template"
	"strings"
	"bytes"
	//	"fmt"

	"sort"
	"io/ioutil"
	"regexp"
	"encoding/json"
	"appengine"
	"appengine/user"
	"appengine/datastore"
)

type WordCount struct {
	Word  string
	Count int
}

type UserWord struct {
	Email string
	Word  string
}

func (uw UserWord) Key() string { return uw.Email + ":" + uw.Word }

type WordCountArray []WordCount

func (p WordCountArray) Len() int { return len(p) }
func (p WordCountArray) Less(i, j int) bool { return p[i].Count > p[j].Count }

//>desc <asc
func (p WordCountArray) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

type menuLink struct {
	Name  string
	Url   string
	Class string
}

func init() {
	http.HandleFunc("/", index)
	http.HandleFunc("/upload", upload)
	http.HandleFunc("/tokenize", tokenize)
	http.HandleFunc("/add", add)
	http.HandleFunc("/words", words)
	http.HandleFunc("/inout", inout)
}

func index(w http.ResponseWriter, r *http.Request) {
	renderPage("/", "index.html", nil, r, w)
}

func upload(w http.ResponseWriter, r *http.Request) {
	renderPage("/upload", "upload.html", nil, r, w)
}

func tokenize(w http.ResponseWriter, r *http.Request) {
	filename, content, err := safeFormFile("somename", r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userWords, _ := userWords(r)
	words := countWords(content, userWords)
	wordsCountResult := toWordCountArray(words)
	sort.Sort(wordsCountResult)

	//rows := toRows(wordsCountResult)

	data := map[string]interface{}{
		"Filename": filename,
		"Words":    wordsCountResult,
	}

	renderPage("/upload", "tokenize.html", data, r, w)
}

func inout(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	u := user.Current(c)
	var url string
	var err error
	if u == nil {
		url, err = user.LoginURL(c, "/")
	} else {
		url, err = user.LogoutURL(c, "/")
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Location", url)
	w.WriteHeader(http.StatusFound)
}

func add(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	u := user.Current(c)

	uw := UserWord{
		Email: u.Email,
		Word:  r.FormValue("word"),
	}

	k := datastore.NewKey(c, "UserWord", uw.Key(), 0, nil)
	_, err := datastore.Put(c, k, &uw)

	var m map[string]string
	if err == nil {
		m = map[string]string{"ok": "true"}
	} else {
		m = map[string]string{"ok": "false", "msg": err.Error()}
	}

	b, _ := json.Marshal(m)

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func words(w http.ResponseWriter, r *http.Request) {
	words, err := userWords(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	renderPage("/words", "words.html", words, r, w)
}

func userWords(r *http.Request) ([]UserWord, error) {
	c := appengine.NewContext(r)
	u := user.Current(c)
	words := make([]UserWord, 0)
	if u != nil {
		q := datastore.NewQuery("UserWord").Filter("Email =", u.Email)
		_, err := q.GetAll(c, &words)
		return words, err
	}
	return words, nil
}

func countWords(val string, userWords []UserWord) map[string]int {
	words := make(map[string]int)
	cleaner, _ := regexp.Compile("[0-9.,!?:;><'\"]|--|'s|'d|'ll")

	for _, word := range strings.Fields(val) {
		word = cleaner.ReplaceAllString(word, "")
		if len(word) == 0 || len(word) == 1 {
			continue
		}
		word = strings.ToLower(word)
		if isUserWord(word, userWords) {
			continue
		}
		if count, ok := words[word]; ok {
			words[word] = count + 1
		} else {
			words[word] = 1
		}
	}
	return words
}

func isUserWord(word string, userWords []UserWord) bool {
	for _, uw := range userWords {
		if word == uw.Word {
			return true
		}
	}
	return false
}

func toWordCountArray(m map[string]int) WordCountArray {
	wordsCountResult := WordCountArray{}
	for k, v := range m {
		wordsCountResult = append(wordsCountResult, WordCount{k, v})
	}
	return wordsCountResult
}

func safeFormFile(fieldName string, r *http.Request) (filename string, content string, err error) {
	file, handler, fileErr := r.FormFile(fieldName)
	if fileErr != nil {
		return "", "", fileErr
	}
	bytesContent, readErr := ioutil.ReadAll(file)
	if readErr != nil {
		return "", "", readErr
	}
	return handler.Filename, string(bytesContent), nil
}

/*
   data := map[string]interface{} {
    "Filename": "123",
    "Words": "3321",
  }

*/
func renderPage(active string, page string, pageData interface{}, r *http.Request, wr http.ResponseWriter) {
	menuData := []menuLink{
		makeMenuLink("home", "/", active),
		makeMenuLink("upload", "/upload", active),
		makeMenuLink("paste", "/paste", active),
		makeMenuLink("words", "/words", active),
		makeMenuLink("login/logout", "/inout", active)}

	menu, e1 := renderToString("_menu.html", menuData)
	content, e2 := renderToString(page, pageData)
	words, e3 := userWords(r)
	sidebar, e4 := renderToString("_sidebar.html", words)

	if e1 != nil {
		http.Error(wr, e1.Error(), http.StatusInternalServerError)
	}
	if e2 != nil {
		http.Error(wr, e2.Error(), http.StatusInternalServerError)
	}
	if e3 != nil {
		http.Error(wr, e3.Error(), http.StatusInternalServerError)
	}
	if e4 != nil {
		http.Error(wr, e4.Error(), http.StatusInternalServerError)
	}


	renderTemplate(menu, content, sidebar, wr)
}

func makeMenuLink(name string, url string, curr string) (result menuLink) {
	class := ""
	if url == curr {
		class = "current_page_item"
	}
	return menuLink{name, url, class}
}

func renderTemplate(menu string, content string, sidebar string, wr http.ResponseWriter) {
	siteStructure := map[string]interface{}{
		"Menu":    menu,
		"Content": content,
		"Sidebar": sidebar,
	}
	render("_template.html", siteStructure, wr)
}

func renderToString(name string, data interface{}) (rendered string, error error) {
	tmpl := tt.New(name)

	b := bytes.NewBuffer([]byte{})
	_, err := tmpl.ParseFiles("html/" + name)
	if err == nil {
		err = tmpl.Execute(b, data)
	}
	return b.String(), err
}

//TODO merge with renderToString ???
func render(name string, data interface{}, wr http.ResponseWriter) {
	tmpl := tt.New(name)

	_, err := tmpl.ParseFiles("html/" + name)
	if err == nil {
		err = tmpl.Execute(wr, data)
	}
	if err != nil {
		http.Error(wr, err.Error(), http.StatusInternalServerError)
	}
}
