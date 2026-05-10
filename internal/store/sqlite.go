package store

import (
	"context"
	"database/sql"
	"strings"

	"clean-portfolio/internal/models"

	_ "github.com/mattn/go-sqlite3"
)

// SQLiteStore is the concrete SQLite implementation of Repository.
type SQLiteStore struct {
	db *sql.DB
}

func OpenSQLite(dsn string) (*SQLiteStore, error) {
	conn, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}
	return &SQLiteStore{db: conn}, nil
}

func (s *SQLiteStore) Close() error { return s.db.Close() }

func (s *SQLiteStore) Ping(ctx context.Context) error {
	return s.db.PingContext(ctx)
}

func (s *SQLiteStore) Migrate(ctx context.Context) error {
	stmts := []string{
		`CREATE TABLE IF NOT EXISTS projects (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			description TEXT,
			github TEXT,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS site_settings (
			id INTEGER PRIMARY KEY CHECK (id = 1),
			site_title TEXT NOT NULL DEFAULT 'popcycle',
			display_name TEXT NOT NULL DEFAULT 'popcycle',
			intro TEXT NOT NULL DEFAULT '',
			about_body TEXT NOT NULL DEFAULT '',
			github_url TEXT NOT NULL DEFAULT '',
			linkedin_url TEXT NOT NULL DEFAULT '',
			email_mailto TEXT NOT NULL DEFAULT '',
			logo_url TEXT NOT NULL DEFAULT '/uploads/rocket.png'
		)`,
	}
	for _, q := range stmts {
		if _, err := s.db.ExecContext(ctx, q); err != nil {
			return err
		}
	}
	if err := s.migrateLegacyLogoColumn(ctx); err != nil {
		return err
	}

	var n int
	if err := s.db.QueryRowContext(ctx, `SELECT COUNT(*) FROM site_settings`).Scan(&n); err != nil {
		return err
	}
	if n == 0 {
		intro := "Hi there.\n\nI'm popcycle.\nBackend developer focused on Golang, Linux and systems programming."
		about := "I build small, reliable systems and minimalist software.\n\nThis site is intentionally plain."
		_, err := s.db.ExecContext(ctx,
			`INSERT INTO site_settings (id, site_title, display_name, intro, about_body) VALUES (1, ?, ?, ?, ?)`,
			"popcycle", "popcycle", intro, about)
		return err
	}
	return nil
}

func (s *SQLiteStore) migrateLegacyLogoColumn(ctx context.Context) error {
	var n int
	err := s.db.QueryRowContext(ctx,
		`SELECT COUNT(*) FROM pragma_table_info('site_settings') WHERE name = 'logo_url'`).Scan(&n)
	if err != nil {
		return err
	}
	if n > 0 {
		return nil
	}
	_, err = s.db.ExecContext(ctx,
		`ALTER TABLE site_settings ADD COLUMN logo_url TEXT NOT NULL DEFAULT '`+
			strings.ReplaceAll(models.DefaultLogoURL, "'", "''")+`'`)
	return err
}

func normalizeLogoURL(raw string) string {
	v := strings.TrimSpace(raw)
	if v == "" {
		return models.DefaultLogoURL
	}
	return v
}

func (s *SQLiteStore) GetSite(ctx context.Context) (models.Site, error) {
	row := s.db.QueryRowContext(ctx, `
		SELECT site_title, display_name, intro, about_body, github_url, linkedin_url, email_mailto, logo_url
		FROM site_settings WHERE id = 1
	`)
	var site models.Site
	if err := row.Scan(&site.SiteTitle, &site.DisplayName, &site.Intro, &site.AboutBody,
		&site.GithubURL, &site.LinkedinURL, &site.EmailMailto, &site.LogoURL); err != nil {
		return site, err
	}
	site.LogoURL = normalizeLogoURL(site.LogoURL)
	return site, nil
}

func (s *SQLiteStore) UpdateSite(ctx context.Context, site models.Site) error {
	logoURL := normalizeLogoURL(site.LogoURL)

	_, err := s.db.ExecContext(ctx, `
		UPDATE site_settings SET
			site_title = ?,
			display_name = ?,
			intro = ?,
			about_body = ?,
			github_url = ?,
			linkedin_url = ?,
			email_mailto = ?,
			logo_url = ?
		WHERE id = 1`,
		strings.TrimSpace(site.SiteTitle),
		strings.TrimSpace(site.DisplayName),
		strings.TrimSpace(site.Intro),
		strings.TrimSpace(site.AboutBody),
		strings.TrimSpace(site.GithubURL),
		strings.TrimSpace(site.LinkedinURL),
		strings.TrimSpace(site.EmailMailto),
		logoURL,
	)
	return err
}

func (s *SQLiteStore) ListProjects(ctx context.Context) ([]models.Project, error) {
	rows, err := s.db.QueryContext(ctx, `
		SELECT id, title, description, github FROM projects ORDER BY id DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []models.Project
	for rows.Next() {
		var p models.Project
		if err := rows.Scan(&p.ID, &p.Title, &p.Description, &p.Github); err != nil {
			return nil, err
		}
		out = append(out, p)
	}
	return out, rows.Err()
}

func (s *SQLiteStore) InsertProject(ctx context.Context, title, description, github string) error {
	_, err := s.db.ExecContext(ctx,
		`INSERT INTO projects(title, description, github) VALUES (?, ?, ?)`,
		strings.TrimSpace(title), strings.TrimSpace(description), strings.TrimSpace(github))
	return err
}

func (s *SQLiteStore) UpdateProject(ctx context.Context, id int, title, description, github string) error {
	_, err := s.db.ExecContext(ctx,
		`UPDATE projects SET title = ?, description = ?, github = ? WHERE id = ?`,
		strings.TrimSpace(title), strings.TrimSpace(description), strings.TrimSpace(github), id)
	return err
}

func (s *SQLiteStore) DeleteProject(ctx context.Context, id int) error {
	_, err := s.db.ExecContext(ctx, `DELETE FROM projects WHERE id = ?`, id)
	return err
}
