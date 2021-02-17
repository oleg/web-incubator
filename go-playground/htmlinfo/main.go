package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path"
)

func main() {
	indexPage := template.Must(template.ParseFiles("templates/index.html"))

	//http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	//
	//})
	http.Handle("/", &PageInfoHandler{
		PageAnalyser: &RegexPageAnalyser{},
		client:       &http.Client{},
		templ:        indexPage,
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func ShowForm(w http.ResponseWriter, r *http.Request) {
	book := PageInfo{}
	fp := path.Join("templates", "index.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, book); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

type PageInfo struct {
	HtmlVersion       string
	PageTitle         string
	Headers           map[string]int //count by level
	InternalLinks     int
	ExternalLinks     int
	InaccessibleLinks int
	HasLoginForm      bool
	Content           string
}

type PageAnalyser interface {
	Analyse(page string) (PageInfo, error)
}

type PageInfoHandler struct {
	PageAnalyser
	client *http.Client
	templ  *template.Template
}

func (h *PageInfoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) //todo client error?
		return
	}
	address := r.FormValue("address")
	if address == "" {
		if err := h.templ.Execute(w, &PageInfo{}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	pageContent, err := h.fetchPage(address)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest) //todo client error?
		return
	}
	pageInfo, err := h.Analyse(pageContent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest) //todo client error?
		return
	}
	if err := h.templ.Execute(w, pageInfo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *PageInfoHandler) fetchPage(address string) (string, error) {
	response, err := h.client.Get("http://" + address)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(bodyBytes), nil
}
