package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"slices"

	"github.com/urfave/cli/v3"
)

func main() {
	cli := cli.Command{
		Name:  "cli",
		Usage: "A sample calculator",
		Flags: []cli.Flag{
			&cli.Int64Flag{
				Name:        "lhs",
				Category:    "operand",
				HideDefault: false,
				Usage:       "first (left hand side) operand",
				Required:    true,
				Aliases:     []string{"l"},
				OnlyOnce:    true,
				Validator: func(value int64) error {
					if value < 0 {
						return errors.New("should not be lower than 0")
					}
					return nil
				},
			},
			&cli.Int64Flag{
				Name:        "rhs",
				Category:    "operand",
				HideDefault: false,
				Usage:       "second (right hand side) operand",
				Required:    true,
				Aliases:     []string{"r"},
				OnlyOnce:    true,
				Validator: func(value int64) error {
					if value < 0 {
						return errors.New("should not be zero")
					}
					return nil
				},
			},
			&cli.StringFlag{
				Name:     "operator",
				Category: "operator",
				Required: true,
				Aliases:  []string{"o"},
				OnlyOnce: true,
				Validator: func(operator string) error {
					if len(operator) == 0 {
						return errors.New("operator should not be empty")
					}
					allowed := []string{"+", "-", "*", "/"}
					if !slices.Contains(allowed, operator) {
						return fmt.Errorf("operator should be %v", allowed)
					}
					return nil
				},
			},
		},
		Action: func(ctx context.Context, command *cli.Command) error {
			operator := command.String("operator")
			lhs := command.Int64("lhs")
			rhs := command.Int64("rhs")
			fmt.Println("operator:", operator)
			fmt.Println("lhs:", lhs)
			fmt.Println("rhs:", rhs)
			return nil
		},
	}
	ctx := context.Background()
	defer func() {
		fmt.Println("Goodbye")
	}()
	err := cli.Run(ctx, os.Args)
	if err != nil {
		slog.ErrorContext(ctx, "There was an error", "->", err)
		os.Exit(1)
	}
}
