package login

import (
	"log/slog"
	"net/http"
	"twit-hub111/internal/db/postgres"
	"twit-hub111/internal/lib/cookies"
)

type LoginService struct {
	log *slog.Logger
	s   *postgres.Storage
	c   *cookies.CacheService
}

func New(
	log *slog.Logger,
	storage *postgres.Storage,
	c *cookies.CacheService,
) *LoginService {
	return &LoginService{
		log: log,
		s:   storage,
		c:   c,
	}
}

func (l *LoginService) Login(w http.ResponseWriter, r *http.Request) {

}
