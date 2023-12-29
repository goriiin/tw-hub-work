package httpServer

import (
	"log/slog"
	"twit-hub111/internal/db/postgres"
	"twit-hub111/internal/httpServer/handlers/auth/login"
	"twit-hub111/internal/httpServer/handlers/auth/reg"
	"twit-hub111/internal/httpServer/handlers/news"
	"twit-hub111/internal/httpServer/handlers/profile"
	"twit-hub111/internal/httpServer/handlers/search"
	"twit-hub111/internal/lib/cookies"
)

type HTTPServer struct {
	L *login.LoginService
	R *reg.RegisterService
	N *news.NewsService
	U *profile.UserService
	S *search.SearchService
}

func New(log *slog.Logger, s *postgres.Storage, cache *cookies.CacheService) *HTTPServer {
	return &HTTPServer{
		L: login.New(log, s, cache),
		R: reg.New(log, s, cache),
		N: news.New(log, s, cache),
		U: profile.New(log, s, cache),
		S: search.New(log, s, cache),
	}
}
