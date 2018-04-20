package main

import (
	//"encoding/json"
	//"fmt"
	"os"

	//jwt "github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"github.com/invokit/vorspiel-lib/mq"
	googlemq "github.com/invokit/vorspiel-lib/google/mq"
)

const defaultPublicUrl = "http://localhost:8080"

func main() {
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		logger.Fatal(err)
	}

	publicUrl := defaultPublicUrl
	if v, ok := os.LookupEnv("PUBLIC_URL"); ok {
		publicUrl = v
	}

	router := BuildRouter()
	mq = BuildMq()

	app := &App{Port: 8080, PublicUrl: publicUrl, MQ: mq, Router: router}

	app.Start()
}

func buildMq() *mq.Client {
	mq := googlemq.New()
	return mq
}
