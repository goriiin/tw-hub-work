package reg

import (
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"twit-hub111/internal/db/postgres"
	"twit-hub111/internal/lib/cookies"
)

type RegisterService struct {
	log *slog.Logger
	s   *postgres.Storage
	c   *cookies.CacheService
}

func New(
	log *slog.Logger,
	storage *postgres.Storage,
	c *cookies.CacheService,
) *RegisterService {
	return &RegisterService{
		log: log,
		s:   storage,
		c:   c,
	}
}

func Users(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("./server/web/ru/profile.html"))

	fmt.Println("Rendering news template")
	err := temp.ExecuteTemplate(w, "body", nil)
	if err != nil {
		_, _ = fmt.Fprintf(w, err.Error())
	}
}

func (reg *RegisterService) Reg(w http.ResponseWriter, r *http.Request) {

}
