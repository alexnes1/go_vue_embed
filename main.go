package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"time"
)

//go:embed front/dist
var frontend embed.FS

func timeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, time.Now())
}

func main() {
	frontendFs, err := fs.Sub(frontend, "front/dist")
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/time", timeHandler)
	http.Handle("/", http.FileServer(http.FS(frontendFs)))

	addr := os.Getenv("ADDR")
	log.Fatal(http.ListenAndServe(addr, nil))
}
