package godoc

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

const (
	HTML     = "html"
	MARKDOWN = "markdown"
	PDF      = "pdf"
)

func LookPandoc() {
	path, err := exec.LookPath("pandoc")
	if err != nil {
		panic("pandoc not found")
	}
	log.Println(path)
}

func ToMarkdown(src string) (markdown string, err error) {
	markdown, err = bash(fmt.Sprintf("pandoc -f %s -t %s", HTML,
		MARKDOWN), src)
	if err != nil {
		return
	}

	return
}

func ToHTML(src string) (html string, err error) {
	html, err = bash(fmt.Sprintf("pandoc -f %s -t %s", MARKDOWN,
		HTML), src)
	if err != nil {
		return
	}

	return
}

func ToPDF(src string, path string) (out string, err error) {
	out, err = bash(fmt.Sprintf("pandoc -f %s -t %s -o %s", HTML,
		PDF, path), src)
	if err != nil {
		return
	}

	return
}

func bash(bash, content string) (out string, err error) {
	var buf bytes.Buffer

	cmd := exec.Command("/bin/sh", "-c", bash)
	cmd.Stdin = strings.NewReader(content)
	cmd.Stderr = &buf
	cmd.Stdout = &buf

	err = cmd.Run()
	if err != nil {
		cmd.Process.Release()
		buf.Reset()
		return
	}

	out = buf.String()

	// Clean up resource
	cmd.Process.Kill()
	buf.Reset()

	return
}
