package main

import (
	"github.com/spf13/cobra"
)

var tagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Create, list, delete a tag object",
	RunE: func(cmd *cobra.Command, args []string) error {
		opts, _, err := NewGitOptions(cmd.Flags(), args)
		if err != nil {
			return err
		}

		_, _, err = opts.Open()
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(tagCmd)
}
