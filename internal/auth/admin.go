package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"net/http"
	"strings"

	"clean-portfolio/internal/config"
)

const cookieName = "portfolio_admin"

// Service verifies the admin password and issues an HMAC session cookie (SRP: single responsibility).
type Service struct {
	adminPwd   string
	sessionKey string
}

func NewService(secrets config.AuthSecrets) *Service {
	ap := strings.TrimSpace(secrets.AdminPassword)
	sk := strings.TrimSpace(secrets.SessionKey)
	if sk == "" {
		sk = ap
	}
	return &Service{
		adminPwd:   ap,
		sessionKey: sk,
	}
}

func (s *Service) Enabled() bool { return s.adminPwd != "" }

func constantTimeEq(a, b string) bool {
	ab := []byte(a)
	bb := []byte(b)
	if len(ab) != len(bb) {
		_ = subtle.ConstantTimeCompare(ab, ab)
		return false
	}
	return subtle.ConstantTimeCompare(ab, bb) == 1
}

func (s *Service) PasswordMatches(pw string) bool {
	return constantTimeEq(strings.TrimSpace(pw), strings.TrimSpace(s.adminPwd))
}

func (s *Service) sessionToken() string {
	if !s.Enabled() || s.sessionKey == "" {
		return ""
	}
	mac := hmac.New(sha256.New, []byte(s.sessionKey))
	mac.Write([]byte("portfolio-admin-session-v1|" + strings.TrimSpace(s.adminPwd)))
	return hex.EncodeToString(mac.Sum(nil))
}

func (s *Service) IsAdmin(r *http.Request) bool {
	if !s.Enabled() {
		return false
	}
	got, err := r.Cookie(cookieName)
	if err != nil {
		return false
	}
	return constantTimeEq(strings.TrimSpace(got.Value), s.sessionToken())
}

func (s *Service) SetSessionCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     cookieName,
		Value:    s.sessionToken(),
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   86400 * 30,
	})
}

func (s *Service) ClearSessionCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:   cookieName,
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})
}

func (s *Service) RequireAdmin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !s.IsAdmin(r) {
			http.Error(w, "forbidden", http.StatusForbidden)
			return
		}
		next(w, r)
	}
}
