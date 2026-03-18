package main

import (
	"context"

	v1 "github.com/hugocbb/url-encurtador/cmd/api/v1"
	"github.com/hugocbb/url-encurtador/internal/config"
)

func main() {
	ctx := context.Background()
	config.NewClientRedis(ctx)
	v1.HandlerRequest()
}
