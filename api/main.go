package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
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

	//defer dbClient.Close()

	rebuildDB(dbClient, dbCtx, "data/mock/")

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173, http://localhost:3000, " + os.Getenv("CLIENT_URL"),
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/api/systems/all/:cursor", systemsHandler)

	app.Get("/api/test/title::title/genre::genre/family::family", func(c *fiber.Ctx) error {
		titleFilter := ""
		genreFilter := ""
		familyFilter := ""
		if c.Params("title") != "*" {
			titleFilter = fmt.Sprintf("@title:%s*", strings.ReplaceAll(c.Params("title"), "_", " "))
		}

		if c.Params("genre") != "*" {
			genreFilter = fmt.Sprintf(" @genre:%s", c.Params("genre"))
		}

		if c.Params("family") != "*" {
			familyFilter = fmt.Sprintf(" @family:%s", strings.ReplaceAll(c.Params("family"), "_", " "))
		}

		res, err := dbClient.Do(dbCtx, "FT.SEARCH", "idx:systems", fmt.Sprintf("%s%s%s", titleFilter, genreFilter, familyFilter), "NOCONTENT").Result()

		if err != nil {
			c.SendString(fmt.Sprintf("%s", err))
			return c.SendStatus(404)
		}

		return c.SendString(fmt.Sprintf("%s", res))
	})

	app.Get("/api/systems/search/title::title/genre::genre/family::family", searchHandler)

	app.Get("/api/systems/similar/:id", func(c *fiber.Ctx) error {
		similarSystems := getSimilarSystems(c.Params("id"))

		for i, system := range similarSystems {
			similarSystems[i].Similar = getTitlesOnly(getSimilarSystems(system.Id))
		}

		return c.JSON(similarSystems)
	})

	app.Get("/api/system/:system", func(c *fiber.Ctx) error {
		responseChannel := make(chan TTRPGSystem)
		go getSystem(responseChannel, c.Params("system"))
		system := <-responseChannel

		return c.JSON(system)
	})

	//app.Static("/", "../client/build")

	log.Fatal(app.Listen(":5000"))
}
