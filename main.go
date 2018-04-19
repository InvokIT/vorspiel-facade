package main

import (
	//"encoding/json"
	"fmt"
	"net/http"
	"os"

	//jwt "github.com/dgrijalva/jwt-go"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/steam"
	"github.com/joho/godotenv"
)

const defaultAddress = ":8080"
const defaultPublicUrl = "http://localhost:8080"

func main() {
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		logger.Fatal(err)
	}

	initAuth()

	router := CreateRouter()
	fmt.Printf("Listening on address %s\n", defaultAddress)
	logger.Fatal(http.ListenAndServe(defaultAddress, router))
}

func initAuth() {
	publicUrl := defaultPublicUrl
	if v, ok := os.LookupEnv("PUBLIC_URL"); ok {
		publicUrl = v
	}

	steamKey := os.Getenv("STEAM_KEY")
	if steamKey == "" {
		logger.Panic("STEAM_KEY is not defined.")
	}

	goth.UseProviders(steam.New(os.Getenv("STEAM_KEY"), fmt.Sprintf("%s/auth/steam/callback", publicUrl)))
}
