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

	router.Get("/ru/news", news.New(log, storage, cacheService).News)
	router.Get("/news", news.New(log, storage, cacheService).News)
	router.Post("/news", news.New(log, storage, cacheService).NewPost)
	router.Get("/ru/search", search.Search)
	router.Get("/search", search.Search)
	router.Get("/profile", profile.Users)
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
