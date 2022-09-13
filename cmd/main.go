package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/ncostamagna/g_ms_user_ex/internal/user"
	"github.com/ncostamagna/g_ms_user_ex/pkg/bootstrap"
)

func main() {

	router := mux.NewRouter()
	_ = godotenv.Load("../.env")
	// sin archivo y sin prefijo
	l := bootstrap.InitLogger()
	db, err := bootstrap.DBConnection()
	if err != nil {
		l.Fatal(err)
	}

	userRepo := user.NewRepo(db, l)
	userSrv := user.NewService(l, userRepo)
	userEnd := user.MakeEndpoints(userSrv)

	router.HandleFunc("/users", userEnd.Create).Methods("POST")
	router.HandleFunc("/users", userEnd.GetAll).Methods("GET")
	router.HandleFunc("/users/{id}", userEnd.Get).Methods("GET")
	router.HandleFunc("/users/{id}", userEnd.Update).Methods("PATCH")
	router.HandleFunc("/users/{id}", userEnd.Delete).Methods("DELETE")

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8081",
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  4 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}
