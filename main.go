package main

import (
	"fmt"
	"os"

	"example.com/m/v2/cli"
)

func main() {
	cli.AddCommands()

	if err := cli.RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}