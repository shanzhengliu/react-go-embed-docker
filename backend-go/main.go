package main

import (
	"embed"
	"io/fs"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

//go:embed ui
var UI embed.FS

const defaultPort = "8080"
const defaultOrigin = "*"

func main() {
	port := defaultPort
	origin := defaultOrigin
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	if os.Getenv("ORIGIN") != "" {
		origin = os.Getenv("PORT")
	}

	ui, _ := fs.Sub(UI, "ui")

	router := mux.NewRouter()

	credentials := handlers.AllowCredentials()
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{origin})
	router.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	router.Path("/").Handler(http.FileServer(http.FS(ui)))

	http.ListenAndServe(":"+port, handlers.CORS(credentials, methods, origins)(router))
}
