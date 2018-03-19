package main

import (
	"log"
	"os"

	cli "github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Commands = []cli.Command{
		{
			Name:  "relayOn",
			Usage: "Turns relay on",
			Action: func(c *cli.Context) error {
				sendMqttMessage(mqttRelayPath, "relayOn")
				return nil
			},
		},
		{
			Name:  "relayOff",
			Usage: "Turns relay off",
			Action: func(c *cli.Context) error {
				sendMqttMessage(mqttRelayPath, "relayOff")
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
