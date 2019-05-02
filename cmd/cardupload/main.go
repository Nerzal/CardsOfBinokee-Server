package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"

	"github.com/urfave/cli"
	resty "gopkg.in/resty.v1"
)

func main() {
	log.Println("Attempting to upload card database..")
	app := cli.NewApp()
	app.Version = "0.1"
	app.Usage = "Upload cards"
	app.Commands = []cli.Command{
		{
			Name:    "post",
			Aliases: []string{"p"},
			Usage:   "Upload json",
			Action: func(c *cli.Context) error {
				path := c.String("path")
				log.Println("Upload path is: " + path)
				log.Println("Attempting to read file...")

				cardBytes, err := ioutil.ReadFile("cards.json")
				if err != nil {
					log.Println(err)
					return err
				} else {
					log.Println("Successfully read card.json")
				}

				resp, err := resty.R().
					SetBody(cardBytes).
					SetHeader("Content-Type", "application/json").
					Post(path)

				if err != nil {
					log.Println(err)
					return err
				}

				if resp.IsError() {
					log.Println("Failed to upload json")
					log.Println(string(resp.Body()))
					return nil
				}

				return nil
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "path, p",
					Value: "localhost:8090/api/v1/cards",
					Usage: "Path to post the json to",
				},
				cli.StringFlag{
					Name:  "json, s",
					Value: "",
					Usage: "json to post",
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully uploaded the cards!")
	}

	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}
