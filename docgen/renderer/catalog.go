package renderer

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Catalog struct {
	Topics   int    `json:"topics"`
	Concepts int    `json:"concepts"`
	Problems int    `json:"problems"`
	Runtime  string `json:"runtime"`
}

func CatalogFromSidebar(categories []SidebarCategory) Catalog {
	out := Catalog{Topics: len(categories), Runtime: "GO"}
	for _, c := range categories {
		out.Concepts += len(c.Concepts)
		out.Problems += len(c.Problems)
	}
	return out
}

func WriteCatalog(outPath string, catalog Catalog) error {
	if err := os.MkdirAll(filepath.Dir(outPath), 0o755); err != nil {
		return err
	}
	b, err := json.MarshalIndent(catalog, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(outPath, append(b, '\n'), 0o644)
}
