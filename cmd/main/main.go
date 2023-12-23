package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"twit-hub111/internal/config"
	"twit-hub111/internal/db/postgres"
	"twit-hub111/internal/http-server/handlers/news"
)

//https://qna.habr.com/q/915835

func main() {
	config.SetupENV()
	cfg := config.MustLoad()

	log := config.SetupLogger(cfg.Env)

	log.Info("starting...", slog.String("env", cfg.Env))
	log.Debug("debug enabled")

	storage, err := postgres.New(cfg.DbConfigPath)
	fmt.Println(cfg.DbConfigPath)
	if err != nil {
		log.Error("db error", err)
		os.Exit(1)
	}

	//storage.DropDB()
	//err = storage.SetDB()
	//if err != nil {
	//	log.Error("set db tables error", err)
	//	os.Exit(1)
	//}

	// TODO: сделать время для поста
	log.Info("DB started")

	err = storage.TestSelect()
	if err != nil {
		log.Error("database tables have not been created", err)
	}

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Get("/news", news.News)
	router.Post("/news", news.NewPost)

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

	// TODO: move timeout to config
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Error("failed to stop server")

		return
	}
}
