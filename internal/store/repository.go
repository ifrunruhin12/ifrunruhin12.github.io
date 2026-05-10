package store

import (
	"context"

	"clean-portfolio/internal/models"
)

// Repository abstracts persistence so handlers depend on behaviour, not SQLite (DIP).
type Repository interface {
	Migrate(ctx context.Context) error

	GetSite(ctx context.Context) (models.Site, error)
	UpdateSite(ctx context.Context, site models.Site) error

	ListProjects(ctx context.Context) ([]models.Project, error)
	InsertProject(ctx context.Context, title, description, github string) error
	UpdateProject(ctx context.Context, id int, title, description, github string) error
	DeleteProject(ctx context.Context, id int) error
}
