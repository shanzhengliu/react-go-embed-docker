package main

import (
	"embed"
	"io/fs"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
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

	router.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	router.PathPrefix("/").Handler(http.FileServer(http.FS(ui)))

	cor := cors.New(cors.Options{
		AllowedOrigins:   []string{origin},
		AllowCredentials: false,
		AllowedHeaders:   []string{"*"},
	})
	corHandler := cor.Handler(router)

	http.ListenAndServe(":"+port, corHandler)
}
