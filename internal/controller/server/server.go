package server

import (
	"context"
	"log"
	"main/internal/entity"
	"net/http"
	"time"
)

type Server struct {
	mux    *http.ServeMux
	server *http.Server
}

func NewServer(listenAddr string) *Server {
	m := http.NewServeMux()

	return &Server{
		mux:    m,
		server: &http.Server{Addr: listenAddr, Handler: m},
	}
}

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
		result, err := router.UCWallet.SendCurrency(ctx, e, amount, currency, receiver, mine)
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
	history, err := router.UCWallet.GetTransactionsHistory(ctx, e)
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
		balance, err := router.UCWallet.GetBalance(ctx, e)
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

		if is_user, err := router.UCLogin.Login(ctx, e); !is_user {
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

		if is_user, err := router.UCRegistration.Register(ctx, e); !is_user {
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

func (s *Server) Start(router *Router, logger *log.Logger) {
	s.mux.HandleFunc("/", router.indexHandler)

	s.mux.HandleFunc("/registration", router.registrationHandler)
	s.mux.HandleFunc("/login", router.loginHandler)

	s.mux.HandleFunc("/main", router.mainPageHandler)

	s.mux.HandleFunc("/get_transactions_history", router.transactionsHistoryHandler)
	s.mux.HandleFunc("/send_currency", router.sendCurrencyHandler)

	logger.Printf("Server is listening on http://%s\n", s.server.Addr)

	if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatal("ListenAndServe:", err)
	}
}

func (s *Server) Shutdown(ctx context.Context) error {
	if err := s.server.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}
