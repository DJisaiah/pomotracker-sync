package server

import (
	"fmt"
	"net/http"
	//"crypto/sha256"
	//"crypto/subtle"
	"log"
)

type authConfig struct {
		username string
		password string
}

type application struct {
	auth authConfig
}

func (app *application) protectedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "this is the protected handler")
	fmt.Fprintf(w, "Welcome %s!\n", app.auth.username)
}

func (app *application) unprotectedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "this is the unprotected handler")
}


func Start() {
	app := application{
		auth: authConfig{
			username: "Isaiah",
			password: "hello",
		},
	}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /lobby", app.unprotectedHandler)
	mux.HandleFunc("GET /login", app.protectedHandler)

	srv := &http.Server{
		Addr: ":8080",
		Handler: mux,
	}
	err := srv.ListenAndServe()
	log.Fatal(err)
}
