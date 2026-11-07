package renderer

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type SearchEntry struct {
	ID       string   `json:"id"`
	Title    string   `json:"title"`
	Type     string   `json:"type"`
	Category string   `json:"category"`
	Tags     []string `json:"tags"`
}

func WriteSearchIndex(outPath string, entries []SearchEntry) error {
	if entries == nil {
		entries = []SearchEntry{}
	}
	if err := os.MkdirAll(filepath.Dir(outPath), 0o755); err != nil {
		return err
	}
	b, err := json.MarshalIndent(entries, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(outPath, append(b, '\n'), 0o644)
}
