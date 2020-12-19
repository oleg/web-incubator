package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

func main() {
	allPosts := []Post{
		{Id: 1, Content: "Hello", Author: "John"},
		{Id: 2, Content: "Bonjour", Author: "Jack"},
		{Id: 3, Content: "Hola", Author: "Jain"},
	}

	err := writeCsvFile("csvfile/posts.csv", allPosts)
	if err != nil {
		panic(err)
	}

	posts, err := readCsvFile("csvfile/posts.csv")
	if err != nil {
		panic(err)
	}
	for _, p := range posts {
		fmt.Println(p)
	}

}

func writeCsvFile(filename string, allPosts []Post) error {
	csvFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)

	for _, post := range allPosts {
		line := []string{strconv.Itoa(post.Id), post.Content, post.Author}
		err := writer.Write(line)
		if err != nil {
			return err
		}
	}
	writer.Flush()
	return nil
}

func readCsvFile(filename string) ([]Post, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var posts []Post
	for _, item := range records {
		id, _ := strconv.Atoi(item[0])
		post := Post{Id: id, Content: item[1], Author: item[2]}
		posts = append(posts, post)
	}
	return posts, nil
}
