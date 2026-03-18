package v1

import (
	"fmt"
	"net/http"
)

func healt(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte(`{"status":"ok"}`))
}

func HandlerRequest() {
	http.HandleFunc("/", healt)

	fmt.Println("Servidor iniciado em http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
