package main

import "github.com/spf13/cobra"

var describeCmd = &cobra.Command{
	Use:   "describe",
	Short: "Give an object a human readable name based on an available ref",
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
	rootCmd.AddCommand(describeCmd)
}
