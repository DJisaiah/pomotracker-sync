package server

import (
	"fmt"
	"io"
	"net/http"
	"encoding/json"
	"log"
)

type authConfig struct {
	Username string
	Password string
	Student bool
	LeftHanded bool
}

type tokenResponse struct {
	token string
	validTil string
}

type application struct {
	auth authConfig
}



func (app *application) protectedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "this is the protected handler")
}

func (app *application) unprotectedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "this is the unprotected handler")
}

func (app *application) register(w http.ResponseWriter, r *http.Request) {
	// only accept POST and json on this endpoint
	validRequest := r.Header.Get("Content-Type") == "application/json"
	if validRequest {
		// serveHTTP does this, we dont actually need to
		defer r.Body.Close()
		jd := json.NewDecoder(r.Body)
		jd.DisallowUnknownFields()
		var ac authConfig
		if err := jd.Decode(&ac); err != nil {
			http.Error(w, "Invalid Payload", http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		t, v := register_user(ac)
		tr := tokenResponse{
			token: t,
			validTil: v,
		}
		if err := json.NewEncoder(w).Encode(tr); err != nil {
			http.Error(w, "something went wrong; cannot serialise token", http.StatusInternalServerError)
			return
		}
	}
}

func (app *application) basicAuth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if ok {
			fmt.Printf("Username is: %s, Password is: %s\n", username, password)
			fmt.Println("serving...")
			next.ServeHTTP(w, r)
			return
		}
		fmt.Println("not okay")
		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})
}


func Start() {
	app := application{}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /lobby", app.unprotectedHandler)
	mux.HandleFunc("GET /login", app.basicAuth(app.protectedHandler))
	mux.HandleFunc("POST /register", app.register)

	srv := &http.Server{
		Addr: ":8080",
		Handler: mux,
	}
	err := srv.ListenAndServe()
	log.Fatal(err)
}
