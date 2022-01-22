package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"

	"github.com/smolagakuba/tv-controller/tv"
)

//go:embed static/*
var frontend embed.FS

func main() {
	s := http.FileServer(getFileSystem())
	http.ListenAndServe(":8080", s)
	tv.Init()
}

func getFileSystem() http.FileSystem {
	fsys, err := fs.Sub(frontend, "static")
	if err != nil {
		log.Fatal(err)
	}
	return http.FS(fsys)
}
