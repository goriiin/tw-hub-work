package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"text/template"
	"twit-hub111/internal/config"
)

//https://qna.habr.com/q/915835

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func news(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("web/news/newsFeed.html"))

	fmt.Println("Rendering news template")
	err := temp.ExecuteTemplate(w, "body", nil)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	//w.Header().Set("Content-Type", "text/html")
	//http.ServeFile(w, r, "web/news/newsFeed.html")
}

func main() {

	cfg := config.MustLoad()

	// TODO: init logger: slog
	log := setupLogger(cfg.Env)

	log.Info("starting...", slog.String("env", cfg.Env))
	log.Debug("debug enabled")

	// TODO: init storage: postgres

	// TODO: init router: chi, "chi render"

	//TODO: init server

	//db.Test()
	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/news"))))
	//http.HandleFunc("/news", news)
	//http.ListenAndServe(":8080", nil)
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(
			os.Stdout,
			&slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}

func setupENV() {
	err := os.Setenv("CONFIG_PATH", "./config/local.yaml")
	if err != nil {
		fmt.Println("err: ", err)
	}
}
