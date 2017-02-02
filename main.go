package main 

import (
    "net/http"
    "github.com/gorilla/mux"
)

func main(){
    r := mux.NewRouter()

    r.HandleFunc("/", GetTokenHandler).Methods("GET")

    http.ListenAndServe(":3000", r)
}

func GetTokenHandler(w http.ResponseWriter, r *http.Request){

    

    w.Write([]byte("Hello world"))
}