package server

import (
	"context"
	"log"
	"main/internal/entity"
	"net/http"
	"time"
)

func (router *Router) mainPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Invalid form", http.StatusBadRequest)
			return
		}
		address := r.FormValue("walletAddress")
		//getBalance. execute mainPage template if incorrect address and return
		balance := []string{"Balance"} //FIXME

		data := struct {
			WAddress string
			WBalance []string
		}{address, balance}

		tmpl.ExecuteTemplate(w, "walletActions", data)

		return
	}

	if err := tmpl.ExecuteTemplate(w, "mainPage", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

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

		http.Redirect(w, r, "/main", http.StatusSeeOther)
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

		data := struct{ Username string }{Username: username}
		if err := tmpl.ExecuteTemplate(w, "registrationSuccess", data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

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
	http.HandleFunc("/main", router.mainPageHandler)

	log.Printf("Server is listening on http://localhost%s\n", appPort)
	if err := http.ListenAndServe(appPort, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
