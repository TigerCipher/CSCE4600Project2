package builtins

import (
	"fmt"
	"io"
	"os"
)

func Cat(w io.Writer, args ...string) error {
	for _, file := range args {
		f, err := os.Open(file)
		if err != nil {
			return fmt.Errorf("failed to open file: %w", err)
		}
		defer f.Close()

		_, err = io.Copy(w, f)
		if err != nil {
			return fmt.Errorf("failed to read file: %w", err)
		}
	}

	return nil
}
