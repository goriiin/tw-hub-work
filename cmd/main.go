package main

import (
	"fmt"
	"net/http"
	"text/template"
)

//https://qna.habr.com/q/915835

func news(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("web/news/newsFeed.html"))

	fmt.Println("Rendering news template")
	err := temp.ExecuteTemplate(w, "body", nil)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	//w.Header().Set("Content-Type", "text/html")
	//http.ServeFile(w, r, "web/news/newsFeed.html")
}

func main() {
	//db.Test()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/news"))))
	http.HandleFunc("/news", news)
	http.ListenAndServe(":8080", nil)
}
