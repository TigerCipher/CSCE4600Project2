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

		count := int64(0)
		count, _ = io.Copy(w, f)
		if count == 0 {
			return fmt.Errorf("failed to read file")
		}
	}

	return nil
}
