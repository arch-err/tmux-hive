package main

import (
	"os"

	"github.com/arch-err/tmux-hive/internal/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		os.Exit(1)
	}
}
