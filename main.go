// Package main implements the git-gomod command
package main

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var rootCmd = &cobra.Command{
	Use:   "git-gomod",
	Short: "Manage gomod tags",
}

func main() {
	err := rootCmd.Execute()
	switch err {
	case nil, pflag.ErrHelp:
		os.Exit(0)
	default:
		log.Fatal(err)
	}
}
