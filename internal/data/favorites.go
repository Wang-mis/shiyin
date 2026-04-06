package data

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func favoritesPath() (string, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "shiyin", "favorites.json"), nil
}

func LoadFavorites() ([]Poem, error) {
	path, err := favoritesPath()
	if err != nil {
		return nil, err
	}
	f, err := os.Open(path)
	if os.IsNotExist(err) {
		return []Poem{}, nil
	}
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var poems []Poem
	if err := json.NewDecoder(f).Decode(&poems); err != nil {
		return nil, err
	}
	return poems, nil
}

func SaveFavorites(poems []Poem) error {
	path, err := favoritesPath()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return json.NewEncoder(f).Encode(poems)
}

func IsFavorite(poems []Poem, title, author string) bool {
	for _, p := range poems {
		if p.Title == title && p.Author == author {
			return true
		}
	}
	return false
}

// ToggleFavorite adds or removes poem from favorites.
// Returns the updated list and true if added, false if removed.
func ToggleFavorite(favorites []Poem, poem Poem) ([]Poem, bool) {
	for i, p := range favorites {
		if p.Title == poem.Title && p.Author == poem.Author {
			return append(favorites[:i:i], favorites[i+1:]...), false
		}
	}
	return append(favorites, poem), true
}
