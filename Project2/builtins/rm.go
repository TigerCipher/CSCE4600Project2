package builtins

import (
	"os"
)

func Remove(args ...string) error {
	if len(args) < 1 {
		return ErrInvalidArgCount
	}

	for _, path := range args {
		os.RemoveAll(path)
	}

	return nil
}
