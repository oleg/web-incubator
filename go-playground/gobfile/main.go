package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

func main() {
	post1 := Post{Id: 1, Content: "Hello", Author: "John"}
	err := store(post1, "gobfile/post1")
	if err != nil {
		panic(err)
	}

	var postRead Post
	err = load(&postRead, "gobfile/post1")
	if err != nil {
		panic(err)
	}
	fmt.Println(postRead)
}

func store(data interface{}, filename string) error {
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(data)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, buffer.Bytes(), 0600)
	if err != nil {
		return err
	}
	return nil
}

func load(data interface{}, filename string) error {
	raw, err := ioutil.ReadFile(filename)
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
