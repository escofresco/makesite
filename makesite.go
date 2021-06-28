package main

import (
	"html/template"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Post struct {
	Title   string
	Content string
}

func main() {
	dat, err := ioutil.ReadFile("first-post.txt")
	check(err)
	post := Post{"asdf", string(dat)}
	f, err := os.Create("first-post.html")
	t := template.Must(template.New("post.tmpl").ParseGlob("*.tmpl"))
	err = t.Execute(f, post)
	if err != nil {
		panic(err)
	}
}
