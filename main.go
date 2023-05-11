package main

import (
	"io"
	"net/http"
	"os"
	"os/exec"
	"time"
)

func convert2pdf(w http.ResponseWriter, req *http.Request) {
	file, handler, err := req.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	defer file.Close()

	// if file exists, create unique name
	docxFile := "/tmp/" + handler.Filename
	if _, err := os.Stat(docxFile); err == nil {
		docxFile = "/tmp/" + time.Now().Format("20060102150405") + handler.Filename
	}

	// create temp file
	f, err := os.Create(docxFile)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer f.Close()

	// copy file to temp file
	_, err = io.Copy(f, file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// convert docx to pdf
	output, err := exec.Command("unoconv", "-f", "pdf", "--stdout", docxFile).Output()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// remove temp file
	err = os.Remove(docxFile)

	// write pdf to response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Disposition", "attachment; filename="+handler.Filename+".pdf")
	w.Header().Set("Content-Type", "application/pdf")
	w.Write(output)
}

func main() {
	http.HandleFunc("/convert-to-pdf", convert2pdf)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}
