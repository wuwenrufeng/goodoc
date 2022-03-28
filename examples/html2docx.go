package main

import (
	"fmt"
	"goodoc"
)

func html2docx() {
	html := "<p>Hello <em>pandoc</em>!</p>"
	goodoc.ToDocx(html, "./hello.docx")
	fmt.Println("------ HTML convert to docx------")
	fmt.Println("HTML:")
	fmt.Println(html)
}
