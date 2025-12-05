package main

import (
	"fmt"
	"html/template"
	"net/http"
	"log"
	"time"
)

var tmpl = template.Must(template.ParseFiles("template/index.html"))

type Guest struct{
	Name string
	Time string
}

var guests []Guest

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		log.Println("Website Visited.")

		if r.Method == http.MethodPost{
			r.ParseForm()

			newGuest := r.FormValue("newguest")

			if newGuest != "" {
				guests = append(guests, Guest{
					Name : newGuest,
					Time: time.Now().Format("2006-01-02 15:04:05"),
				})
			}
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return

		}

		tmpl.Execute(w, guests)
	})

	fmt.Println("Your server is running in http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}