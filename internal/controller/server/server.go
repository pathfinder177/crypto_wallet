package server

import (
	"context"
	"log"
	"main/internal/entity"
	"net/http"
	"time"
)

func (router *Router) sendCurrencyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Invalid form", http.StatusBadRequest)
			return
		}
		amount := r.FormValue("amount")
		currency := r.FormValue("currency")
		sender := r.FormValue("address")
		receiver := r.FormValue("receiver")
		mine := r.FormValue("mine")

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		e := entity.Wallet{Address: sender}
		result, err := router.WalletUC.SendCurrency(ctx, e, amount, currency, receiver, mine)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		data := struct {
			WAmount   string
			WCurrency string
			WSender   string
			WReceiver string
			WResult   string
		}{amount, currency, sender, receiver, result}

		tmpl.ExecuteTemplate(w, "successSendCurrency", data)
	}
}

func (router *Router) transactionsHistoryHandler(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")
	if address == "" {
		http.Error(w, "missing address", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	e := entity.Wallet{Address: address}
	history, err := router.WalletUC.GetTransactionsHistory(ctx, e)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data := struct {
		WAddress  string
		WCurrency string
		WHistory  []string
	}{address, "", history}

	tmpl.ExecuteTemplate(w, "getTransactionsHistory", data)
}

func (router *Router) mainPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Invalid form", http.StatusBadRequest)
			return
		}
		address := r.FormValue("walletAddress")
		if address == "" {
			http.Error(w, "missing address", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		e := entity.Wallet{Address: address}
		balance, err := router.WalletUC.GetBalance(ctx, e)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

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

		if username == "" || password == "" {
			http.Error(w, "missing login or password", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
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

		if username == "" || password == "" {
			http.Error(w, "missing login or password", http.StatusBadRequest)
			return
		} else if len(password) < 8 {
			http.Error(w, "password is less than 8 symbols!", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
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

	http.HandleFunc("/get_transactions_history", router.transactionsHistoryHandler)
	http.HandleFunc("/send_currency", router.sendCurrencyHandler)

	log.Printf("Server is listening on http://localhost%s\n", appPort)
	if err := http.ListenAndServe(appPort, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
