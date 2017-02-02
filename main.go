package main 

import (
    "net/http"
    "time"
    "github.com/gorilla/mux"
    "github.com/dgrijalva/jwt-go"
)

var secret = []byte("secrety") //get this from os.env

func main(){
    r := mux.NewRouter()

    r.HandleFunc("/", GetTokenHandler).Methods("GET")

    http.ListenAndServe(":3000", r)
}

func GetTokenHandler(w http.ResponseWriter, r *http.Request){

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
            "username" : "mohan",
            "exp" : time.Now().Add(time.Hour * 24).Unix(),
        })

    tokenString, _ := token.SignedString(secret)
    
    w.Write([]byte(tokenString))
}