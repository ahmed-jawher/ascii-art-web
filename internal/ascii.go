package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func IsValidBanner(name string) bool {
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
	input = strings.ReplaceAll(input, "\r\n", "\n")
	input = strings.ReplaceAll(input, "\r", "\n")

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
