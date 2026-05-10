package view

import (
	"html/template"
	"io"
	"log"
	"net/http"

	"clean-portfolio/internal/models"
)

// Execute writes a parsed template path to w (shared by HTTP handlers and static export).
func Execute(w io.Writer, tmplPath string, data models.PageData) error {
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		return err
	}
	return tmpl.Execute(w, data)
}

func RenderHTML(res http.ResponseWriter, tmplPath string, data models.PageData) {
	if err := Execute(res, tmplPath, data); err != nil {
		http.Error(res, "template error", http.StatusInternalServerError)
		log.Println(err)
	}
}
