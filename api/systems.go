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

		system.Id = strings.Split(key, ":")[1]
		system.Similar = getTitlesOnly(getSimilarSystems(system.Id))

		// for i, similar := range similarSet {
		// 	similarSet[i] = dbClient.Get(dbCtx, "similar:"+similar).Val()
		// }

		systems = append(systems, system)
	}

	iterGenres := dbClient.SScan(dbCtx, "genres", 0, "", 0).Iterator()
	genres := make([]string, 0)

	for iterGenres.Next(dbCtx) {
		genres = append(genres, iterGenres.Val())
	}

	if err != nil {
		panic(err)
	}

	return c.JSON(fiber.Map{
		"cursor":  cursor,
		"systems": systems,
		"genres":  genres,
	})
}
