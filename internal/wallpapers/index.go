package wallpapers

import (
	"log"
	"os"
	"path/filepath"
)

var Wallpapers = make(map[string][]string)

func IndexWallpapers() {
	info, err := os.Stat("./wallpapers")
	if err != nil || !info.IsDir() {
		log.Println("creating wallpapers directory")
		os.Mkdir("./wallpapers", 0755)
		log.Println("successfully created wallpapers directory")
	}

	entries, err := os.ReadDir("./wallpapers")
	if err != nil {
		log.Fatal(err)
	}

	wallpapers := make(map[string][]string)

	for _, entry := range entries {
		if entry.IsDir() {
			path := filepath.Join("./wallpapers", entry.Name())

			subentries, err := os.ReadDir(path)
			if err != nil {
				log.Println(err)
				continue
			}

			for _, wp := range subentries {
				wallpapers[entry.Name()] = append(wallpapers[entry.Name()], wp.Name())
			}
		} else {
			log.Printf("remove %v from the wallpapers directory as it is not a directory", entry.Name())
		}
	}

	log.Printf("wallpapers: %v", wallpapers)

	Wallpapers = wallpapers
}
