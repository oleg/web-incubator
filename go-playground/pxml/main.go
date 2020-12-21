package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type Post struct {
	XMLName  xml.Name  `xml:"post"`
	Id       string    `xml:"id,attr"`
	Content  string    `xml:"content"`
	Author   Author    `xml:"author"`
	Xml      string    `xml:",innerxml"`
	Comments []Comment `xml:"comments>comment"`
}

type Author struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

type Comment struct {
	Id      string `xml:"id,attr"`
	Content string `xml:"content"`
	Author  Author `xml:"author"`
}

var post = &Post{
	Id:      "1",
	Content: "hello world",
	Author: Author{
		Id:   "2",
		Name: "Joe",
	},
}

func main() {
	fmt.Printf("post:\n%v\n", unmarshalPost("pxml/post.xml"))
	comments, err := decodeComments()
	if err != nil {
		panic(err)
	}
	fmt.Printf("comments:\n%v", comments)

	err = marshalPost("pxml/post-m.xml", post)
	if err != nil {
		panic(err)
	}
}

func marshalPost(filename string, post *Post) error {
	output, err := xml.MarshalIndent(post, "", "\t")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, []byte(xml.Header+string(output)), 0644)
}

func unmarshalPost(filename string) *Post {
	xmlFile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer xmlFile.Close()

	xmlData, err := ioutil.ReadAll(xmlFile)
	var post Post
	xml.Unmarshal(xmlData, &post)
	return &post
}

func decodeComments() ([]Comment, error) {
	xmlFile, err := os.Open("pxml/post.xml")
	if err != nil {
		return nil, err
	}
	defer xmlFile.Close()
	comments := make([]Comment, 0)
	decoder := xml.NewDecoder(xmlFile)
	for {
		t, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "comment" {
				var comment Comment
				err = decoder.DecodeElement(&comment, &se)
				if err != nil {
					return nil, err
				}
				comments = append(comments, comment)
			}
		}
	}
	return comments, nil
}
