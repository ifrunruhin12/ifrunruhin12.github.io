// Package models defines read models used by handlers and persistence.
package models

const DefaultLogoURL = "/uploads/rocket.png"

type Project struct {
	ID          int
	Title       string
	Description string
	Github      string
}

type Site struct {
	SiteTitle   string
	DisplayName string
	Intro       string
	AboutBody   string
	GithubURL   string
	LinkedinURL string
	EmailMailto string
	LogoURL     string
}

type PageData struct {
	Site     Site
	Projects []Project
	IsAdmin  bool
}
