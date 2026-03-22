package v1

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hugocbb/url-encurtador/internal/config"
	"github.com/hugocbb/url-encurtador/internal/controllers"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5174")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func healt(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":"ok"}`))
}

func HandlerRequest() {
	ctx := context.Background()
	rdb := config.NewClientRedis(ctx)
	urlHandler := controllers.NewUrlController(rdb)

	mux := http.NewServeMux()
	mux.HandleFunc("/", healt)
	mux.HandleFunc("POST /", urlHandler.CreateUrl)
	mux.HandleFunc("GET /{code}", urlHandler.Redirect)

	fmt.Println("Servidor iniciado em http://localhost:8080")
	http.ListenAndServe(":8080", corsMiddleware(mux))
}
