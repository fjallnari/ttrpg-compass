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
	Title       string `redis:"title"`
	Edition     string `redis:"edition"`
	Description string `redis:"desc"`
	Url         string `redis:"url"`

	Complexity    int `redis:"cmpx"`
	Progression   int `redis:"prgr"`
	Narrative     int `redis:"narr"`
	Combat        int `redis:"comb"`
	Exploration   int `redis:"expl"`
	Balance       int `redis:"blnc"`
	Versatility   int `redis:"vers"`
	Customization int `redis:"cust"`

	Genre string `redis:"genre"`
	Type  string `redis:"type"`
	Gm    string `redis:"gm"`
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
	entries, err := os.ReadDir("../data/mock")

	if err != nil {
		panic(err)
	}

	for _, entry := range entries {
		if filepath.Ext(entry.Name()) == ".toml" {
			doc, err := os.ReadFile(fmt.Sprintf("../data/mock/%s", entry.Name()))

			if err != nil {
				panic(err)
			}

			err = toml.Unmarshal([]byte(doc), &system)

			if err != nil {
				panic(err)
			}

			systemId := fmt.Sprintf("system:%s", strings.Split(entry.Name(), ".")[0])

			if _, err := client.Pipelined(ctx, func(rdb redis.Pipeliner) error {
				rdb.HSet(ctx, systemId, "title", system.Title)
				rdb.HSet(ctx, systemId, "edition", system.Edition)
				rdb.HSet(ctx, systemId, "desc", system.Description)
				rdb.HSet(ctx, systemId, "url", system.Url)

				rdb.HSet(ctx, systemId, "cmpx", system.Complexity)
				rdb.HSet(ctx, systemId, "prgr", system.Progression)
				rdb.HSet(ctx, systemId, "narr", system.Narrative)
				rdb.HSet(ctx, systemId, "comb", system.Combat)
				rdb.HSet(ctx, systemId, "expl", system.Exploration)
				rdb.HSet(ctx, systemId, "blnc", system.Balance)
				rdb.HSet(ctx, systemId, "vers", system.Versatility)
				rdb.HSet(ctx, systemId, "cust", system.Customization)

				rdb.HSet(ctx, systemId, "genre", system.Genre)
				rdb.HSet(ctx, systemId, "type", system.Type)
				rdb.HSet(ctx, systemId, "gm", system.Gm)
				return nil
			}); err != nil {
				panic(err)
			}

			// if err != nil {
			// 	panic(err)
			// }
			fmt.Printf("%s	\033[32mOK\033[0m\n", system.Title)
		}
	}
}
