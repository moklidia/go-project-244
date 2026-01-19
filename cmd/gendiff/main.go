package main

import (
	"code/internal/diff"
	"code/internal/formatter"
	"code/internal/parser"
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
			format := cmd.String("format")
			result, err := run(file1, file2, format)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(result)

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(file1, file2, format string) (string, error) {
	data1, err := os.ReadFile(file1)
	if err != nil {
		return "", err
	}
	data2, err := os.ReadFile(file2)
	if err != nil {
		return "", err
	}
	parsedData1, err := parser.Parse(string(data1))
	if err != nil {
		return "", fmt.Errorf("error parsing %s: %w", file1, err)
	}

	parsedData2, err := parser.Parse(string(data2))
	if err != nil {
		return "", fmt.Errorf("error parsing %s: %w", file2, err)
	}

	diffData := diff.GenerateDiff(parsedData1, parsedData2)

	formattedDiff := formatter.Format(diffData, format)

	return formattedDiff, nil
}
