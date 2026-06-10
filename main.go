package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net"
	"net/http"
	"os"

	webview "github.com/webview/webview_go"
)

//go:embed frontend/*
var frontendFiles embed.FS

func main() {
	port := createLocalServer()

	w := webview.New(true)
	defer w.Destroy()

	w.SetTitle("Tasks")
	w.SetSize(1280, 720, 0)

	w.Bind("hello", func() string {
		return "Hello from Webview!"
	})

	w.Navigate(fmt.Sprintf("http://localhost:%d/index.html", port))
	w.Run()
}

func createLocalServer() int {
	subFS, err := fs.Sub(frontendFiles, "frontend")
	if err != nil {
		panic(err)
	}

	http.Handle("/", http.FileServer(http.FS(subFS)))

	// get a free port
	ln, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}
	port := ln.Addr().(*net.TCPAddr).Port
	ln.Close()

	server := &http.Server{Addr: fmt.Sprintf(":%d", port)}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Server error: %s\n", err)
			os.Exit(1)
		}
	}()

	return port
}
