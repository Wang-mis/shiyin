package data

import (
	"encoding/json"
	"embed"
	"fmt"
)

//go:embed tang300.json ci300.json
var dataFS embed.FS

// AvailableCollections lists all built-in collections with their display names and file keys.
var AvailableCollections = []struct {
	Name string
	Key  string
}{
	{Name: "唐诗三百首", Key: "tang300"},
	{Name: "宋词三百首", Key: "ci300"},
}

// Load loads poems from the embedded data by key ("tang300", "ci300", or "all").
func Load(key string) ([]Poem, error) {
	switch key {
	case "tang300":
		return loadFile("tang300.json")
	case "ci300":
		return loadFile("ci300.json")
	case "all":
		tang, err := loadFile("tang300.json")
		if err != nil {
			return nil, err
		}
		ci, err := loadFile("ci300.json")
		if err != nil {
			return nil, err
		}
		return append(tang, ci...), nil
	default:
		return nil, fmt.Errorf("unknown collection: %s", key)
	}
}

// LoadByName loads poems by display name (e.g. "唐诗三百首").
func LoadByName(name string) ([]Poem, error) {
	for _, c := range AvailableCollections {
		if c.Name == name {
			return Load(c.Key)
		}
	}
	return nil, fmt.Errorf("collection not found: %s", name)
}

func loadFile(path string) ([]Poem, error) {
	b, err := dataFS.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", path, err)
	}
	var poems []Poem
	if err := json.Unmarshal(b, &poems); err != nil {
		return nil, fmt.Errorf("parse %s: %w", path, err)
	}
	return poems, nil
}
