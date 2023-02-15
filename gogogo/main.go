package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type User struct {
	ID       int
	Email    string
	Password string
}

func (u User) GetPassword() string {
	return u.Password
}

func (u User) GetFlag() string {
	data, err := os.ReadFile("flag")

	if err != nil {
		log.Panic(err)
	}

	return string(data)
}

func main() {
	user1 := User{1, "hcamp@hcamp.com", "hackingcamp123"}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var tmpl = fmt.Sprintf(`
			<html>
			<head>
			<title>gogogo</title>
			</head>
			<h1>let's go to hackingcamp gogogogo</h1>
			<p>%s</p>
			</html>`,
			r.URL.Query()["q"])

		t := template.Must(template.New("page").Parse(tmpl))

		err := t.Execute(w, user1)

		if err != nil {
			fmt.Println(err)
		}
	})
	http.ListenAndServe(":3000", nil)
}
