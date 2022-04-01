package main

import (
	"fmt"
	"goodoc"
)

func html2pdf() {
	html := "<p>Hello <em>pandoc</em>!</p>"
	goodoc.ToPDFWithPath(html, "./hello.pdf")
	fmt.Println("------ HTML convert to pdf------")
	fmt.Println("HTML:")
	fmt.Println(html)
}
