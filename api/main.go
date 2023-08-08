package main

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/redis/go-redis/v9"
)

var dbClient *redis.Client
var dbCtx context.Context

func getSystem(res chan TTRPGSystem, systemId string) {
	var system TTRPGSystem

	// Scan all fields into the model.
	if err := dbClient.HGetAll(dbCtx, systemId).Scan(&system); err != nil {
		panic(err)
	}

	res <- system
}

func main() {
	app := fiber.New()

	dbClient, dbCtx = setupDBClient()

	rebuildDB(dbClient, dbCtx, "../data/stress_test/")

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/api/systems/:cursor", systemsHandler)

	app.Get("/api/system/:system", func(c *fiber.Ctx) error {
		responseChannel := make(chan TTRPGSystem)
		go getSystem(responseChannel, c.Params("system"))
		system := <-responseChannel

		return c.JSON(system)
	})

	app.Static("/", "../client/dist")

	app.Listen(":5000")
}
