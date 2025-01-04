package main

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func main() {
	serveHTTP(context.Background())
	select {}
}

func serveHTTP(_ context.Context) {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("ok"))
		log.Printf("health check, err:%v", err)
	})
	http.ListenAndServe(":9080", nil)
}
