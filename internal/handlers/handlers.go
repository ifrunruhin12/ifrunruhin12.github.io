// Package handlers contains HTTP adapters; it depends on store.Repository via interfaces (DIP).
package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"clean-portfolio/internal/auth"
	"clean-portfolio/internal/httpx"
	"clean-portfolio/internal/models"
	"clean-portfolio/internal/paths"
	"clean-portfolio/internal/store"
	"clean-portfolio/internal/view"
)

type Handler struct {
	Store store.Repository
	Auth  *auth.Service
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	if !httpx.IsHomePath(httpx.CanonicalPath(r.URL.Path)) {
		http.NotFound(w, r)
		return
	}

	ctx := r.Context()
	site, err := h.Store.GetSite(ctx)
	if err != nil {
		http.Error(w, "db error", http.StatusInternalServerError)
		return
	}
	projects, err := h.Store.ListProjects(ctx)
	if err != nil {
		http.Error(w, "db error", http.StatusInternalServerError)
		return
	}
	view.RenderHTML(w, paths.RenderHome, models.PageData{
		Site:     site,
		Projects: projects,
		IsAdmin:  h.Auth.IsAdmin(r),
	})
}

func (h *Handler) About(w http.ResponseWriter, r *http.Request) {
	if !httpx.IsAboutPath(httpx.CanonicalPath(r.URL.Path)) {
		http.NotFound(w, r)
		return
	}
	ctx := r.Context()
	site, err := h.Store.GetSite(ctx)
	if err != nil {
		http.Error(w, "db error", http.StatusInternalServerError)
		return
	}
	view.RenderHTML(w, paths.RenderAbout, models.PageData{
		Site:    site,
		IsAdmin: h.Auth.IsAdmin(r),
	})
}

func (h *Handler) AdminLoginPage(w http.ResponseWriter, r *http.Request) {
	if h.Auth.IsAdmin(r) {
		http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
		return
	}
	ctx := r.Context()
	site, err := h.Store.GetSite(ctx)
	if err != nil {
		http.Error(w, "db error", http.StatusInternalServerError)
		return
	}
	view.RenderHTML(w, paths.AdminLogin, models.PageData{Site: site})
}

func (h *Handler) AdminDashboard(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	site, err := h.Store.GetSite(ctx)
	if err != nil {
		http.Error(w, "db error", http.StatusInternalServerError)
		return
	}
	projects, err := h.Store.ListProjects(ctx)
	if err != nil {
		http.Error(w, "db error", http.StatusInternalServerError)
		return
	}
	view.RenderHTML(w, paths.AdminDashboard, models.PageData{
		Site:     site,
		Projects: projects,
		IsAdmin:  true,
	})
}

func (h *Handler) AdminLoginPost(w http.ResponseWriter, r *http.Request) {
	if !h.Auth.Enabled() {
		http.Error(w, "admin disabled (set ADMIN_PASSWORD)", http.StatusForbidden)
		return
	}
	if !h.Auth.PasswordMatches(r.FormValue("password")) {
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
		return
	}
	h.Auth.SetSessionCookie(w)
	http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
}

func (h *Handler) AdminLogout(w http.ResponseWriter, r *http.Request) {
	h.Auth.ClearSessionCookie(w)
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

func (h *Handler) AddProject(w http.ResponseWriter, r *http.Request) {
	title := strings.TrimSpace(r.FormValue("title"))
	if title == "" {
		http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
		return
	}
	desc := strings.TrimSpace(r.FormValue("description"))
	gh := strings.TrimSpace(r.FormValue("github"))
	if err := h.Store.InsertProject(r.Context(), title, desc, gh); err != nil {
		http.Error(w, "db error", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
}

func (h *Handler) UpdateProject(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimSpace(r.FormValue("id")))
	if err != nil || id <= 0 {
		http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
		return
	}
	title := strings.TrimSpace(r.FormValue("title"))
	if title == "" {
		http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
		return
	}
	desc := strings.TrimSpace(r.FormValue("description"))
	gh := strings.TrimSpace(r.FormValue("github"))
	if err := h.Store.UpdateProject(r.Context(), id, title, desc, gh); err != nil {
		http.Error(w, "db error", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
}

func (h *Handler) DeleteProjectPath(requestPath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := strings.TrimPrefix(requestPath, "/delete/")
		idStr = strings.TrimSpace(idStr)
		id, err := strconv.Atoi(idStr)
		if err != nil || id <= 0 || strings.Contains(idStr, "/") {
			http.NotFound(w, r)
			return
		}
		if err := h.Store.DeleteProject(r.Context(), id); err != nil {
			http.Error(w, "db error", http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
	}
}

func (h *Handler) UpdateSite(w http.ResponseWriter, r *http.Request) {
	s := models.Site{
		SiteTitle:   strings.TrimSpace(r.FormValue("site_title")),
		DisplayName: strings.TrimSpace(r.FormValue("display_name")),
		Intro:       strings.TrimSpace(r.FormValue("intro")),
		AboutBody:   strings.TrimSpace(r.FormValue("about_body")),
		GithubURL:   strings.TrimSpace(r.FormValue("github_url")),
		LinkedinURL: strings.TrimSpace(r.FormValue("linkedin_url")),
		EmailMailto: strings.TrimSpace(r.FormValue("email_mailto")),
		LogoURL:     strings.TrimSpace(r.FormValue("logo_url")),
	}
	if s.SiteTitle == "" {
		s.SiteTitle = "Portfolio"
	}
	if s.DisplayName == "" {
		s.DisplayName = s.SiteTitle
	}
	if s.LogoURL == "" {
		s.LogoURL = models.DefaultLogoURL
	}
	if err := h.Store.UpdateSite(r.Context(), s); err != nil {
		http.Error(w, "db error", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
}
