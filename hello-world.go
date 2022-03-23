package main

import (
    "encoding/json"
    "fmt"
		"html/template"
    "net/http"
)

type User struct {
    Firstname string `json:"firstname"`
    Lastname  string `json:"lastname"`
    Age       int    `json:"age"`
}

type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

func main() {
		tmpl := template.Must(template.ParseFiles("forms.html"))

		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				if r.Method != http.MethodPost {
						tmpl.Execute(w, nil)
						return
				}

				
				details := ContactDetails{
						Email:   r.FormValue("email"),
						Subject: r.FormValue("subject"),
						Message: r.FormValue("message"),

				}

				fmt.Println(details)
        json.NewEncoder(w).Encode(details)

				// do something with details
				_ = details

				tmpl.Execute(w, struct{ Success bool }{true})
		})

    http.HandleFunc("/decode", func(w http.ResponseWriter, r *http.Request) {
        var user User
        json.NewDecoder(r.Body).Decode(&user)

        fmt.Fprintf(w, "%s %s is %d years old!", user.Firstname, user.Lastname, user.Age)
    })

    http.HandleFunc("/encode", func(w http.ResponseWriter, r *http.Request) {
        peter := User{
            Firstname: "John",
            Lastname:  "Doe",
            Age:       25,
        }

        json.NewEncoder(w).Encode(peter)
    })

    http.ListenAndServe(":8080", nil)
}