package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"path/filepath"
	"venv/services/internal/routers"
)

func main() {
	router := routers.NewRouter()

	http.Handle("/static/", http.FileServer(http.Dir(filepath.Join(".", "static"))))

	fmt.Println("Server is running on http://127.0.0.1:8000/")
	log.Fatal(http.ListenAndServe(":8000", router))
}
