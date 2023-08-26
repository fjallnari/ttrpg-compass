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
	}

	if c.Params("genre") != "*" {
		genreFilter = fmt.Sprintf(" @genre:%s", c.Params("genre"))
	}

	if c.Params("family") != "*" {
		familyFilter = fmt.Sprintf(" @family:%s", strings.ReplaceAll(c.Params("family"), "_", " "))
	}

	res, err := dbClient.Do(dbCtx, "FT.SEARCH", "idx:systems", fmt.Sprintf("%s%s%s", titleFilter, genreFilter, familyFilter), "NOCONTENT").Result()
	fmt.Println(res)

	if err != nil {
		fmt.Println(err)
		return c.SendStatus(404)
	}

	foundSystems := res.([]interface{})[1:]

	for _, systemKey := range foundSystems {
		var system TTRPGSystem
		if err := dbClient.HGetAll(dbCtx, fmt.Sprintf("%s", systemKey)).Scan(&system); err != nil {
			fmt.Println(err)
			panic(err)
		}

		system.Id = strings.Split(fmt.Sprintf("%s", systemKey), ":")[1]
		system.Similar = getTitlesOnly(getSimilarSystems(system.Id))

		systems = append(systems, system)
	}

	return c.JSON(systems)
}
