package git

import (
	"log"
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5"
)

// FindGitDir finds the
func FindGitDir(dir string) (string, error) {
	dir, err := filepath.Abs(dir)
	switch {
	case err == nil:
		// success
	case os.IsNotExist(err):
		// invalid directory
		return "", git.ErrRepositoryNotExists
	default:
		// other errors
		return "", err
	}

	for {
		log.Println("dir:", dir)
		name := filepath.Join(dir, git.GitDirName)
		_, err := os.Stat(name)
		switch {
		case err == nil:
			// found
			return dir, nil
		case !os.IsNotExist(err):
			// other errors
			return "", err
		}

		// up one level
		next := filepath.Dir(dir)
		log.Println("dir:", dir, "->", next)
		if dir == next {
			break
		}

		dir = next
	}

	return "", git.ErrRepositoryNotExists
}

// OpenRepository ...
func OpenRepository(dir string) (*git.Repository, error) {
	dir, err := FindGitDir(dir)
	if err != nil {
		return nil, err
	}

	return git.PlainOpen(dir)
}
