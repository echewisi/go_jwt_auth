package main

import (
	"net/http"
)
var users= map[string]string{
	"user1": "password1",
	"user2": "password2",
}

type Credentials struct {
	Username string   'json:"username"'
	Password string  'json:"password"'
};

type Claims struct {

}

var jwtkey= []byte("secret_key")

func Login(w http.ResponseWriter, r *http.Request){

}

func Home(w http.ResponseWriter, r *http.Request){

}

func Refresh(w http.ResponseWriter, r *http.Request){

}