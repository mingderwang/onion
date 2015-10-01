// generated by ginger from go generate -- DO NOT EDIT
/*
	MIT License
	Copyright Ming-der Wang<ming@log4analytics.com>
*/
package main

import (
	//"errors"
	"io/ioutil"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"gopkg.in/yaml.v2"
)

func defaultConfig() Config {
	config := Config{":8080", "root", "", "localhost", "/data/Onion.db", "", ""}
	return config
}

func getConfig(c *cli.Context) (Config, error) {
	yamlPath := c.GlobalString("config")
	config := Config{}

	if _, err := os.Stat(yamlPath); err != nil {
		//create default value
		config = defaultConfig()
		return config, nil
	}

	ymlData, err := ioutil.ReadFile(yamlPath)
	if err != nil {
		// create default value
		config = defaultConfig()
		return config, nil
	}

	err = yaml.Unmarshal([]byte(ymlData), &config)
	return config, err
}

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.TextFormatter{})

	// Output to stderr instead of stdout, could also be a file.
	log.SetOutput(os.Stderr)

	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)
}

func mainLoop(app *cli.App) {
	app.Action = func(c *cli.Context) {
		//host := c.GlobalString("host")
		cfg, err := getConfig(c)
		if err != nil {
			log.WithFields(log.Fields{
				"file:": "main.go",
				"func:": "main",
				"line:": 74,
			}).Fatal("getConfig return error")
			return
		}
		svc := OnionService{}

		if err = svc.Run(cfg); err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "goslack"
	app.Usage = "micro service main"
	app.Version = "0.0.3"

	app.Flags = []cli.Flag{
		//	cli.StringFlag{"host", "http://localhost:8080", "use sevice host", "APP_HOST"},
		cli.StringFlag{"config, c", "config.yaml", "config file to use", "APP_CONFIG"},
	}

	mainLoop(app)

	app.Commands = []cli.Command{
		{
			Name:  "migratedb",
			Usage: "Perform database migrations",
			Action: func(c *cli.Context) {
				cfg, err := getConfig(c)
				if err != nil {
					log.Fatal(err)
					return
				}

				svc := OnionService{}

				if err = svc.Migrate(cfg); err != nil {
					log.Fatal(err)
				}
			},
		},
	}

	app.Run(os.Args)
}
