package main

import (
	"github.com/spf13/pflag"

	"github.com/amery/git-gomod/pkg/git"
)

const (
	gitDirFlag      = "git-dir"
	gitWorkTreeFlag = "work-tree"
)

func NewGitOptions(flags *pflag.FlagSet, args []string) (*git.Options, []string, error) {
	opts := new(git.Options)
	opts.LoadFlags(flags, gitDirFlag, gitWorkTreeFlag)
	opts.LoadEnv()

	switch {
	case opts.Ready():
		return opts, args, nil
	case len(args) > 0:
		err := opts.LoadFromArg(args[0])
		if err != nil {
			return nil, nil, err
		}

		// success
		return opts, args[1:], nil
	default:
		// let's hope opts.Open() detects things
		return opts, args, nil
	}
}

func init() {
	pFlags := rootCmd.PersistentFlags()
	pFlags.String(gitDirFlag, "./.git", "path to the repository (\".git\" directory)")
	pFlags.String(gitWorkTreeFlag, ".", "path to the working tree")
}
