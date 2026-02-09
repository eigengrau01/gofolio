package main

import (
	"embed"
	"log"
	"net/http"
	"github.com/eigengrau01/gofolio/components"

	"github.com/a-h/templ"
)

//go:embed static/*
var static embed.FS

//go:generate npx tailwindcss build -i static/css/style.css -o static/css/tailwind.css -m

func main() {
	homePage := components.Maingrid()
	pagesHandler := http.NewServeMux()
	pagesHandler.Handle("/", templ.Handler(homePage))
	pagesHandler.Handle("/static/", http.FileServer(http.FS(static)))

	log.Fatalln(http.ListenAndServe(":8080", pagesHandler))
}
