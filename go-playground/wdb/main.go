package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "user=gwp password=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}
}
func main() {
	post1 := Post{Content: "hello", Author: "joe"}
	err := post1.Create()
	if err != nil {
		panic(err)
	}

	post2 := Post{Content: "bonjour", Author: "ji"}
	err = post2.Create()
	if err != nil {
		panic(err)
	}

	post, err := getPost(post2.Id)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n\n", post)

	posts, err := Posts(10)
	if err != nil {
		panic(err)
	}
	for _, p := range posts {
		fmt.Printf("%#v\n", p)
	}

	err = posts[0].Delete()
	if err != nil {
		panic(err)
	}
	posts[3].Content = "A"
	posts[3].Author = "B"
	err = posts[3].Update()
	if err != nil {
		panic(err)
	}

	posts, err = Posts(10)
	if err != nil {
		panic(err)
	}
	fmt.Println()
	for _, p := range posts {
		fmt.Printf("%#v\n", p)
	}
}

type Post struct {
	Id      int
	Content string
	Author  string
}

func (post *Post) Create() (err error) {
	statement := "insert into posts (content, author) values($1, $2) returning id"
	stmt, err := db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	return
}

func getPost(id int) (post Post, err error) {
	post = Post{}
	err = db.
		QueryRow("select id, content, author from posts where id = $1", id).
		Scan(&post.Id, &post.Content, &post.Author)

	return
}

func (post *Post) Update() (err error) {
	_, err = db.Exec("update posts set content = $2, author = $3 where id = $1",
		post.Id, post.Content, post.Author)
	return
}

func (post *Post) Delete() (err error) {
	_, err = db.Exec("delete from posts where id = $1", post.Id)
	return
}

func Posts(limit int) (posts []Post, err error) {
	rows, err := db.Query("select id, content, author from posts order by id asc limit $1", limit)
	if err != nil {
		return
	}
	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.Id, &post.Content, &post.Author)
		if err != nil {
			return
		}
		posts = append(posts, post)
	}
	rows.Close()
	return
}
