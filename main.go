package main

import (
	"context"
	"ex-go/internal/app"
)

func main() {

	ctx := context.Background()

	app.Start(ctx)
}
