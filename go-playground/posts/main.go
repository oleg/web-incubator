package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"io"
	"net/http"
	"path"
	"strconv"
)

type Post struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "user=gwp dbname=gwp password=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func main() {
	server := http.Server{Addr: "127.0.0.1:8080"}
	http.HandleFunc("/post/", handleRequest)
	server.ListenAndServe()
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = handleGet(w, r)
	case "POST":
		err = handlePost(w, r)
	case "PUT":
		err = handlePut(w, r)
	case "DELETE":
		err = handleDelete(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleGet(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return err
	}
	post, err := retrieve(id)
	if err != nil {
		return err
	}
	output, err := json.MarshalIndent(&post, "", "\t")
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(output)
	if err != nil {
		return err
	}
	return nil
}

func handlePost(w http.ResponseWriter, r *http.Request) error {
	body := make([]byte, r.ContentLength)
	_, err := r.Body.Read(body)
	if err != nil && err != io.EOF {
		fmt.Println("failed to read body", err)
		return err
	}
	var post Post
	err = json.Unmarshal(body, &post)
	if err != nil {
		fmt.Println("failed to unmarshal post", err)
		return err
	}
	err = store(&post)
	if err != nil {
		fmt.Println("failed to store post", err)
		return err
	}
	w.WriteHeader(http.StatusOK)
	return nil
}

func handlePut(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return err
	}
	post, err := retrieve(id)
	if err != nil {
		return err
	}
	body := make([]byte, r.ContentLength)
	_, err = r.Body.Read(body)
	if err != nil && err != io.EOF {
		return err
	}
	err = json.Unmarshal(body, &post)
	if err != nil {
		return err
	}
	err = update(&post)
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusOK)
	return nil
}

func handleDelete(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return err
	}
	post, err := retrieve(id)
	if err != nil {
		return err
	}
	err = delete(&post)
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusOK)
	return nil
}

func retrieve(id int) (Post, error) {
	post := Post{}
	err := db.QueryRow("select id, content, author from posts where id = $1", id).
		Scan(&post.Id, &post.Author, &post.Content)
	return post, err
}

func store(post *Post) error {
	stmt, err := db.Prepare("insert into posts(content, author) values ($1, $2) returning id")
	if err != nil {
		return err
	}
	defer stmt.Close()
	return stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
}

func update(post *Post) error {
	_, err := db.Exec("update posts set content = $2, author = $3 where id = $1",
		post.Id, post.Content, post.Author)
	return err
}

func delete(post *Post) error {
	_, err := db.Exec("delete from posts where id = $1", post.Id)
	return err
}
