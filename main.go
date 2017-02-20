package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"net/http"
	"time"

	"database/sql"
	_ "github.com/lib/pq"
)

var secret = []byte("secrety") //get this from os.env

var DB *sql.DB

func main() {

	var err error

	DB, err = sql.Open("postgres", "password=password  user=user dbname=my_db sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}

	_, err = DB.Exec("CREATE TABLE IF NOT EXISTS login(name varchar, password varchar, salt varchar)")

	if err != nil {
		fmt.Println(err)
	}

	// _, err = db.Exec("insert into login values('mohan', 'meme', 'meme')")

	if err != nil {
		fmt.Println(err)
	}

	r := mux.NewRouter()

	r.HandleFunc("/auth", GetTokenHandler).Methods("POST")

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	http.ListenAndServe(":3000", r)
}

func GetTokenHandler(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()

	if err != nil {
		panic(err)
	}

	username := r.Form["name"][0]
	password := r.Form["password"][0]
	redirectUrl := r.Form["redirect"]

	tokenString, err := authorize(username, password)

	if err != nil {
		fmt.Fprintln(w, err)
	}

	cookie := &http.Cookie{Name: "jwtcookie", Value: tokenString, MaxAge: 3600, Secure: false, HttpOnly: true, Raw: tokenString}
	http.SetCookie(w, cookie)

	if redirectUrl != nil {
		http.Redirect(w, r, redirectUrl[0], 302)
	} else {
		w.Write([]byte(tokenString))
	}

}

func authorize(username string, password string) (token string, autherr error) {

	var dbpassword string

	rows, err := DB.Query("Select password from login where name=$1", username)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		err = rows.Scan(&dbpassword)

		if err != nil {
			panic(err)
		}

		break

	}

	if password == dbpassword {

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": "mohan",
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

		tokenString, _ := token.SignedString(secret)

		return tokenString, nil

	} else {
		autherr = fmt.Errorf("Cannot authorize %q", username)

		return token, autherr
	}

}
