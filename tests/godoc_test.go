package tests

import (
	"goodoc/goodoc"
	"testing"
)

func TestToMarkdown(t *testing.T) {
	tests := []struct {
		name string
		src  string
	}{
		// TODO: Add test cases.
		{"", "<p>Hello <em>pandoc</em>!</p>"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			markdown, err := goodoc.ToMarkdown(tt.src)
			if err != nil {
				t.Errorf("to markdown error: %v", err)
			}
			t.Log(markdown)
		})
	}
}

func TestToHTML(t *testing.T) {
	tests := []struct {
		name string
		src  string
	}{
		// TODO: Add test cases.
		{"", `---
		title: Test
		...
		
		# Test!
		
		This is a test of *pandoc*.
		
		- list one
		- list two`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			html, err := goodoc.ToHTML(tt.src)
			if err != nil {
				t.Errorf("to html error: %v", err)
			}
			t.Log(html)
		})
	}
}

func TestToPDF(t *testing.T) {
	tests := []struct {
		name string
		src  string
		path string
	}{
		// TODO: Add test cases.
		{"", "<p>Hello <em>pandoc</em>!</p>", "./test1.pdf"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := goodoc.ToPDF(tt.src, tt.path)
			if err != nil {
				t.Errorf("to pdf error: %v, out: %v", err, out)
			}
		})
	}
}
