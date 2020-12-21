package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Post struct {
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

var testPost = &Post{
	Id:      1,
	Content: "hello world",
	Author: Author{
		Id:   2,
		Name: "Joe",
	},
	Comments: []Comment{
		{
			Id:      10,
			Content: "Good",
			Author:  "A.J.",
		}, {
			Id:      20,
			Content: "Yes, Good",
			Author:  "JJ",
		},
	},
}

func main() {
	readPost, err := unmarshalPost("pjson/post.json")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", readPost)

	err = marshalPost("pjson/post-m.json", testPost)
	if err != nil {
		panic(err)
	}
}

func marshalPost(filename string, p *Post) error {
	data, err := json.MarshalIndent(p, "", "\t")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, data, 0644)
}

func unmarshalPost(filename string) (Post, error) {
	var post Post
	jsonFile, err := os.Open(filename)
	if err != nil {
		return post, err
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return post, err
	}

	err = json.Unmarshal(jsonData, &post)
	if err != nil {
		return post, err
	}
	return post, nil
}
