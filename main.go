package main

import (
  "fmt"
  "log"
  "net/http"
  "os"
  "github.com/gorilla/pat"
  "github.com/markbates/goth"
  "github.com/markbates/goth/providers/steam"
  jwt "github.com/dgrijalva/jwt-go"
)

func main() {
  goth.UseProviders(
    steam.New(
      os.Getenv("STEAM_KEY"),
      fmt.Sprintf("%s/auth/steam/callback", os.Getenv("PUBLIC_URL"))
    )
  )

  router := pat.New()
  router.Get("/auth/{provider}/callback", AuthCallbackHandler)
  router.Get("/logout/{provider}", LogoutHandler)
  router.Get("/auth/{provider}", AuthHandler)

  log.Fatal(http.ListenAndServe(":8080", router))
}

func AuthCallbackHandler(res http.ResponseWriter, req *http.Request) {
  user, err := gothic.CompleteUserAuth(res, req)
  if err != nil {
    fmt.Fprint(res, err)
    res.WriteHeader(http.StatusUnauthorized)
		return
	}

  jwt := BuildJwt(user)

  // TODO Send JWT to client somehow
  res.Header()["JWT"] = jwt
  res.WriteHeader(http.StatusOK)
}

func LogoutHandler(res http.ResponseWriter, req *http.Request) {

}

func AuthHandler(res http.ResponseWriter, req *http.Request) {

}

func BuildJwt(goth.User) (jwt.Token) {
  
}
