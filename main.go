package main

import (
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/karaageeee/go-batch-boilerplate/config"
	"github.com/karaageeee/go-batch-boilerplate/service"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	// "github.com/karaageeee/go-batch-boilerplate/db"
)

func main() {
	setup()
	log.Info("Batch Start")
	start(os.Args)
	log.Info("Batch Finish")

}

func setup() {

	// set timezone
	time.Local = time.FixedZone("Asia/Tokyo", 9*60*60)

	// Load .env
	dotEnvErr := godotenv.Load()
	if dotEnvErr != nil {
		log.Warn("Failed to load .env file")
	}

	// set up logger
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	if os.Getenv("ENV") == "production" {
		log.SetLevel(log.InfoLevel)
	} else {
		log.SetLevel(log.DebugLevel)
	}

	// setup db
	// db.Setup()
}

// start is CLI setting and execute services
func start(args []string) {
	app := &cli.App{
		Usage: "Execute with options",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "type",
				Aliases:  []string{"t"},
				Required: true,
				Usage:    "batch type is requiered to process : demo, etc...",
			},
			&cli.StringFlag{
				Name:    "targetDate",
				Aliases: []string{"td"},
				Usage:   "[Not Required] Format : 2006-01-02.",
			},
		},
		Action: func(c *cli.Context) error {
			// name := "default"
			// if c.NArg() > 0 {
			// 	name = c.Args().Get(0)
			// }
			batchType := c.String("type")
			log.Info("Batch Type : " + batchType)

			targetDate := c.String("targetDate")
			if targetDate == "" {
				log.Info("targetDate option is empty")
			} else {
				log.Info("targetDate : " + targetDate)
			}

			switch batchType {
			case config.BatchTypeDemo:
				return service.Demo()
			}
			return nil
		},
	}

	err := app.Run(args)
	if err != nil {
		log.Fatal(err)
	}
}
