package main

import (
	"fmt"
	"godoc/godoc"
)

func html2pdf() {
	html := "<p>Hello <em>pandoc</em>!</p>"
	godoc.ToPDF(html, "./hello.pdf")
	fmt.Println("------ HTML convert to pdf------")
	fmt.Println("HTML:")
	fmt.Println(html)
}
