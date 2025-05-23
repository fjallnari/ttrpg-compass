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
	Id          string
	Title       string `redis:"title"`
	Edition     string `redis:"edition"`
	Description string `redis:"desc"`
	Publisher   string `redis:"publisher"`

	Complexity    int `redis:"cmpx"`
	Progression   int `redis:"prgr"`
	Narrative     int `redis:"narr"`
	Combat        int `redis:"comb"`
	Exploration   int `redis:"expl"`
	Balance       int `redis:"blnc"`
	Versatility   int `redis:"vers"`
	Customization int `redis:"cust"`

	Genre  string `redis:"genre"`
	Family string `redis:"family"`
	Gm     string `redis:"gm"`

	Similar []string `redis:"similar"`
}

type SimilarSystem struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

func setupDBClient() (*redis.Client, context.Context) {
	ctx := context.Background()

	opts, err := redis.ParseURL(os.Getenv("REDIS_URL"))

	if err != nil {
		panic(err)
	}

	fmt.Printf("Setting up the DB client\n")

	return redis.NewClient(opts), ctx
}

// Rebuilds the database from the data directory
// Might want to add selective rebuild
func rebuildDB(client *redis.Client, ctx context.Context, dataDir string) {
	fmt.Printf("Building DB...\n")
	var system TTRPGSystem
	entries, err := os.ReadDir(dataDir)

	if err != nil {
		panic(err)
	}

	for _, entry := range entries {
		if filepath.Ext(entry.Name()) == ".toml" {
			doc, err := os.ReadFile(fmt.Sprintf("%s%s", dataDir, entry.Name()))

			if err != nil {
				panic(err)
			}

			err = toml.Unmarshal([]byte(doc), &system)

			if err != nil {
				panic(err)
			}

			systemTextId := strings.Split(entry.Name(), ".")[0]
			systemId := "system:" + systemTextId

			if _, err := client.Pipelined(ctx, func(rdb redis.Pipeliner) error {
				rdb.HSet(ctx, systemId, "title", system.Title)
				rdb.HSet(ctx, systemId, "edition", system.Edition)
				rdb.HSet(ctx, systemId, "desc", system.Description)
				rdb.HSet(ctx, systemId, "publisher", system.Publisher)

				rdb.HSet(ctx, systemId, "cmpx", system.Complexity)
				rdb.HSet(ctx, systemId, "prgr", system.Progression)
				rdb.HSet(ctx, systemId, "narr", system.Narrative)
				rdb.HSet(ctx, systemId, "comb", system.Combat)
				rdb.HSet(ctx, systemId, "expl", system.Exploration)
				rdb.HSet(ctx, systemId, "blnc", system.Balance)
				rdb.HSet(ctx, systemId, "vers", system.Versatility)
				rdb.HSet(ctx, systemId, "cust", system.Customization)

				rdb.HSet(ctx, systemId, "genre", system.Genre)
				rdb.HSet(ctx, systemId, "family", system.Family)
				rdb.HSet(ctx, systemId, "gm", system.Gm)

				return nil
			}); err != nil {
				panic(err)
			}

			// similar systems are stored in a separate set for each system
			for _, similar := range system.Similar {
				client.SAdd(ctx, "similar:"+systemTextId, similar)
			}

			client.SAdd(ctx, "genres", strings.ToLower(system.Genre))
			client.SAdd(ctx, "families", strings.ToLower(system.Family))

			fmt.Printf("%s	\033[32mOK\033[0m\n", system.Title)
		}
	}
	fmt.Printf("Loaded %d systems\n", len(entries))

	//dbClient.Do(ctx, "FT.DROPINDEX", "idx:systems")

	// _, err = dbClient.Do(ctx, "FT.CREATE", "idx:systems", "ON HASH", "PREFIX 1 system:", "SCHEMA title TEXT NOSTEM genre TEXT family TEXT").Result()
	// if err != nil {
	// 	panic(err)
	// }

	//! needs to be run manually - panic: Unknown argument `ON HASH`
	// docker exec -it redis-stack redis-cli
	// FT.CREATE idx:systems ON HASH PREFIX 1 system: SCHEMA title TEXT NOSTEM genre TEXT family TEXT
}
