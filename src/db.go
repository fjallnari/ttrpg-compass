package main

import (
	"context"
	"os"
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

func feedDB(client *redis.Client, ctx context.Context) TTRPGSystem {
	var system TTRPGSystem
	doc, err := os.ReadFile("../data/cairn.toml")

	if err != nil {
		panic(err)
	}

	err = toml.Unmarshal([]byte(doc), &system)

	if err != nil {
		panic(err)
	}

	systemId := strings.ToLower(system.Title)

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

	return system
}
