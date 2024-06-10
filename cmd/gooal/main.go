package main

import (
	"context"
	"time"

	"github.com/hay-i/gooal/internal/db"
	"github.com/hay-i/gooal/internal/router"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, dbErr := db.Initialize(ctx)
	defer func() {
		if dbErr = client.Disconnect(ctx); dbErr != nil {
			panic(dbErr)
		}
	}()
	defer cancel()

	e := router.Initialize(client)

	e.Logger.Fatal(e.Start(":1323"))
}
