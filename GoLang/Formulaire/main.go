package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintf(w, "<html><head><title>Formulaire</title><style>label,input{display:block;margin-bottom:10px;}</style></head><body><form method='POST'><label for='nom'>Nom :</label><input type='text' id='nom' name='nom'><label for='email'>Email :</label><input type='email' id='email' name='email'><button type='submit'>Envoyer</button></form></body></html>")
	} else if r.Method == "POST" {
		nom := r.FormValue("nom")
		email := r.FormValue("email")

		fmt.Printf("Nom : %s\nEmail : %s\n", nom, email)

		fmt.Fprintf(w, "Merci pour votre soumission !")
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
