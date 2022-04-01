package goodoc

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
	DOCX     = "docx"
)

func init() {
	lookPath()
}

func lookPath() {
	path, err := exec.LookPath("pandoc")
	if err != nil {
		panic("pandoc not found")
	}
	log.Printf("pandoc path is %s", path)
	path, err = exec.LookPath("latex")
	if err != nil {
		panic("latex not found")
	}
	log.Printf("latex path is %s", path)
}

func ToMarkdown(src string) (markdown string, err error) {
	markdown, err = bash(fmt.Sprintf("pandoc -f %s -t %s", HTML,
		MARKDOWN), src)
	if err != nil {
		return
	}

	return
}

func ToMarkdownByte(src []byte) (markdown []byte, err error) {
	markdown, err = bashByte(fmt.Sprintf("pandoc -f %s -t %s", HTML,
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

func ToPDFWithPath(src string, path string) (out string, err error) {
	out, err = bash(fmt.Sprintf("pandoc -f %s -t %s -o %s", HTML,
		PDF, path), src)
	if err != nil {
		return
	}

	return
}

func ToPDF(src []byte) (out []byte, err error) {
	out, err = bashByte(fmt.Sprintf("pandoc -f %s -t %s", HTML,
		PDF), src)
	if err != nil {
		return
	}

	return
}

func ToDocxWithPath(src string, path string) (out string, err error) {
	out, err = bash(fmt.Sprintf("pandoc -f %s -t %s -o %s", HTML,
		DOCX, path), src)
	if err != nil {
		return
	}

	return
}

func ToDocx(src []byte) (out []byte, err error) {
	out, err = bashByte(fmt.Sprintf("pandoc -f %s -t %s", HTML,
		DOCX), src)
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

func bashByte(bash string, content []byte) (out []byte, err error) {
	var buf bytes.Buffer

	cmd := exec.Command("/bin/sh", "-c", bash)
	cmd.Stdin = bytes.NewReader(content)
	cmd.Stdout = &buf

	err = cmd.Run()
	if err != nil {
		cmd.Process.Release()
		buf.Reset()
		return
	}

	out = buf.Bytes()

	// Clean up resource
	cmd.Process.Kill()
	buf.Reset()

	return
}
