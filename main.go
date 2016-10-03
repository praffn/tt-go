package main

import (
	"os"

	"github.com/boltdb/bolt"
	"github.com/urfave/cli"

	"./trat"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	db, err := bolt.Open("tt.db", 0600, nil)
	check(err)
	defer db.Close()

	app := cli.NewApp()

	app.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: "start timer for given job",
			Action: func(c *cli.Context) error {
				trat.Start(db, c.Args().First(), "", "")
				return nil
			},
		},
		{
			Name:  "stop",
			Usage: "stop timer for given job",
			Action: func(c *cli.Context) error {
				trat.Stop(db, c.Args().First(), "", "")
				return nil
			},
		},
	}

	app.Run(os.Args)
}
