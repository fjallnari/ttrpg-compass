package main

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func searchHandler(c *fiber.Ctx) error {
	var systems []TTRPGSystem = make([]TTRPGSystem, 0)
	titleFilter := ""
	genreFilter := ""
	familyFilter := ""

	if c.Params("title") != "*" {
		titleFilter = fmt.Sprintf("@title:%s*", strings.ReplaceAll(c.Params("title"), "_", " "))
		c.SendString(titleFilter)
		return c.SendStatus(418)
	}

	if c.Params("genre") != "*" {
		genreFilter = fmt.Sprintf(" @genre:%s", c.Params("genre"))
		c.SendString(genreFilter)
		return c.SendStatus(418)
	}

	if c.Params("family") != "*" {
		familyFilter = fmt.Sprintf(" @family:%s", strings.ReplaceAll(c.Params("family"), "_", " "))
		c.SendString(familyFilter)
		return c.SendStatus(418)
	}

	res, err := dbClient.Do(dbCtx, "FT.SEARCH", "idx:systems", fmt.Sprintf("%s%s%s", titleFilter, genreFilter, familyFilter), "NOCONTENT").Result()

	if err != nil {
		c.SendString(fmt.Sprintf("%s", err))
		return c.SendStatus(404)
	}

	foundSystems := res.([]interface{})[1:]

	for _, systemKey := range foundSystems {
		var system TTRPGSystem
		if err := dbClient.HGetAll(dbCtx, fmt.Sprintf("%s", systemKey)).Scan(&system); err != nil {
			c.SendString(fmt.Sprintf("%s", systemKey))
			return c.SendStatus(500)
		}

		system.Id = strings.Split(fmt.Sprintf("%s", systemKey), ":")[1]
		system.Similar = getTitlesOnly(getSimilarSystems(system.Id))

		systems = append(systems, system)
	}

	return c.JSON(systems)
}
