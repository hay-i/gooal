package main

import (
	"context"
	"time"

	"github.com/hay-i/chronologger/db"
	"github.com/hay-i/chronologger/router"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, dbErr := db.Initialize(ctx)
	database := client.Database("chronologger")
	defer func() {
		if dbErr = client.Disconnect(ctx); dbErr != nil {
			panic(dbErr)
		}
	}()
	defer cancel()

	// TODO: Do we want to do this via a script instead?
	db.Seed(ctx, database)

	e := router.Initialize(client, ctx)

	e.Logger.Fatal(e.Start(":1323"))
}
