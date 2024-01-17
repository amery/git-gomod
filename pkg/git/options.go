package git

import (
	"io/fs"
	"syscall"

	"github.com/go-git/go-git/v5"
	"github.com/spf13/pflag"
)

const (
	GitDirName     = git.GitDirName
	GitDirPrefix   = "gitdir: "
	GitDirEnv      = "GIT_DIR"
	GitWorkTreeEnv = "GIT_WORK_TREE"
)

// Options ...
type Options struct {
	GitDir      *string
	GitWorkTree *string
}

// LoadEnv loads missing GitOptions values from the environment.
func (opts *Options) LoadEnv() {
	if opts.GitDir == nil {
		if s, ok := syscall.Getenv(GitDirEnv); ok {
			opts.GitDir = &s
		}
	}

	if opts.GitWorkTree == nil {
		if s, ok := syscall.Getenv(GitWorkTreeEnv); ok {
			opts.GitWorkTree = &s
		}
	}
}

// LoadFlags reads missing GitOptions values from cobra flags.
func (opts *Options) LoadFlags(flags *pflag.FlagSet, gitDirName, gitWorkTreeName string) {
	if opts.GitDir == nil {
		if s, ok := getFlagChanged(flags, gitDirName); ok {
			opts.GitDir = &s
		}
	}

	if opts.GitWorkTree == nil {
		if s, ok := getFlagChanged(flags, gitWorkTreeName); ok {
			opts.GitWorkTree = &s
		}
	}
}

func getFlagChanged(flags *pflag.FlagSet, flagName string) (string, bool) {
	flag := flags.Lookup(flagName)
	if flag != nil && flag.Changed {
		return flag.Value.String(), true
	}

	return "", false
}

func (opts *Options) LoadFromArg(dir string) error
func (opts *Options) Ready() bool
func (opts *Options) Open() (*git.Repository, fs.FS, error)
