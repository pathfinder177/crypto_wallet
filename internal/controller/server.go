package controller

import (
	"context"
	"log"
	"main/internal/entity"
	"main/internal/usecase/registration"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Server struct {
	RUC *registration.UseCase
}

func (s *Server) registrationHandler(w http.ResponseWriter, r *http.Request) {
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

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
		defer cancel()

		e := entity.Registration{Username: username, HashedPassword: string(passwordHash)}
		s.RUC.Register(ctx, e) //FIXME

		w.Write([]byte("User " + username + " registered successfully!"))
		return
	}

	if err := tmpl.ExecuteTemplate(w, "registration", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if err := tmpl.ExecuteTemplate(w, "index", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func StartServer(appPort string, reg *registration.UseCase) {

	server := Server{RUC: reg}

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/registration", server.registrationHandler)
	// http.HandleFunc("/login", loginHandler)

	log.Printf("Server is listening on http://localhost%s\n", appPort)
	if err := http.ListenAndServe(appPort, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
