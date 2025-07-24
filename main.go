package main

import (
	"ascii-art-stylize/internal"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
)

func main() {
	internal.Templates = template.Must(template.ParseGlob("templates/*.html"))

	// Serve static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			internal.HandleNotFound(w, r)
			return
		}
		if err := internal.Templates.ExecuteTemplate(w, "index.html", &internal.PageData{}); err != nil {
			internal.HandleInternalError(w, r)
		}
	})

	http.HandleFunc("/ascii-art", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			internal.HandleBadRequest(w, r, "Method not allowed")
			return
		}

		text := r.FormValue("text")
		banner := r.FormValue("banner")

		if text == "" {
			internal.HandleBadRequest(w, r, "Text is required")
			return
		}

		bannerFile := banner + ".txt"
		if !internal.IsValidBanner(bannerFile) {
			internal.HandleBadRequest(w, r, "Invalid banner selected")
			return
		}

		art, err := internal.AsciiArt(text, bannerFile)
		if err != nil {
			if strings.Contains(err.Error(), "unsupported character") {
				internal.HandleBadRequest(w, r, err.Error())
				return
			}
			internal.HandleInternalError(w, r)
			return
		}

		if err := internal.Templates.ExecuteTemplate(w, "index.html", &internal.PageData{
			Result: art,
			Text:   text,
		}); err != nil {
			internal.HandleInternalError(w, r)
		}
	})

	fmt.Println("Server running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Server error: %v\n", err)
		os.Exit(1)
	}
}