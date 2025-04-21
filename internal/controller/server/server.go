package server

import (
	"context"
	"log"
	"main/internal/entity"
	"net/http"
	"time"
)

func (router *Router) loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
		defer cancel()

		e := entity.Login{Username: username, Password: password}

		if is_user, err := router.LoginUC.Login(ctx, e); !is_user {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		} else if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write([]byte("Welcome " + username + "!"))
		return
	}

	if err := tmpl.ExecuteTemplate(w, "login", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (router *Router) registrationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")
		if len(password) < 8 {
			w.Write([]byte("Password is less than 8 symbols!"))
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
		defer cancel()

		e := entity.Registration{Username: username, Password: password}

		if is_user, err := router.RegistrationUC.Register(ctx, e); !is_user {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} else if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write([]byte("User " + username + " registered successfully!"))
		return
	}

	if err := tmpl.ExecuteTemplate(w, "registration", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (router *Router) indexHandler(w http.ResponseWriter, r *http.Request) {
	if err := tmpl.ExecuteTemplate(w, "index", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func StartServer(appPort string, router *Router) {

	http.HandleFunc("/", router.indexHandler)
	http.HandleFunc("/registration", router.registrationHandler)
	http.HandleFunc("/login", router.loginHandler)

	log.Printf("Server is listening on http://localhost%s\n", appPort)
	if err := http.ListenAndServe(appPort, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
