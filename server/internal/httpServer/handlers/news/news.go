package news

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"twit-hub111/internal/db/postgres"
	"twit-hub111/internal/lib/cookies"
)

// TODO: отслеживание лайков

type NewsService struct {
	log *slog.Logger
	s   *postgres.Storage
	c   *cookies.CacheService
}

func New(
	log *slog.Logger,
	storage *postgres.Storage,
	c *cookies.CacheService,
) *NewsService {
	return &NewsService{
		log: log,
		s:   storage,
		c:   c,
	}
}

func (n *NewsService) News(w http.ResponseWriter, r *http.Request) {
	var temp *template.Template
	//cookie, err := r.Cookie("token")
	//
	//token, err := jwt.NewToken(domain.TokenUser{1, "123@mail.ru"})
	//c := http.Cookie{
	//	Name:     "token",
	//	Value:    token,
	//	Path:     "/",
	//	Expires:  time.Now().Add(time.Hour),
	//	MaxAge:   3600,
	//	HttpOnly: true,
	//	Secure:   false,
	//	SameSite: http.SameSiteLaxMode,
	//}
	//
	//// Устанавливаем cookie в браузере
	//http.SetCookie(w, &c)
	//flag, err := n.c.IsCookieValid(w, r)
	//if err != nil {
	//	http.Redirect(w, r, r.URL.Path[0:4]+"/login", http.StatusInternalServerError)
	//}
	//
	//if !flag {
	//	http.Redirect(w, r, r.URL.Path[0:4]+"/login", http.StatusUnauthorized)
	//}

	cookie, err := r.Cookie("token")

	id, _ := n.c.GetUserIdFromToken(cookie.Value)
	fmt.Println(id)

	if r.URL.Path[0:3] == "/ru" {
		temp = template.Must(template.ParseFiles("web/ru/news/newsFeed.gohtml"))
	}

	if r.URL.Path[0:3] == "/en" {
		temp = template.Must(template.ParseFiles("web/en/news/newsFeed.gohtml"))
	}

	err = temp.ExecuteTemplate(w, "body", nil)

	//err := temp.ExecuteTemplate(w, "body", info)
	if err != nil {
		_, _ = fmt.Fprintf(w, err.Error())
	}
}

func (n *NewsService) NewPost(w http.ResponseWriter, r *http.Request) {
	var data map[string]string

	//cookie, err := r.Cookie("token")
	//tok := cookie.Value
	//fmt.Println(tok)
	//userId, err := n.c.GetUserIdFromToken(tok)
	//if err != nil {
	//
	//}
	//fmt.Println(userId)

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO получаю id с куки и отдаю в базу
	fmt.Println("Received text:", data)
}

func (n *NewsService) RenderNews(w http.ResponseWriter, r *http.Request) {

	//cookie, err := r.Cookie("token")
	//tok := cookie.Value
	//flag, _ := n.c.IsCookieValid(w, r)
	//if !flag {
	//	http.Redirect(w, r, r.URL.Path[0:4]+"/login", http.StatusUnauthorized)
	//}
	//id, err := n.c.GetUserIdFromToken(tok)
	//if err != nil {
	//	w.WriteHeader(http.StatusInternalServerError)
	//}

	cookie, _ := r.Cookie("token")

	id, _ := n.c.GetUserIdFromToken(cookie.Value)
	fmt.Println(id)

	ppp, _ := n.s.PostsFromSubs(id)

	_ = json.NewEncoder(w).Encode(ppp)
}
