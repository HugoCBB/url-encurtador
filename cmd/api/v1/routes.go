package v1

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hugocbb/url-encurtador/internal/config"
	"github.com/hugocbb/url-encurtador/internal/controllers"
)

func healt(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":"ok"}`))
}

func HandlerRequest() {
	ctx := context.Background()
	rdb := config.NewClientRedis(ctx)
	urlHandler := controllers.NewUrlController(rdb)

	http.HandleFunc("/", healt)
	http.HandleFunc("POST /", urlHandler.CreateUrl)
	http.HandleFunc("GET /{code}", urlHandler.Redirect)

	fmt.Println("Servidor iniciado em http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
