package news

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

// TODO: отслеживание лайков

func News(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("./server/web/static/news/newsFeed.html"))

	fmt.Println("Rendering news template")
	err := temp.ExecuteTemplate(w, "body", nil)
	if err != nil {
		_, _ = fmt.Fprintf(w, err.Error())
	}
}

func NewPost(w http.ResponseWriter, r *http.Request) {
	var data map[string]string

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("Received text:", data)
}
