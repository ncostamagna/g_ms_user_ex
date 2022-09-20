package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/ncostamagna/g_ms_user_ex/internal/user"
	"github.com/ncostamagna/g_ms_user_ex/pkg/bootstrap"
	"github.com/ncostamagna/g_ms_user_ex/pkg/handler"
	"github.com/ncostamagna/g_ms_user_ex/pkg/logger"
)

func main() {

	_ = godotenv.Load()
	db, err := bootstrap.DBConnection()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	var userSrv user.Service
	{
		log := logger.New("users")
		repo := user.NewRepo(db, log)
		userSrv = user.NewService(log, repo)
	}

	h := handler.NewUserHTTPServer(ctx, user.MakeEndpoints(userSrv))
	port := os.Getenv("PORT")

	srv := &http.Server{
		Handler:      accessControl(h),
		Addr:         fmt.Sprintf("127.0.0.1:%s", port),
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  4 * time.Second,
	}

	errCh := make(chan error)

	go func() {
		errCh <- srv.ListenAndServe()
	}()

	err = <-errCh
	if err != nil {
		log.Fatal(err)
	}

}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, HEAD")
		w.Header().Set("Access-Control-Allow-Headers", "Accept,Authorization,Cache-Control,Content-Type,DNT,If-Modified-Since,Keep-Alive,Origin,User-Agent,X-Requested-With")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}
