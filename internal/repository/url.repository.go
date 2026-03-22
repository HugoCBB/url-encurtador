package repository

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"github.com/hugocbb/url-encurtador/internal/models"
	"github.com/redis/go-redis/v9"
)

func Save(ctx context.Context, rdb *redis.Client, input models.UrlRecord) error {

	data, err := json.Marshal(input)
	if err != nil {
		return err
	}

	expiration := time.Hour * 24

	if err := rdb.Set(ctx, input.Id, data, expiration).Err(); err != nil {
		return err

	}

	return nil
}

func GetByShortCode(ctx context.Context, rdb *redis.Client, shortCode string) (models.UrlRecord, error) {
	url, err := rdb.Get(ctx, shortCode).Result()
	if err == redis.Nil {
		return models.UrlRecord{}, fmt.Errorf("url não encontrada")
	}

	if err != nil {
		return models.UrlRecord{}, fmt.Errorf("erro ao buscar url: %v", err)
	}

	var record models.UrlRecord
	err = json.Unmarshal([]byte(url), &record)
	if err != nil {
		return models.UrlRecord{}, fmt.Errorf("erro ao decodificar dados: %v", err)
	}
	return record, nil
}

func GenerateShortCodeUrl(url string) string {
	hasher := md5.New()
	hasher.Write([]byte(url))
	hashBytes := hasher.Sum(nil)

	hashHex := hex.EncodeToString(hashBytes)
	return hashHex[:8]
}
