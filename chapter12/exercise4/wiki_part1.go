package main

import (
	"fmt"
	"io/ioutil"
)

type Page struct {
	Title string
	Body  []byte
}

// Save writes a to a text file with Title as the name and body as the content.
func (p *Page) Save() {
	err := ioutil.WriteFile(p.Title, p.Body, 777)
	if err != nil {
		fmt.Println("Write error", err)
	}
}

// Load, when given a title string reads the corresponding text file.
func Load(s string) (p *Page) {
	p = new(Page)
	p.Title = s
	f, err := ioutil.ReadFile(s)
	if err != nil {
		fmt.Println("Read error", err)
	}
	p.Body = f
	return p
}

func main() {
	p := new(Page)
	p.Title = "test.txt"
	p.Body = []byte("The defined file mode bits are the most significant bits of the FileMode.")
	p.Save()

	f := Load("test.txt")
	fmt.Println(f)
}
