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
	cookie, err := r.Cookie("token")
	flag, err := n.c.IsCookieValid(cookie)
	if err != nil {
		http.Redirect(w, r, r.URL.Path[0:4]+"/news", http.StatusInternalServerError)
	}

	if !flag {
		http.Redirect(w, r, r.URL.Path[0:4]+"/login", http.StatusUnauthorized)
	}

	if r.URL.Path[0:3] == "/ru" {
		temp = template.Must(template.ParseFiles("server/web/ru/news/newsFeed.html"))
	} else {
		temp = template.Must(template.ParseFiles("server/web/en/news/newsFeed.html"))
	}

	tok := cookie.Value

	userId, err := n.c.GetUserIdFromToken(tok)
	if err != nil {

	}
	fmt.Println(userId)
	fmt.Println("Rendering news template")

	//info, err := twits.ShowFeed(context.Background(), n.s, userId)
	if err != nil {
		_, _ = fmt.Fprintf(w, err.Error())
	}
	err = temp.ExecuteTemplate(w, "body", nil)
	//err = temp.ExecuteTemplate(w, "body", info)
	if err != nil {
		_, _ = fmt.Fprintf(w, err.Error())
	}
}

func (n *NewsService) NewPost(w http.ResponseWriter, r *http.Request) {
	var data map[string]string

	cookie, err := r.Cookie("token")
	tok := cookie.Value

	userId, err := n.c.GetUserIdFromToken(tok)
	fmt.Println(userId)
	if err != nil {

	}
	fmt.Println(userId)

	if err = json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO получаю id с куки и отдаю в базу
	fmt.Println("Received text:", data)
}
