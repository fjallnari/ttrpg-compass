package main

import (
	"context"
	"strconv"

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
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/api/systems/:cursor?", func(c *fiber.Ctx) error {
		var system TTRPGSystem
		var cursor uint64 = 0
		var err error
		var keys []string

		systems := make([]TTRPGSystem, 0)

		if c.Params("cursor") != "" && c.Params("cursor") != "0" {
			cursor, err = strconv.ParseUint(c.Params("cursor"), 10, 64)
			if err != nil {
				panic(err)
			}
			// fmt.Printf("Cursor: %s\n", c.Params("cursor"))
		}

		keys, cursor, err = dbClient.Scan(dbCtx, cursor, "system:*", 15).Result()

		// fmt.Printf("Cursor: %d\n", cursor)
		// fmt.Printf("Keys length: %d\n", len(keys))

		if err != nil {
			panic(err)
		}

		for _, key := range keys {
			// fmt.Printf("Key: %s\n", key)
			if err := dbClient.HGetAll(dbCtx, key).Scan(&system); err != nil {
				panic(err)
			}

			systems = append(systems, system)
		}

		return c.JSON(fiber.Map{
			"cursor":  cursor,
			"systems": systems,
		})
	})

	app.Get("/api/system/:system", func(c *fiber.Ctx) error {
		responseChannel := make(chan TTRPGSystem)
		go getSystem(responseChannel, c.Params("system"))
		system := <-responseChannel

		return c.JSON(system)
	})

	app.Static("/", "../client/dist")

	app.Listen(":5000")
}
