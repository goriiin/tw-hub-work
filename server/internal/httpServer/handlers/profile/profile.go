package profile

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"html/template"
	"log/slog"
	"net/http"
	"strconv"
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

type UserProfile struct {
	Username string
}

type UUU struct {
	Id       string
	Username string
}

func (u *UserService) User(w http.ResponseWriter, r *http.Request) {
	var temp *template.Template

	if r.URL.Path[0:3] == "/ru" {
		temp = template.Must(template.ParseFiles("web/ru/user/user.gohtml"))
	}

	if r.URL.Path[0:3] == "/en" {
		temp = template.Must(template.ParseFiles("web/en/user/user.gohtml"))
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

	idS := chi.URLParam(r, "id")

	id, _ := strconv.Atoi(idS)

	uuu, _ := u.s.SearchUserID(id)

	user := UUU{Id: idS, Username: uuu.Nick}

	err := temp.ExecuteTemplate(w, "body", user)

	//err = temp.ExecuteTemplate(w, "body", info)
	if err != nil {
		_, _ = fmt.Fprintf(w, err.Error())
	}
}

func (u *UserService) RenderNewsProf(w http.ResponseWriter, r *http.Request) {
	//idS := chi.URLParam(r, "id")
	//
	//id, _ := strconv.Atoi(idS)
	//

	fmt.Println("RENDER_NEWS_PROFILE")
	cookie, _ := r.Cookie("token")

	id, _ := u.c.GetUserIdFromToken(cookie.Value)

	ppp, _ := u.s.MyPost(id)

	_ = json.NewEncoder(w).Encode(ppp)
}

func (u *UserService) RenderNews(w http.ResponseWriter, r *http.Request) {
	idS := chi.URLParam(r, "id")

	fmt.Println("RENDER_NEWS", idS)

	id, _ := strconv.Atoi(idS)

	ppp, _ := u.s.MyPost(id)

	_ = json.NewEncoder(w).Encode(ppp)
}

func (u *UserService) Profile(w http.ResponseWriter, r *http.Request) {
	var temp *template.Template
	if r.URL.Path[0:3] == "/ru" {
		temp = template.Must(template.ParseFiles("web/ru/profile/profile.gohtml"))
	}

	if r.URL.Path[0:3] == "/en" {
		temp = template.Must(template.ParseFiles("web/en/profile/profile.gohtml"))
	}

	cookie, _ := r.Cookie("token")
	fmt.Println(cookie)

	id, _ := u.c.GetUserIdFromToken(cookie.Value)
	fmt.Println(id)

	aaa, _ := u.s.SearchUserID(id)

	fmt.Println(aaa)

	user := UserProfile{Username: aaa.Nick}

	_ = temp.ExecuteTemplate(w, "body", user)
}

func (u *UserService) Follow(w http.ResponseWriter, r *http.Request) {
	// TODO: берем id и смотрим, есть ли подписка, если нет, то возвращаем ok, при наличии подписки, также ok
}

func (u *UserService) IsFollow(w http.ResponseWriter, r *http.Request) {
	// TODO: проверка есть ли подписка
}

func (u *UserService) UnFollow(w http.ResponseWriter, r *http.Request) {

}
