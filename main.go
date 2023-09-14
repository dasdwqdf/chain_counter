package main

import (
	bmp "chain_counter/beatmap"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	var snap, minObjects int

	app := &cli.App{
		Name:  "chain_counter",
		Usage: "use to count the chains of a beatmap",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:        "snap",
				Usage:       "distance between the notes, measured in 1/snap of a beat",
				Value:       4,
				Aliases:     []string{"s"},
				Destination: &snap,
			},
			&cli.IntFlag{
				Name:        "min-objects",
				Usage:       "minimum of objects to be considered in a chain",
				Value:       3,
				Aliases:     []string{"m"},
				Destination: &minObjects,
			},
		},
		Action: func(context *cli.Context) error {
			beatmap := bmp.ImportBeatmapData(context.Args().First())
			fmt.Println(bmp.GetBeatmapChains(beatmap, snap, minObjects))
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
