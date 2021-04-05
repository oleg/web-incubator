package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

//todo 0, 1 element
func dirTree(out io.Writer, path string, files bool) error {
	return dirTreeLevel(out, path, files, "")
}
func dirTreeLevel(out io.Writer, path string, files bool, prefix string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	entries, err := f.Readdir(0)
	if err != nil {
		return err
	}

	names := make([]string, 0)
	for _, v := range entries {
		if v.IsDir() {
			names = append(names, v.Name())
		}
		if !v.IsDir() && files {
			name := fmt.Sprintf("%s (%db)", v.Name(), v.Size())
			if v.Size() == 0 {
				name = fmt.Sprintf("%s (empty)", v.Name())
			}
			names = append(names, name)
		}
	}
	sort.Strings(names)

	last := len(names) - 1
	for i, v := range names {
		line := "├───"
		prev := "│"
		if i == last {
			line = "└───"
			prev = ""
		}
		out.Write([]byte(prefix + line + v + "\n"))
		dirTreeLevel(out, path+"/"+v, files, prefix+prev+"\t")
	}
	return nil
}
