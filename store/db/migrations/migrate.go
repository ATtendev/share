package main

import (
	"context"
	"flag"

	"entgo.io/ent/dialect/sql/schema"
	"github.com/ATtendev/share/config"
	"github.com/ATtendev/share/internal/log"
	"github.com/ATtendev/share/store/db/ent"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "./config.yaml", "file config")
	flag.Parse()

	c, err := config.NewConfig(configPath)
	if err != nil {
		log.Panicf("config.NewConfig error: %v", err)
	}
	db := ent.NewClient(
		ent.Log(log.Info), // logger
		ent.Driver(c.DatabaseConf.NewNoCacheDriver()),
		ent.Debug(), // debug mode
	)
	if err := db.Schema.Create(context.Background(), schema.WithForeignKeys(false), schema.WithDropColumn(true),
		schema.WithDropIndex(true)); err != nil {
		panic(err)
	}
}
