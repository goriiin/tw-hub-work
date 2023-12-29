package profile

import (
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"twit-hub111/internal/db/postgres"
	"twit-hub111/internal/lib/cookies"
)

type UserService struct {
	log *slog.Logger
	s   *postgres.Storage
	c   *cookies.CacheService
}

func New(
	log *slog.Logger,
	storage *postgres.Storage,
	c *cookies.CacheService,
) *UserService {
	return &UserService{
		log: log,
		s:   storage,
		c:   c,
	}
}

func (u *UserService) User(w http.ResponseWriter, r *http.Request) {
	var temp *template.Template

	//cookie, err := r.Cookie("token")
	//flag, err := u.c.IsCookieValid(cookie)
	//if err != nil {
	//	http.Redirect(w, r, r.URL.Path[0:4]+"/wtf", http.StatusInternalServerError)
	//}
	//
	//if !flag {
	//	http.Redirect(w, r, r.URL.Path[0:4]+"/login", http.StatusUnauthorized)
	//}

	if r.URL.Path[0:3] == "/ru" {
		temp = template.Must(template.ParseFiles("web/ru/profile/profile.gohtml"))
	}

	if r.URL.Path[0:3] == "/en" {
		temp = template.Must(template.ParseFiles("web/en/profile/profile.gohtml"))
	}

	//tok := cookie.Value

	//userId, err := u.c.GetUserIdFromToken(tok)
	//if err != nil {
	//
	//}
	//fmt.Println(userId)
	//fmt.Println("Rendering news template")
	//
	//ui, _ := profile.UserInfo(context.Background(), u.s, userId)
	//up, _ := profile.UserPosts(context.Background(), u.s, userId)

	//info, err := profile.InfoToJSON(ui, up)
	//if err != nil {
	//	_, _ = fmt.Fprintf(w, err.Error())
	//}
	err := temp.ExecuteTemplate(w, "body", nil)

	//err = temp.ExecuteTemplate(w, "body", info)
	if err != nil {
		_, _ = fmt.Fprintf(w, err.Error())
	}
}

func (u *UserService) Profile(w http.ResponseWriter, r *http.Request) {

}

func (u *UserService) Follow(w http.ResponseWriter, r *http.Request) {
	// TODO: берем id и смотрим, есть ли подписка, если нет, то возвращаем ok, при наличии подписки, также ok
}

func (u *UserService) IsFollow(w http.ResponseWriter, r *http.Request) {
	// TODO: проверка есть ли подписка
}

func (u *UserService) UnFollow(w http.ResponseWriter, r *http.Request) {

}
