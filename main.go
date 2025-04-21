package main

import (
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// entity
type RegistrationData struct {
	username string
	password string
}

// business cases
func indexHandler(w http.ResponseWriter, r *http.Request) {
	if err := tmpl.ExecuteTemplate(w, "index", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")
		if len(password) < 8 {
			w.Write([]byte("Password is less than 8 symbols!"))
			return
		}

		passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		passwordHashStr := string(passwordHash)
		// dbWrite(username, passwordHashStr)

		w.Write([]byte("User " + username + " registered successfully!"))
		return
	}

	if err := tmpl.ExecuteTemplate(w, "signup", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func signinHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		w.Write([]byte("Welcome back, " + username + "!"))
		return
	}

	if err := tmpl.ExecuteTemplate(w, "signin", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	port := ":3004"

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/signup", signupHandler)
	http.HandleFunc("/signin", signinHandler)

	log.Printf("Server is listening on http://localhost%s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
