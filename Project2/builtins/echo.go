package builtins

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func Echo(w io.Writer, args ...string) error {
	if len(args) < 1 {
		return ErrInvalidArgCount
	}

	// Parse command-line flags
	flag.Parse()

	// Iterate over the arguments
	for _, arg := range args {
		// Check if the argument starts with a "$" indicating an environment variable
		if strings.HasPrefix(arg, "$") {
			// Get the environment variable value
			envVar := strings.TrimPrefix(arg, "$")
			value := os.Getenv(envVar)

			// Print the environment variable value
			fmt.Print(w, value+" ")
		} else {
			// Print the argument
			fmt.Print(w, arg+" ")
		}
	}

	// Print the end character
	fmt.Print(w, "\n")

	return nil
}
