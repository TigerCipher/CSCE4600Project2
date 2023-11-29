package builtins

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func Echo(args ...string) error {
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
			fmt.Print(value + " ")
		} else {
			// Print the argument
			fmt.Print(arg + " ")
		}
	}

	// Print the end character
	fmt.Print("\n")

	return nil
}
