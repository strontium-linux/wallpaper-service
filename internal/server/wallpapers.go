package server

import (
	"encoding/json"
	"net/http"
	"path/filepath"
	"strings"
	"wallpaper-service/internal/wallpapers"
)

func GetWallpapersMetadata(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	wallpapers := wallpapers.Wallpapers

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(wallpapers); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetWallpaper(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	path := r.URL.Query().Get("path")

	fullpath := filepath.Join("./wallpapers", path)

	base, _ := filepath.Abs("./wallpapers")
	full, _ := filepath.Abs(fullpath)
	if !strings.HasPrefix(full, base+string(filepath.Separator)) {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	http.ServeFile(w, r, full)
}
