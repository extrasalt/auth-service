package main 

import (
"net/http"
"time"
"github.com/gorilla/mux"
"github.com/dgrijalva/jwt-go"
"fmt"
)

var secret = []byte("secrety") //get this from os.env

func main(){
    r := mux.NewRouter()

    r.HandleFunc("/auth", GetTokenHandler).Methods("POST")

    r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

    

    http.ListenAndServe(":3000", r)
}

func GetTokenHandler(w http.ResponseWriter, r *http.Request){


    err := r.ParseForm()

    if err != nil {
        panic(err)
    }

    username := r.Form["name"][0]
    password := r.Form["password"][0]   

    fmt.Println(username, password)

    if username == "mohan" && password == "meme" {

        token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
            "username" : "mohan",
            "exp" : time.Now().Add(time.Hour * 24).Unix(),
            })

        tokenString, _ := token.SignedString(secret)

        cookie := &http.Cookie{Name: "jwtcookie", Value: tokenString, MaxAge: 3600, Secure: false, HttpOnly: true, Raw: tokenString}
        http.SetCookie(w, cookie)

        w.Write([]byte(tokenString))
    } else {
        w.Write([]byte("Auth failed"))
    }
}