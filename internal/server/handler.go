package server

import (
	"fmt"
	"io"
	"net/http"
	//"crypto/sha256"
	//"crypto/subtle"
	"encoding/json"
	"log"
)

type authConfig struct {
		Username string
		Password string
		Student bool
		LeftHanded bool
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
		fmt.Fprintln(w, "here's your token")
		jd := json.NewDecoder(r.Body)
		jd.DisallowUnknownFields()
		for {
			var ac authConfig
			if err := jd.Decode(&ac); err == io.EOF {
				break		
			} else if err != nil {
				fmt.Fprintln(w, "Invalid payload")
				break
			}
			fmt.Fprintf(
				w, "username is: %s, password is: %s, student: %t, leftHanded: %t\n",
				ac.Username, ac.Password, ac.Student, ac.LeftHanded,
			)
		}
		return
	}
	fmt.Fprintln(w, "Invalid Request.")
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
