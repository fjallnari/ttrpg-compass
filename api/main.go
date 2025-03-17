package main

import (
	"context"
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var dbClient *redis.Client
var dbCtx context.Context

func getSimilarSystems(systemId string) []TTRPGSystem {
	var system TTRPGSystem
	similar := make([]TTRPGSystem, 3)
	similarKeys, _, err := dbClient.SScan(dbCtx, "similar:"+systemId, 0, "", 0).Result()

	if err != nil {
		panic(err)
	}

	for i, key := range similarKeys {
		if err := dbClient.HGetAll(dbCtx, "system:"+key).Scan(&system); err != nil {
			panic(err)
		}
		similar[i] = system
		similar[i].Id = key
	}

	return similar
}

func getTitlesOnly(systems []TTRPGSystem) []string {
	titles := make([]string, 3)
	for i, system := range systems {
		titles[i] = system.Title
	}
	return titles
}

func main() {
	godotenv.Load()
	app := fiber.New()

	dbClient, dbCtx = setupDBClient()

	_, err := dbClient.Ping(dbCtx).Result() // Test connection
	if err != nil {
		log.Fatal(err)
	}

	defer dbClient.Close()

	if os.Getenv("REBUILD_DB") == "true" {
		rebuildDB(dbClient, dbCtx, "data/mock/")
	}

	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173", "http://localhost:3000", os.Getenv("CLIENT_URL")},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept"},
	}))

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	app.Get("/api/systems/all/:cursor", systemsHandler)

	app.Get("/api/systems/search/title::title/genre::genre/family::family", searchHandler)

	app.Get("/api/systems/similar/:id", func(c fiber.Ctx) error {
		similarSystems := getSimilarSystems(c.Params("id"))

		for i, system := range similarSystems {
			similarSystems[i].Similar = getTitlesOnly(getSimilarSystems(system.Id))
		}

		return c.JSON(similarSystems)
	})

	log.Fatal(app.Listen(":5000"))
}
