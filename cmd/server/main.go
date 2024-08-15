package main

import (
	"context"
	"os"

	"github.com/ysomad/go-project-structure/internal/app/server"
	"github.com/ysomad/go-project-structure/internal/slogx"
)

func main() {
	ctx := context.Background()

	if err := server.Run(ctx, os.Args); err != nil {
		slogx.FatalContext(ctx, err.Error())
	}
}
