package main

import (
	"encoding/json"
	"fmt"
	"log"
	"maxsol/service"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/mitchellh/mapstructure"
)

var JwtKey = []byte(os.Getenv("JWT_KEY"))

var Users = []User{
	{
		Username: "user1",
		Password: "password1",
	},
	{
		Username: "user2",
		Password: "password2",
	},
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type JwtToken struct {
	Token string `json:"token"`
}

type Exception struct {
	Message string `json:"message"`
}

type Response struct {
	Data string `json:"data"`
}

func CreateToken(w http.ResponseWriter, r *http.Request) {
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"password": user.Password,
		"exp":      time.Now().Add(time.Hour * time.Duration(1)).Unix(),
	})
	tokenString, error := token.SignedString(JwtKey)
	if error != nil {
		fmt.Println(error)
	}
	json.NewEncoder(w).Encode(JwtToken{Token: tokenString})
}

func ProtectedEndpoint(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	token, _ := jwt.Parse(params["token"][0], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error")
		}
		return JwtKey, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var user User
		mapstructure.Decode(claims, &user)
		json.NewEncoder(w).Encode(user)
	} else {
		json.NewEncoder(w).Encode(Exception{Message: "Invalid authorization token"})
	}
}

func ValidateMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, error := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("there was an error")
					}
					return JwtKey, nil
				})
				if error != nil {
					json.NewEncoder(w).Encode(Exception{Message: error.Error()})
					return
				}
				if token.Valid {
					next.ServeHTTP(w, r)
				} else {
					json.NewEncoder(w).Encode(Exception{Message: "Invalid authorization token"})
				}
			}
		} else {
			json.NewEncoder(w).Encode(Exception{Message: "An authorization header is required"})
		}
	})
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequest() {
	category := service.Category{}
	product := service.Product{}
	supplier := service.Supplier{}

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/authenticate", CreateToken).Methods("POST")
	myRouter.HandleFunc("/", ValidateMiddleware(homePage)).Methods("GET")
	myRouter.HandleFunc("/Goroutines", ValidateMiddleware(service.RunResponse)).Methods("GET")
	myRouter.HandleFunc("/Category", ValidateMiddleware(category.GetAll)).Methods("GET")
	myRouter.HandleFunc("/Category", ValidateMiddleware(category.Create)).Methods("POST")
	myRouter.HandleFunc("/Category/{id}", ValidateMiddleware(category.Update)).Methods("PUT")
	myRouter.HandleFunc("/Category/{id}", ValidateMiddleware(category.Delete)).Methods("DELETE")
	myRouter.HandleFunc("/Category/{id}", ValidateMiddleware(category.GetById)).Methods("GET")
	myRouter.HandleFunc("/Category/Name/{name}", ValidateMiddleware(category.FindByName)).Methods("GET")
	myRouter.HandleFunc("/Product", ValidateMiddleware(product.GetAll)).Methods("GET")
	myRouter.HandleFunc("/Product", ValidateMiddleware(product.Create)).Methods("POST")
	myRouter.HandleFunc("/Product/{id}", ValidateMiddleware(product.Update)).Methods("PUT")
	myRouter.HandleFunc("/Product/{id}", ValidateMiddleware(product.Delete)).Methods("DELETE")
	myRouter.HandleFunc("/Product/{id}", ValidateMiddleware(product.GetById)).Methods("GET")
	myRouter.HandleFunc("/Product/Name/{name}", ValidateMiddleware(product.FindByName)).Methods("GET")
	myRouter.HandleFunc("/Supplier", ValidateMiddleware(supplier.GetAll)).Methods("GET")
	myRouter.HandleFunc("/Supplier", ValidateMiddleware(supplier.Create)).Methods("POST")
	myRouter.HandleFunc("/Supplier/{id}", ValidateMiddleware(supplier.Update)).Methods("PUT")
	myRouter.HandleFunc("/Supplier/{id}", ValidateMiddleware(supplier.Delete)).Methods("DELETE")
	myRouter.HandleFunc("/Supplier/{id}", ValidateMiddleware(supplier.GetById)).Methods("GET")
	myRouter.HandleFunc("/Supplier/Name/{name}", ValidateMiddleware(supplier.FindByName)).Methods("GET")
	log.Fatal(http.ListenAndServe(":1234", myRouter))
}

func main() {
	fmt.Println("Starting the application...")
	handleRequest()
}
