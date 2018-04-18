package main

import (
	//"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"html/template"

	//jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/pat"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/steam"
	"github.com/joho/godotenv"
)

var authSuccessTemplate *template.Template = nil
const DefaultAddress = ":8080"
const DefaultPublicUrl = "http://localhost:8080"

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	if tmpl, err := template.ParseFiles("auth-success.html"); err == nil {
		authSuccessTemplate = tmpl
	} else {
		panic(err)
	}

	publicUrl, ok := os.LookupEnv("PUBLIC_URL")
	if !ok {
		publicUrl = DefaultPublicUrl
	}
	log.Printf("publicUrl = %s", publicUrl)

	address, ok := os.LookupEnv("ADDRESS")
	if !ok {
		address = DefaultAddress
	}
	log.Printf("address = %s", address)

	goth.UseProviders(steam.New(os.Getenv("STEAM_KEY"), fmt.Sprintf("%s/auth/steam/callback", publicUrl)))

	router := pat.New()
	router.Get("/auth/{provider}/callback", AuthCallbackHandler)
	router.Get("/logout/{provider}", LogoutHandler)
	router.Get("/auth/{provider}", AuthHandler)

	log.Printf("Listening on address %s", address)
	log.Fatal(http.ListenAndServe(address, router))
}

func AuthCallbackHandler(res http.ResponseWriter, req *http.Request) {
	log.Print("AuthCallbackHandler entered")

	gothUser, err := gothic.CompleteUserAuth(res, req)
	if err != nil {
		log.Print(err)
		fmt.Fprint(res, err)
		res.WriteHeader(http.StatusUnauthorized)
		return
	}

	log.Printf("AuthCallbackHandler got user: %s", gothUser)

	if err := authSuccessTemplate.Execute(res, gothUser); err != nil {
		log.Print(err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	//res.WriteHeader(http.StatusOK)
}

func LogoutHandler(res http.ResponseWriter, req *http.Request) {
	log.Print("LogoutHandler entered")

	gothic.Logout(res, req)
	//res.Header().Set("Location", "/")
	res.WriteHeader(http.StatusOK)
}

func AuthHandler(res http.ResponseWriter, req *http.Request) {
	log.Print("AuthHandler entered")

	// try to get the user without re-authenticating
	if gothUser, err := gothic.CompleteUserAuth(res, req); err == nil {
		if err := authSuccessTemplate.Execute(res, gothUser); err != nil {
			log.Fatal(err)
			res.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	gothic.BeginAuthHandler(res, req)
}
