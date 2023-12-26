package search

import (
	"log/slog"
	"net/http"
	"twit-hub111/internal/db/postgres"
	"twit-hub111/internal/lib/cookies"
)

type SearchService struct {
	log *slog.Logger
	s   *postgres.Storage
	c   *cookies.CacheService
}

func New(
	log *slog.Logger,
	storage *postgres.Storage,
	c *cookies.CacheService,
) *SearchService {
	return &SearchService{
		log: log,
		s:   storage,
		c:   c,
	}
}

func (s *SearchService) Search(w http.ResponseWriter, r *http.Request) {

}
