package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Contact struct {
	ID           string   `json:"id"`
	Service      string   `json:"service"`
	Category     string   `json:"category"`
	Numbers      []string `json:"numbers"`
	Availability string   `json:"availability"`
	Status       string   `json:"status"`
	SourceID     string   `json:"sourceId"`
	SourceURL    string   `json:"sourceUrl"`
	CheckedAt    string   `json:"checkedAt"`
}
type Dataset struct {
	Version    string    `json:"version"`
	ReviewedAt string    `json:"reviewedAt"`
	Warning    string    `json:"warning"`
	Contacts   []Contact `json:"contacts"`
}

func main() {
	raw, e := os.ReadFile("data/contacts.json")
	if e != nil {
		log.Fatal(e)
	}
	var d Dataset
	if e = json.Unmarshal(raw, &d); e != nil {
		log.Fatal(e)
	}
	origin := os.Getenv("ALLOWED_ORIGIN")
	if origin == "" {
		origin = "https://essential.digitalghana.dev"
	}
	write := func(w http.ResponseWriter, v any) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_ = json.NewEncoder(w).Encode(v)
	}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		write(w, map[string]any{"status": "ok", "version": d.Version, "contacts": len(d.Contacts), "reviewedAt": d.ReviewedAt})
	})
	mux.HandleFunc("GET /v1/contacts", func(w http.ResponseWriter, r *http.Request) { write(w, d) })
	mux.HandleFunc("GET /v1/offline.json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Disposition", "attachment; filename=ghanaessential-"+d.Version+".json")
		write(w, d)
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+port, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Content-Type-Options", "nosniff")
		if r.Header.Get("Origin") == origin {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
		}
		mux.ServeHTTP(w, r)
	})))
}
