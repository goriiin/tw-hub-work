package user

import (
	"context"
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"twit-hub111/internal/domain"
)

type User interface {
	InsertUser(ctx context.Context, u *domain.User) (int, error)
}

func New(log *slog.Logger, user User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func Users(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("./web/static/profile/profile.html"))

	fmt.Println("Rendering news template")
	err := temp.ExecuteTemplate(w, "body", nil)
	if err != nil {
		_, _ = fmt.Fprintf(w, err.Error())
	}
}
