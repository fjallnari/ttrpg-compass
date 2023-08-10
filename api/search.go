package main

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func searchHandler(c *fiber.Ctx) error {
	var systems []TTRPGSystem = make([]TTRPGSystem, 0)
	genreFilter := ""
	titleFilter := ""

	if c.Params("title") != "*" {
		titleFilter = fmt.Sprintf("@title:%s*", strings.ReplaceAll(c.Params("title"), "_", " "))
	}

	if c.Params("genre") != "*" {
		genreFilter = fmt.Sprintf(" @genre:%s", c.Params("genre"))
	}

	res, err := dbClient.Do(dbCtx, "FT.SEARCH", "idx:systems", fmt.Sprintf("%s%s", titleFilter, genreFilter), "NOCONTENT").Result()

	if err != nil {
		return c.SendStatus(404)
	}

	foundSystems := res.([]interface{})[1:]

	for _, systemKey := range foundSystems {
		var system TTRPGSystem
		if err := dbClient.HGetAll(dbCtx, fmt.Sprintf("%s", systemKey)).Scan(&system); err != nil {
			panic(err)
		}

		system.Id = strings.Split(fmt.Sprintf("%s", systemKey), ":")[1]
		system.Similar = getTitlesOnly(getSimilarSystems(system.Id))

		systems = append(systems, system)
	}

	return c.JSON(systems)
}
