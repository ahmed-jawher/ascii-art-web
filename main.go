package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type PageData struct {
	Result string
	Error  string
	Text   string
}

// Custom error handlers
func handleNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	templates.ExecuteTemplate(w, "404.html", nil)
}

func handleInternalError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	templates.ExecuteTemplate(w, "500.html", nil)
}

func handleBadRequest(w http.ResponseWriter, r *http.Request, message string) {
	w.WriteHeader(http.StatusForbidden)
	templates.ExecuteTemplate(w, "403.html", &PageData{Error: message})
}

var templates *template.Template

func main() {
	templates = template.Must(template.ParseGlob("templates/*.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			handleNotFound(w, r)
			return
		}
		if err := templates.ExecuteTemplate(w, "index.html", &PageData{}); err != nil {
			handleInternalError(w, r)
		}
	})

	http.HandleFunc("/ascii-art", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			handleBadRequest(w, r, "Method not allowed")
			return
		}

		text := r.FormValue("text")
		banner := r.FormValue("banner")

		if text == "" {
			handleBadRequest(w, r, "Text is required")
			return
		}

		bannerFile := banner + ".txt"
		if !isValidBanner(bannerFile) {
			handleBadRequest(w, r, "Invalid banner selected")
			return
		}

		art, err := AsciiArt(text, bannerFile)
		if err != nil {
			if strings.Contains(err.Error(), "unsupported character") {
				handleBadRequest(w, r, err.Error())
				return
			}
			handleInternalError(w, r)
			return
		}

		if err := templates.ExecuteTemplate(w, "index.html", &PageData{
			Result: art,
			Text:   text,
		}); err != nil {
			handleInternalError(w, r)
		}
	})

	fmt.Println("Server running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Server error: %v\n", err)
		os.Exit(1)
	}
}

func isValidBanner(name string) bool {
	validBanners := []string{"standard.txt", "shadow.txt", "thinkertoy.txt"}
	for _, banner := range validBanners {
		if banner == name {
			return true
		}
	}
	return false
}

// AsciiArt generates ASCII art from input text using the specified banner
func AsciiArt(input string, banner string) (string, error) {
	// Handle different types of line endings
	input = strings.ReplaceAll(input, "\r\n", "\n") // Convert Windows line endings
	input = strings.ReplaceAll(input, "\r", "\n")   // Convert old Mac line endings

	bannerPath := filepath.Join("banners", banner)
	fontBytes, err := os.ReadFile(bannerPath)
	if err != nil {
		return "", fmt.Errorf("read font: %w", err)
	}
	fontLines := strings.Split(string(fontBytes), "\n")

	var out strings.Builder
	words := strings.Split(input, "\n")

	for _, word := range words {
		if word == "" {
			out.WriteByte('\n')
			continue
		}

		for row := 0; row < 8; row++ {
			for _, ch := range word {
				if ch < ' ' || ch > '~' {
					return "", fmt.Errorf("unsupported character: %q", ch)
				}
				idx := row + (int(ch-' ') * 9) + 1
				out.WriteString(fontLines[idx])
			}
			out.WriteByte('\n')
		}
	}
	return out.String(), nil
}
