package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/hugocbb/url-encurtador/internal/models"
	"github.com/hugocbb/url-encurtador/internal/repository"
	"github.com/redis/go-redis/v9"
)

type UrlController struct {
	rdb *redis.Client
}

func NewUrlController(rdb *redis.Client) *UrlController {
	return &UrlController{rdb: rdb}
}

func (c *UrlController) CreateUrl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var input models.UrlRecordInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	expiration := time.Hour * 24
	shortCode := repository.GenerateShortCodeUrl(input.Url)

	data := models.UrlRecord{
		Id:        shortCode,
		ShortCode: shortCode,
		OldUrl:    input.Url,
		Exp:       expiration,
		Create_at: time.Now().Format("02/01/2006"),
	}

	if err := repository.Save(r.Context(), c.rdb, data); err != nil {
		log.Fatal("Erro ao salvar url:", err)
	}

	newUrl := "http://localhost:8080/" + shortCode
	output := models.UrlRecordOutput{
		Url: newUrl,
		Exp: data.Exp.String(),
	}

	json.NewEncoder(w).Encode(output)

}

func (c *UrlController) Redirect(w http.ResponseWriter, r *http.Request) {
	shortCode := r.PathValue("code")
	record, err := repository.GetByShortCode(r.Context(), c.rdb, shortCode)
	if err != nil {
		http.Error(w, "Link expirado ou inexistente", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, record.OldUrl, http.StatusMovedPermanently)

}
