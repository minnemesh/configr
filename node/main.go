package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

type State struct {
}

func (state *State) cmdGet(args *cli.Context) error {
	fmt.Printf("Getting config: %v\n", args.String("name"))
	return nil
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
