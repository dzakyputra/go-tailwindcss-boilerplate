package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

//go:embed static/*.html
var resources embed.FS

var t = template.Must(template.ParseFS(resources, "static/*.html"))

func main() {
	port := getPort()

	http.HandleFunc("/", index)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	log.Println("listening on", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func index(w http.ResponseWriter, req *http.Request) {
	host := getHost(req)
	data := make(map[string]interface{})
	data["Host"] = host
	t.ExecuteTemplate(w, "index.html", data)
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}

func getHost(req *http.Request) string {
	host := "https://" + req.Host
	if strings.Contains(req.Host, "localhost") {
		host = "http://" + req.Host
	}
	return host
}
