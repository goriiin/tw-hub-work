package search

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"twit-hub111/internal/db/postgres"
	"twit-hub111/internal/lib/cookies"
)

type SearchService struct {
	log *slog.Logger
	s   *postgres.Storage
	c   *cookies.CacheService
}

func New(
	log *slog.Logger,
	storage *postgres.Storage,
	c *cookies.CacheService,
) *SearchService {
	return &SearchService{
		log: log,
		s:   storage,
		c:   c,
	}
}

func (s *SearchService) Search(w http.ResponseWriter, r *http.Request) {
	var temp *template.Template
	if r.URL.Path[0:3] == "/ru" {
		temp = template.Must(template.ParseFiles("web/ru/search/search.gohtml"))
	}

	if r.URL.Path[0:3] == "/en" {
		temp = template.Must(template.ParseFiles("web/en/search/search.gohtml"))
	}

	//tok := cookie.Value

	//userId, err := n.c.GetUserIdFromToken(tok)
	//if err != nil {
	//
	//}

	//info, err := twits.ShowFeed(context.Background(), n.s, userId)
	//if err != nil {
	//    _, _ = fmt.Fprintf(w, err.Error())
	//}
	//err := temp.ExecuteTemplate(w, "body", nil)

	err := temp.ExecuteTemplate(w, "body", nil)
	if err != nil {
		_, _ = fmt.Fprintf(w, err.Error())
	}
}

type Author struct {
	Id   int    `json:"id"`
	Nick string `json:"nick,omitempty"`
}

func (s *SearchService) SearchNick(w http.ResponseWriter, r *http.Request) {
	// слой взаимодействия с базой данных

	a := []Author{Author{
		Id:   1,
		Nick: "JohnDoe",
	},
		Author{
			Id:   2,
			Nick: "John",
		},
	}
	fmt.Println("ОТПРАВИЛ")

	_ = json.NewEncoder(w).Encode(a)
}
