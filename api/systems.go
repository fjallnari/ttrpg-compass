package main

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func systemsHandler(c *fiber.Ctx) error {
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
	}

	keys, cursor, err = dbClient.Scan(dbCtx, cursor, "system:*", 20).Result()

	if err != nil {
		panic(err)
	}

	for _, key := range keys {
		if err := dbClient.HGetAll(dbCtx, key).Scan(&system); err != nil {
			panic(err)
		}

		system.Id = strings.Split(key, ":")[1]
		system.Similar = getTitlesOnly(getSimilarSystems(system.Id))

		systems = append(systems, system)
	}

	genres := make([]string, 0)
	iterGenres := dbClient.SScan(dbCtx, "genres", 0, "", 0).Iterator()

	for iterGenres.Next(dbCtx) {
		genres = append(genres, iterGenres.Val())
	}

	families := make([]string, 0)
	iterFamilies := dbClient.SScan(dbCtx, "families", 0, "", 0).Iterator()

	for iterFamilies.Next(dbCtx) {
		families = append(families, iterFamilies.Val())
	}

	if err != nil {
		panic(err)
	}

	return c.JSON(fiber.Map{
		"cursor":   cursor,
		"systems":  systems,
		"genres":   genres,
		"families": families,
	})
}
