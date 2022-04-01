package main

import (
	"fmt"
	"goodoc"
	"io/ioutil"
	"net/http"
	"net/url"
)

func transfer(w http.ResponseWriter, r *http.Request) {
	format := r.PostFormValue("format")
	file, _, err := r.FormFile("upload")
	if err != nil || file == nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "file not found!")
		return
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "file read failed!")
		return
	}
	var out []byte
	switch format {
	case goodoc.PDF:
		out, err = goodoc.ToPDF(data)
	case goodoc.DOCX:
		w.Header().Set("Content-Type", "application/msword")
		out, err = goodoc.ToDocx(data)
	case goodoc.MARKDOWN:
		format = "md"
		out, err = goodoc.ToMarkdownByte(data)
	default:
		out, err = goodoc.ToPDF(data)
	}
	if err != nil {
		fmt.Fprintf(w, "tarnsfer %s failed", format)
		return
	}
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", url.QueryEscape(fmt.Sprintf("goodoc.%s", format))))
	w.Write([]byte(out))
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/upload", transfer)
	server.ListenAndServe()
}
