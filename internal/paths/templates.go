package paths

// RenderHome and RenderAbout are Go html/template sources used by the local server and exporter.
const (
	RenderHome  = "templates/render/index.gohtml"
	RenderAbout = "templates/render/about.gohtml"

	// ExportedIndexPath and ExportedAboutPath are files written by cmd/export for GitHub Pages (plain HTML).
	ExportedIndexPath = "index.html"
	ExportedAboutPath = "about.html"

	AdminLogin     = "templates/admin_login.html"
	AdminDashboard = "templates/admin.html"
)
