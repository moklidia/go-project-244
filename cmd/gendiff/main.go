package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:      "gendiff",
		Usage:     "Compares two configuration files and shows a difference.",
		ArgsUsage: "<file1> <file2>",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "format",
				Aliases: []string{"f"},
				Usage:   "output format",
				Value:   "stylish",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			if cmd.Args().Len() < 2 {
				return fmt.Errorf("error: requires 2 arguments\n\nUsage: %s", cmd.ArgsUsage)
			}
			file1 := cmd.Args().Get(0)
			file2 := cmd.Args().Get(1)
			data1, err := os.ReadFile(file1)
			if err != nil {
				log.Fatal(err)
			}
			data2, err := os.ReadFile(file2)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("Comparing %s and %s\n", data1, data2)

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
