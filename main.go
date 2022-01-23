package main

import (
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/smolagakuba/tv-controller/pkg/api"
	"github.com/smolagakuba/tv-controller/pkg/tv"
)

//go:embed static/*
var frontend embed.FS

func main() {
	port := flag.String("port", "", "Name of the rs-232 port to use (ex. /dev/ttyS0)")
	flag.Parse()
	if *port == "" {
		log.Fatal("You must provide name of the port to use (ex. -port=/dev/ttyS0)")
	}
	tv := tv.Init(*port)
	api := api.Init(tv)
	frontend := http.FileServer(getFileSystem())
	r := chi.NewRouter()
	r.Route("/api", func(r chi.Router) {
		r.Handle("/*", api)
	})
	r.Handle("/*", frontend)
	fmt.Println("--------- 🚀 The server is running! 🚀 ----------")
	fmt.Println("• Visit http://localhost:8080 on this machine")
	fmt.Println("• Or http://" + getOutboundIP() + ":8080 on the local network")
	fmt.Println("-------------------------------------------------")
	http.ListenAndServe(":8080", r)
}

func getFileSystem() http.FileSystem {
	fsys, err := fs.Sub(frontend, "static")
	if err != nil {
		log.Fatal(err)
	}
	return http.FS(fsys)
}

func getOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String()
}
