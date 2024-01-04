package reg

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"twit-hub111/internal/db/postgres"
	"twit-hub111/internal/domain"
	"twit-hub111/internal/lib/cookies"
	"twit-hub111/internal/lib/jwt"
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

func (reg *RegisterService) Reg(w http.ResponseWriter, r *http.Request) {
	var temp *template.Template
	if r.URL.Path[0:3] == "/ru" {
		temp = template.Must(template.ParseFiles("web/ru/sign_up/singup.gohtml"))
	}

	if r.URL.Path[0:3] == "/en" {
		temp = template.Must(template.ParseFiles("web/en/sign_up/singup.gohtml"))
	}

	err := temp.ExecuteTemplate(w, "body", nil)
	if err != nil {
		_, _ = fmt.Fprintf(w, err.Error())
	}
}

func (reg *RegisterService) RegData(w http.ResponseWriter, r *http.Request) {
	var rrr domain.RegData

	err := json.NewDecoder(r.Body).Decode(&rrr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	fmt.Println(rrr)
	id, err := reg.s.InsertUser(&rrr)
	if err != nil || id < 1 {
		reg.log.Error("Insert err", err)
		http.Redirect(w, r, r.URL.Path[0:4]+"/reg", http.StatusUnauthorized)
		return
	}

	reg.log.Info("New User", id)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	user := domain.TokenUser{
		Id:    id,
		Email: rrr.Email,
	}

	token, _ := jwt.NewToken(user)

	reg.c.SetTokenCookie(w, token)

	err = json.NewEncoder(w).Encode(map[string]string{"token": "123"})
}
