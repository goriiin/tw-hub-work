package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//начало

func main() {
	http.HandleFunc("/sign-up", handleSignIn)
	http.HandleFunc("/reg", handle)
	http.ListenAndServe(":8080", nil)
}

func handle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "ошибка парсинга", http.StatusBadRequest)
			return
		}

		email := r.Form.Get("email")
		username := r.Form.Get("username")
		password := r.Form.Get("password")

		fmt.Println(email, username, password)
		fmt.Fprintf(w, "получили:\nEmail: %s\nUsername: %s\nPassword: %s", email, username, password)
	} else {
		http.Error(w, "НЕТ ОБРАБОТЧИКА ПОД МЕТОД", http.StatusMethodNotAllowed)
	}

}

func handleSignIn(w http.ResponseWriter, r *http.Request) {
	html, err := ioutil.ReadFile("/home/dmitry/Desktop/twit-hub/web/html/sign_in/sign_in.html")
	if err != nil {
		http.Error(w, "Error reading HTML file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write(html)
}
