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

	rebuildDB(dbClient, dbCtx, "../data/mock/")

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/api/systems", func(c *fiber.Ctx) error {
		var system TTRPGSystem
		systems := make([]TTRPGSystem, 0)
		iter := dbClient.Scan(dbCtx, 0, "*", 0).Iterator()

		for iter.Next(dbCtx) {
			if err := dbClient.HGetAll(dbCtx, iter.Val()).Scan(&system); err != nil {
				panic(err)
			}

			systems = append(systems, system)
		}

		if err := iter.Err(); err != nil {
			panic(err)
		}

		return c.JSON(systems)
	})

	app.Get("/api/system/:system", func(c *fiber.Ctx) error {
		responseChannel := make(chan TTRPGSystem)
		go getSystem(responseChannel, c.Params("system"))
		system := <-responseChannel

		return c.JSON(system) //c.SendString(fmt.Sprintf("%v", system))
	})

	app.Static("/", "../client/dist")

	app.Listen(":5000")
}
