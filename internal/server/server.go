package server

import (
	"context"
	"log"
	"net/http"
	"strings"

	"clean-portfolio/internal/auth"
	"clean-portfolio/internal/config"
	"clean-portfolio/internal/handlers"
	"clean-portfolio/internal/httpx"
	"clean-portfolio/internal/store"
)

func Run() error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	st, err := store.OpenSQLite(cfg.DBPath)
	if err != nil {
		return err
	}
	defer st.Close()

	ctx := context.Background()
	if err := st.Ping(ctx); err != nil {
		return err
	}
	if err := st.Migrate(ctx); err != nil {
		return err
	}

	h := &handlers.Handler{
		Store: st,
		Auth:  auth.NewService(config.LoadAuthSecrets()),
	}

	staticFS := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
	uploadFS := http.StripPrefix("/uploads/", http.FileServer(http.Dir("uploads")))

	main := func(w http.ResponseWriter, r *http.Request) {
		p := httpx.CanonicalPath(r.URL.Path)

		switch {
		case strings.HasPrefix(p+"/", "/static/"):
			staticFS.ServeHTTP(w, r)
			return
		case strings.HasPrefix(p+"/", "/uploads/"):
			uploadFS.ServeHTTP(w, r)
			return

		case httpx.IsHomePath(p) && r.Method == http.MethodGet:
			h.Home(w, r)
		case httpx.IsAboutPath(p) && r.Method == http.MethodGet:
			h.About(w, r)
		case p == "/admin" && r.Method == http.MethodGet:
			h.AdminLoginPage(w, r)
		case p == "/admin/dashboard" && r.Method == http.MethodGet:
			h.Auth.RequireAdmin(h.AdminDashboard)(w, r)
		case p == "/admin/login" && r.Method == http.MethodPost:
			h.AdminLoginPost(w, r)
		case p == "/admin/logout" && r.Method == http.MethodPost:
			h.AdminLogout(w, r)
		case p == "/add-project" && r.Method == http.MethodPost:
			h.Auth.RequireAdmin(h.AddProject)(w, r)
		case p == "/update-project" && r.Method == http.MethodPost:
			h.Auth.RequireAdmin(h.UpdateProject)(w, r)
		case p == "/update-site" && r.Method == http.MethodPost:
			h.Auth.RequireAdmin(h.UpdateSite)(w, r)
		case strings.HasPrefix(p, "/delete/") && r.Method == http.MethodPost:
			h.Auth.RequireAdmin(h.DeleteProjectPath(p))(w, r)
		default:
			http.NotFound(w, r)
		}
	}

	log.Printf("listening on http://localhost%s (set ADMIN_PASSWORD to enable admin)", cfg.ListenAddr)
	return http.ListenAndServe(cfg.ListenAddr, http.HandlerFunc(main))
}
