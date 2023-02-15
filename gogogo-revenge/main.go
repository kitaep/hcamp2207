package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

type User struct {
	ID       int
	Email    string
	Password string
	GetFlag  func(a int) string
}

func main() {
	user1 := User{1, "hcamp@hcamp.com", "hackingcamp123", func(a int) string {
		data, err := os.ReadFile("flag")

		if err != nil {
			log.Panic(err)
		}

		return string(data)
	}}

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
