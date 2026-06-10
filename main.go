package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/diogocorigo/tasks/backend/controllers"
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

	/**
	 * Expose all the controller functinos into the WebView.
	 * This is the only way to call them from the frontend.
	 */
	w.Bind("getTasks", controllers.GetTasks)
	w.Bind("createTask", controllers.CreateTask)

	w.Bind("hello", func() string {
		return "Hello from Webview!"
	})

	w.Navigate(fmt.Sprintf("http://localhost:%d/", port))
	w.Run()
}

/**
 * Create a local HTTP server to serve the frontend files.
 * If the `APP_DEBUG` environment variable is set to `true`, the server will
 * proxy requests to Vite's dev server. Otherwise, it will serve the embedded
 * frontend files.
 */
func createLocalServer() int {
	var handler http.Handler

	DEBUG := os.Getenv("APP_DEBUG")

	if DEBUG == "true" {
		viteURL, _ := url.Parse("http://localhost:5173")
		handler = httputil.NewSingleHostReverseProxy(viteURL)
	} else {
		subFS, err := fs.Sub(frontendFiles, "frontend/dist")
		if err != nil {
			panic(err)
		}
		handler = http.FileServer(http.FS(subFS))
	}

	http.Handle("/", handler)

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
