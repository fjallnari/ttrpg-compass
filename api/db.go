package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pelletier/go-toml/v2"
	"github.com/redis/go-redis/v9"
)

type TTRPGSystem struct {
	Title      string   `redis:"title"`
	Desc       string   `redis:"desc"`
	Url        string   `redis:"url"`
	Complexity string   `redis:"cmpx"`
	Genre      string   `redis:"genre"`
	System     string   `redis:"system"`
	Dice       string   `redis:"dice"`
	Gm         string   `redis:"gm"`
	Tags       []string `redis:"tags"`
}

func setupDBClient() (*redis.Client, context.Context) {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return client, ctx
}

// Rebuilds the database from the data directory
// Might want to add selective rebuild
func rebuildDB(client *redis.Client, ctx context.Context) {
	fmt.Printf("Building DB...\n")
	var system TTRPGSystem
	entries, err := os.ReadDir("../data")

	if err != nil {
		panic(err)
	}

	for _, entry := range entries {
		if filepath.Ext(entry.Name()) == ".toml" {
			doc, err := os.ReadFile(fmt.Sprintf("../data/%s", entry.Name()))

			if err != nil {
				panic(err)
			}

			err = toml.Unmarshal([]byte(doc), &system)

			if err != nil {
				panic(err)
			}

			systemId := strings.ReplaceAll(strings.ToLower(system.Title), " ", "")

			if _, err := client.Pipelined(ctx, func(rdb redis.Pipeliner) error {
				rdb.HSet(ctx, systemId, "title", system.Title)
				rdb.HSet(ctx, systemId, "desc", system.Desc)
				rdb.HSet(ctx, systemId, "url", system.Url)
				rdb.HSet(ctx, systemId, "cmpx", system.Complexity)
				rdb.HSet(ctx, systemId, "genre", system.Genre)
				rdb.HSet(ctx, systemId, "system", system.System)
				rdb.HSet(ctx, systemId, "dice", system.Dice)
				rdb.HSet(ctx, systemId, "gm", system.Gm)
				return nil
			}); err != nil {
				panic(err)
			}

			// if err != nil {
			// 	panic(err)
			// }
			fmt.Printf("%s		\033[32mOK\033[0m\n", system.Title)
		}
	}
}
