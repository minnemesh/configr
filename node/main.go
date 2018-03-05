package main

import (
	"fmt"
	"log"
	"os"

	"github.com/minnemesh/configr/common/types"
	"github.com/minnemesh/configr/node/config"
	"github.com/minnemesh/configr/node/fetcher"
	"github.com/urfave/cli"
)

type State struct {
}

func (state *State) cmdGet(args *cli.Context) error {
	nodeConfig, err := config.ReadConfig()
	if err != nil {
		err := fmt.Errorf("Error loading node config: %v", err)
		return err
	}

	app := args.String("name")
	fmt.Printf("Getting config for app '%v'\n", app)
	appconfig, present := nodeConfig.Config[app]
	if !present {
		return fmt.Errorf("Could not find app %v in node config.", app)
	}

	encConfig, err := state.fetchEncryptedAppConfig(&appconfig)
	if err != nil {
		return fmt.Errorf("Could not fetch config: %v", err)
	}

	fmt.Println(string(encConfig.Data))

	return nil
}

func (state *State) fetchEncryptedAppConfig(appconfig *config.AppConfig) (types.EncryptedAppConfig, error) {
	for _, fetchconfig := range appconfig.Fetch {
		if fetchconfig.Method == "http" {
			fetcher := fetcher.HTTPFetcher{}
			encConfig, err := fetcher.Fetch(&fetchconfig)
			if err != nil {
				log.Printf("Error fetching config via %v. %v", fetchconfig.Method, err)
				continue
			} else {
				log.Println("Fetched config via ", fetchconfig.Method)
				return encConfig, nil
			}
		} else {
			log.Printf("Unknown fetch method '%v' encountered", fetchconfig.Method)
		}
	}

	return types.EncryptedAppConfig{}, fmt.Errorf("All configured fetch options failed.")
}

func main() {
	s := State{}

	app := cli.NewApp()
	app.Name = "configr"
	app.Usage = "fetch decentralized configuration data"
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{
		{
			Name:   "get",
			Usage:  "get a configuration",
			Action: s.cmdGet,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:   "name",
					Usage:  "name of the config to get",
					Value:  "default",
					EnvVar: "CONFIG_NAME",
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
