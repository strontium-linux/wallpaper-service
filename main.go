package main

import (
	"flag"
	"log"
	"net/http"
	"time"
	"wallpaper-service/internal/server"
	"wallpaper-service/internal/wallpapers"
)

func main() {
	log.Println("wallpaper service starting")

	wallpapers.IndexWallpapers()

	ticker := time.NewTicker(4 * time.Hour)
	defer ticker.Stop()

	go func() {
		for t := range ticker.C {
			log.Printf("time of index: %v", t.UnixMilli())
			wallpapers.IndexWallpapers()
		}
	}()

	port := flag.String("port", ":8080", "Port to serve on")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/wallpapers/get-wallpapers-metadata", server.GetWallpapersMetadata)
	mux.HandleFunc("/wallpapers/get-wallpaper", server.GetWallpaper)
	log.Printf("serving http server on port %v", *port)
	http.ListenAndServe(*port, mux)
}
