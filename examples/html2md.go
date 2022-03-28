package main

import (
	"fmt"
	"goodoc"
)

func html2md() {
	html := "<p>Hello <em>pandoc</em>!</p>"
	md, _ := goodoc.ToMarkdown(html)
	fmt.Println("------ HTML convert to markdown------")
	fmt.Println("HTML:")
	fmt.Println(html)
	fmt.Println("")
	fmt.Println("Markdown:")
	fmt.Println(md)
}
