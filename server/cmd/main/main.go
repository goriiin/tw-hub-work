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
	"twit-hub111/internal/httpServer"
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

	appCache := cache.New(-1, 60*time.Minute)
	cacheService := cookies.NewCacheService(appCache, log)

	_ = cacheService
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	fs := http.FileServer(http.Dir("./web/static"))
	router.Handle("/web/static/*", http.StripPrefix("/web/static", fs))

	httpServ := httpServer.New(log, storage, cacheService)

	router.Get("/ru/login", httpServ.L.Login)
	router.Get("/ru/reg", httpServ.R.Reg)
	router.Get("/ru/news", httpServ.N.News)
	router.Get("/ru/search", httpServ.S.Search)

	router.Get("/ru/user/{id}", httpServ.U.User)
	// переход на собственный профиль
	router.Get("/ru/profile", httpServ.U.Profile)
	router.Get("/ru/user/{id}/follow", httpServ.U.Profile)
	router.Get("/ru/user/{id}/is_follow", httpServ.U.Profile)
	router.Get("/ru/user/{id}/unfollow", httpServ.U.Profile)
	router.Get("/ru/profile", httpServ.U.Profile)

	router.Get("/en/login", httpServ.L.Login)
	router.Get("/en/reg", httpServ.R.Reg)
	router.Get("/en/news", httpServ.N.News)
	router.Get("/en/search", httpServ.S.Search)
	router.Get("/en/user/{id}", httpServ.U.User)
	// переход на собственный профиль
	router.Get("/en/profile", httpServ.U.Profile)
	router.Get("/en/user/{id}/follow", httpServ.U.Profile)
	router.Get("/en/user/{id}/is_follow", httpServ.U.Profile)
	router.Get("/en/user/{id}/unfollow", httpServ.U.Profile)

	router.Post("/en/login", httpServ.L.LogData)
	router.Post("/ru/login", httpServ.L.LogData)

	router.Post("/en/reg", httpServ.R.RegData)
	router.Post("/ru/reg", httpServ.R.RegData)

	router.Post("/ru/news", httpServ.N.NewPost)
	router.Post("/en/news", httpServ.N.NewPost)

	router.Get("/search/{nickname}", httpServ.S.SearchNick)
	router.Get("/news/render", httpServ.N.RenderNews)
	router.Get("/profile/render", httpServ.U.RenderNewsProf)
	router.Get("/user/{id}/render", httpServ.U.RenderNews)

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
		if err = srv.ListenAndServe(); err != nil {
			log.Error("failed to start server", err)
		}
	}()

	log.Info("server started")

	<-done
	log.Info("stopping server")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err = srv.Shutdown(ctx); err != nil {
		log.Error("failed to stop server", err)
	}
}
