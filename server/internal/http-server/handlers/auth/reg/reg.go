package reg

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
		temp = template.Must(template.ParseFiles("web/ru/sign_up/singup.html"))
	}

	if r.URL.Path[0:3] == "/en" {
		temp = template.Must(template.ParseFiles("web/en/sign_up/singup.html"))
	}

	err := temp.ExecuteTemplate(w, "body", nil)
	if err != nil {
		_, _ = fmt.Fprintf(w, err.Error())
	}
}

type regData struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (reg *RegisterService) RegData(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {

	}
	var rrr regData
	json.Unmarshal(body, &rrr)

	fmt.Println(rrr)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(map[string]string{"token": "123"})
}
