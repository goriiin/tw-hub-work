package login

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log/slog"
	"net/http"
	"twit-hub111/internal/db/postgres"
	"twit-hub111/internal/lib/cookies"
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
		temp = template.Must(template.ParseFiles("server/web/ru/sign_in/sign_in.html"))
	}

	if r.URL.Path[0:3] == "/en" {
		temp = template.Must(template.ParseFiles("server/web/en/sign_in/sign_in.html"))
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
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {

	}
	var lll loginData
	json.Unmarshal(body, &lll)

	fmt.Println(lll)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(map[string]string{"token": "123"})
}
