package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/patrickmn/go-cache"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"twit-hub111/internal/config"
	"twit-hub111/internal/db/postgres"
	"twit-hub111/internal/http-server/handlers/auth/login"
	"twit-hub111/internal/http-server/handlers/auth/reg"
	"twit-hub111/internal/http-server/handlers/news"
	"twit-hub111/internal/http-server/handlers/profile"
	"twit-hub111/internal/http-server/handlers/search"
	"twit-hub111/internal/lib/cookies"
)

//https://qna.habr.com/q/915835

func main() {
	config.SetupENV()
	cfg := config.MustLoad()

	fmt.Println(cfg)
	log := config.SetupLogger(cfg.Env)

	log.Info("starting...", slog.String("env", cfg.Env))
	log.Debug("debug enabled")

	storage, err := postgres.New(cfg.DbConfigPath)
	fmt.Println(cfg.DbConfigPath)
	if err != nil {
		log.Error("db error", err)
		os.Exit(1)
	}

	log.Info("DB started")

	err = storage.TestSelect()
	if err != nil {
		log.Error("database tables have not been created", err)
	}

	appCache := cache.New(-1, 60*time.Minute)
	cacheService := cookies.NewCacheService(appCache, log)

	_ = cacheService
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Get("/ru/login", login.New(log, storage, cacheService).Login)
	router.Get("/ru/reg", reg.New(log, storage, cacheService).Reg)
	router.Get("/ru/news", news.New(log, storage, cacheService).News)
	router.Get("/ru/search", search.New(log, storage, cacheService).Search)
	router.Get("/ru/search/{nickname}", search.New(log, storage, cacheService).SearchNick)
	router.Get("/ru/user/{id}", profile.New(log, storage, cacheService).Users)

	router.Get("/en/login", login.New(log, storage, cacheService).Login)
	router.Get("/en/reg", reg.New(log, storage, cacheService).Reg)
	router.Get("/en/news", news.New(log, storage, cacheService).News)
	router.Get("/en/search", search.New(log, storage, cacheService).Search)
	router.Get("/en/search/{nickname}", search.New(log, storage, cacheService).SearchNick)
	router.Get("/en/user/{id}", profile.New(log, storage, cacheService).Users)

	router.Post("/en/login", login.New(log, storage, cacheService).LogData)
	router.Post("/ru/login", login.New(log, storage, cacheService).LogData)

	router.Post("/en/reg", reg.New(log, storage, cacheService).RegData)
	router.Post("/ru/reg", reg.New(log, storage, cacheService).RegData)

	router.Post("/news", news.New(log, storage, cacheService).NewPost)

	log.Info("starting server", slog.String("address", cfg.Address))

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Error("failed to start server")
		}
	}()

	log.Info("server started")

	<-done
	log.Info("stopping server")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Error("failed to stop server")
	}
}
