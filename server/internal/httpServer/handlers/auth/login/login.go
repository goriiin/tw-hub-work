package login

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

	data, err := l.s.UserHashPass(lll.Email)
	fmt.Println(data, lll)
	if err != nil {
		l.log.Error("DB ERROR", err)
	}

	if lll.Password == data.Pass {
		user := domain.TokenUser{
			Id:    data.Id,
			Email: data.Email,
		}

		token, _ := jwt.NewToken(user)

		cookie := http.Cookie{
			Name:     "token",
			Value:    token,
			MaxAge:   3500,
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
		}
		http.SetCookie(w, &cookie)

		l.c.SetTokenCookie(w, token)

		t, err := r.Cookie("token")
		if err != nil {
			fmt.Println("ошибка в получении", err)
		} else {
			if t.Value != "" {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				err = json.NewEncoder(w).Encode(map[string]string{"token": "123"})
			}

		}
		fmt.Println(t)
	} else {
		http.Redirect(w, r, r.URL.Path[0:4]+"/login", http.StatusNotFound)
	}
}
