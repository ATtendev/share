package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/qeesung/image2ascii/convert"

	"github.com/ATtendev/share/config"
	"github.com/ATtendev/share/internal/log"
	"github.com/ATtendev/share/server"
	"github.com/ATtendev/share/store/db"
	"github.com/ATtendev/share/store/geo"
	"github.com/urfave/cli/v2"
)

func init() {
	log.Init("debug")
}

func main() {
	var cfgPath string
	app := &cli.App{
		Name:  "Share",
		Usage: "make an explosive entrance",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "config",
				Usage:       "setup config",
				Destination: &cfgPath,
				Required:    true,
			},
		},
		Action: func(*cli.Context) error {

			cfgs, err := config.NewConfig(cfgPath)
			if err != nil {
				log.Panicf("config.NewConfig error: %v", err)
			}

			// Create convert options
			convertOptions := convert.DefaultOptions
			convertOptions.FixedWidth = 100
			convertOptions.FixedHeight = 40

			// Create the image converter
			converter := convert.NewImageConverter()
			fmt.Print(converter.ImageFile2ASCIIString(cfgs.ServerConf.Banner, &convertOptions))

			sigChan := make(chan os.Signal, 1)
			signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
			// connect to geo database
			geoDB := geo.NewConnection()
			// connect to database
			storeDB := db.NewConnection(&cfgs.DatabaseConf)

			s, _ := server.NewServer(context.Background(), cfgs, storeDB, geoDB)
			go func() {
				sig := <-sigChan
				log.Infof("force shutdown %s", sig)
				s.Shutdown(context.Background())
			}()
			s.Start()
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Errorf("Can't start share %s", err.Error())
	}
}
