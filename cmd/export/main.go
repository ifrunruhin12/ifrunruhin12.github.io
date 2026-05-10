package main

import (
	"context"
	"flag"
	"log"
	"os"
	"path/filepath"

	"clean-portfolio/internal/config"
	"clean-portfolio/internal/export"
	"clean-portfolio/internal/paths"
	"clean-portfolio/internal/store"
)

// Run after local edits (server may be stopped; reads the SQLite file directly):
//
//	go run ./cmd/export
//
// Writes plain HTML for GitHub Pages. Commit/push generated index.html and about.html.
func main() {
	outRoot := flag.String("out", ".", "repo root (writes index.html and about.html here)")
	flag.Parse()

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	st, err := store.OpenSQLite(cfg.DBPath)
	if err != nil {
		log.Fatal(err)
	}
	defer st.Close()

	ctx := context.Background()
	if err := st.Ping(ctx); err != nil {
		log.Fatal(err)
	}
	if err := st.Migrate(ctx); err != nil {
		log.Fatal(err)
	}

	if fi, err := os.Stat(*outRoot); err != nil || !fi.IsDir() {
		log.Fatalf("bad -out dir: %q", *outRoot)
	}

	if err := export.WriteGitHubPagesHTML(ctx, st, *outRoot); err != nil {
		log.Fatal(err)
	}

	idx := filepath.Join(*outRoot, paths.ExportedIndexPath)
	abt := filepath.Join(*outRoot, paths.ExportedAboutPath)
	log.Printf("wrote %s and %s — commit/push those for GitHub Pages", idx, abt)
}
