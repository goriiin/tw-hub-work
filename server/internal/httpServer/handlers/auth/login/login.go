package login

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"time"
	"twit-hub111/internal/db/postgres"
	"twit-hub111/internal/domain"
	"twit-hub111/internal/lib/cookies"
	"twit-hub111/internal/lib/jwt"
)

type LoginService struct {
	log *slog.Logger
	s   *postgres.Storage
	c   *cookies.CacheService
}

func New(
	log *slog.Logger,
	storage *postgres.Storage,
	c *cookies.CacheService,
) *LoginService {
	return &LoginService{
		log: log,
		s:   storage,
		c:   c,
	}
}

func (l *LoginService) Login(w http.ResponseWriter, r *http.Request) {
	var temp *template.Template
	if r.URL.Path[0:3] == "/ru" {
		temp = template.Must(template.ParseFiles("web/ru/sign_in/sign_in.gohtml"))
	}

	if r.URL.Path[0:3] == "/en" {
		temp = template.Must(template.ParseFiles("web/en/sign_in/sign_in.gohtml"))
	}

	err := temp.ExecuteTemplate(w, "body", nil)
	if err != nil {
		_, _ = fmt.Fprintf(w, err.Error())
	}

}

type loginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (l *LoginService) LogData(w http.ResponseWriter, r *http.Request) {
	var lll loginData
	err := json.NewDecoder(r.Body).Decode(&lll)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	fmt.Println(lll)

	data, err := l.s.UserHashPass(lll.Email)
	if err != nil {
		l.log.Error("DB ERROR", err)
	}

	if lll.Password == data.Pass {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		user := domain.TokenUser{
			Id:    data.Id,
			Email: data.Email,
		}

		token, _ := jwt.NewToken(user)

		l.c.SetTokenCookie(w, token, time.Hour*10)

		err = json.NewEncoder(w).Encode(map[string]string{"token": "123"})
	} else {
		http.Redirect(w, r, r.URL.Path[0:4]+"/login", http.StatusNotFound)
	}
}
