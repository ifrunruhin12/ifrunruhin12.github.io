package httpx

import (
	"path"
	"strings"
)

// CanonicalPath normalizes outgoing request paths the same way the router uses them.
func CanonicalPath(raw string) string {
	p := path.Clean("/" + raw)
	if p != "/" && strings.HasSuffix(p, "/") {
		p = strings.TrimSuffix(p, "/")
	}
	return p
}

func IsHomePath(canonicalPath string) bool {
	return canonicalPath == "/" || canonicalPath == "/index.html"
}

func IsAboutPath(canonicalPath string) bool {
	return canonicalPath == "/about" || canonicalPath == "/about.html"
}
