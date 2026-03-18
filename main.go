package main

import (
	"context"

	"github.com/hugocbb/url-encurtador/internal/config"
)

func main() {
	ctx := context.Background()
	config.NewClientRedis(ctx)

}
