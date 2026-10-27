package main

import (
	"embed"
	"fmt"
	"strings"
)

//go:embed sections/*.md
var sectionFiles embed.FS

// Section is one part of the DDD book.
type Section struct {
	ID    string
	Title string
	Body  string
}

// loadSections reads the embedded files, in book order.
// File names look like "09-entities.md": the ID is "entities".
func loadSections() ([]Section, error) {
	entries, err := sectionFiles.ReadDir("sections")
	if err != nil {
		return nil, err
	}
	var out []Section
	for _, e := range entries {
		data, err := sectionFiles.ReadFile("sections/" + e.Name())
		if err != nil {
			return nil, err
		}
		body := string(data)
		title, _, _ := strings.Cut(body, "\n")
		id := strings.TrimSuffix(e.Name(), ".md")
		if _, rest, ok := strings.Cut(id, "-"); ok {
			id = rest
		}
		out = append(out, Section{ID: id, Title: strings.TrimPrefix(title, "# "), Body: body})
	}
	return out, nil
}

func find(sections []Section, id string) (Section, bool) {
	for _, s := range sections {
		if s.ID == id {
			return s, true
		}
	}
	return Section{}, false
}

// search returns up to limit short matches for a case-insensitive query.
func search(sections []Section, query string, limit int) []string {
	q := strings.ToLower(query)
	var hits []string
	for _, s := range sections {
		lower := strings.ToLower(s.Body)
		if i := strings.Index(lower, q); i >= 0 {
			hits = append(hits, fmt.Sprintf("%s: …%s…", s.ID, snippet(s.Body, i, len(q))))
			if len(hits) == limit {
				break
			}
		}
	}
	return hits
}

func snippet(text string, at, n int) string {
	start, end := max(0, at-80), min(len(text), at+n+80)
	// Stay on rune boundaries.
	for start > 0 && !isRuneStart(text[start]) {
		start--
	}
	for end < len(text) && !isRuneStart(text[end]) {
		end++
	}
	return strings.Join(strings.Fields(text[start:end]), " ")
}

func isRuneStart(b byte) bool { return b&0xC0 != 0x80 }
