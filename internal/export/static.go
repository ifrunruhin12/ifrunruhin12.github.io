package export

import (
	"context"
	"os"
	"path/filepath"

	"clean-portfolio/internal/models"
	"clean-portfolio/internal/paths"
	"clean-portfolio/internal/store"
	"clean-portfolio/internal/view"
)

// WriteGitHubPagesHTML renders the portfolio from SQLite into root index.html and about.html (public; no admin chrome).
func WriteGitHubPagesHTML(ctx context.Context, repo store.Repository, rootDir string) error {
	site, err := repo.GetSite(ctx)
	if err != nil {
		return err
	}
	projects, err := repo.ListProjects(ctx)
	if err != nil {
		return err
	}
	pub := models.PageData{
		Site:     site,
		Projects: projects,
		IsAdmin:  false,
	}

	writeOne := func(tmplRel, htmlName string) error {
		dest := filepath.Join(rootDir, htmlName)
		f, err := os.Create(dest)
		if err != nil {
			return err
		}
		err = view.Execute(f, tmplRel, pub)
		_ = f.Close()
		return err
	}

	if err := writeOne(paths.RenderHome, paths.ExportedIndexPath); err != nil {
		return err
	}
	return writeOne(paths.RenderAbout, paths.ExportedAboutPath)
}
