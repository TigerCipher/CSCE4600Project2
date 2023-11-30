package builtins

import (
	"bytes"
	"os"
	"testing"
)

func TestCat(t *testing.T) {
	// Create a buffer to capture the output
	output := bytes.Buffer{}

	// Test case 1: Single file
	err := Cat(&output, "cd.go")
	if err != nil {
		t.Errorf("Cat returned an error: %v", err)
	}

	expectedOutput, err := os.ReadFile("cd.go")
	if err != nil {
		t.Errorf("Failed to read expected output file: %v", err)
	}

	if !bytes.Equal(output.Bytes(), expectedOutput) {
		t.Errorf("Cat produced incorrect output for single file. Expected: %q, Got: %q", expectedOutput, output.Bytes())
	}

	// Test case 2: Multiple files
	output.Reset()
	err = Cat(&output, "cd.go", "echo.go")
	if err != nil {
		t.Errorf("Cat returned an error: %v", err)
	}

	expectedOutput, err = os.ReadFile("cd.go")
	if err != nil {
		t.Errorf("Failed to read expected output file: %v", err)
	}

	expectedOutput2, err := os.ReadFile("echo.go")
	if err != nil {
		t.Errorf("Failed to read expected output file: %v", err)
	}

	expectedOutput = append(expectedOutput, expectedOutput2...)

	if !bytes.Equal(output.Bytes(), expectedOutput) {
		t.Errorf("Cat produced incorrect output for multiple files. Expected: %q, Got: %q", expectedOutput, output.Bytes())
	}
}
