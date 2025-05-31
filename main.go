package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
)

type PageData struct {
	Result string
	Error  string
}

func main() {
	templates := template.Must(template.ParseGlob("templates/*.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		templates.ExecuteTemplate(w, "index.html", &PageData{})
	})

	http.HandleFunc("/ascii-art", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		text := r.FormValue("text")
		banner := r.FormValue("banner")

		if text == "" {
			templates.ExecuteTemplate(w, "index.html", &PageData{Error: "Text is required"})
			return
		}

		// Choose banner file based on selection
		bannerFile := banner + ".txt"
		if !isValidBanner(bannerFile) {
			templates.ExecuteTemplate(w, "index.html", &PageData{Error: "Invalid banner selected"})
			return
		}

		art, err := AsciiArt(text, bannerFile)
		if err != nil {
			templates.ExecuteTemplate(w, "index.html", &PageData{Error: err.Error()})
			return
		}

		templates.ExecuteTemplate(w, "index.html", &PageData{Result: art})
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
	

	fontBytes, err := os.ReadFile(banner)
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
