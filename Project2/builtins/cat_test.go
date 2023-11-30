package builtins

import (
	"bytes"
	"os"
	"testing"
)

func TestCat(t *testing.T) {
	// Create a buffer to capture the output
	output := bytes.Buffer{}

	// fail to open file
	err := Cat(&output, "nonexistentfile")
	errOnWin := "failed to open file: open nonexistentfile: The system cannot find the file specified."
	errOnLinux := "failed to open file: open nonexistentfile: no such file or directory"
	if err.Error() != errOnWin && err.Error() != errOnLinux {
		t.Errorf("Cat failed to return the correct error when opening a nonexistent file. Expected: %q, Got: %q", errOnWin, err.Error())
	}

	// Fail to read file

	file, permErr := os.Create("restricted.txt")
	if permErr != nil {
		t.Fatalf("Failed to create temporary file: %v", permErr)
	}
	file.Close()

	permErr = os.Chmod("restricted.txt", 0000)
	if permErr != nil {
		t.Fatalf("Failed to change permissions on temporary file: %v", permErr)
	}

	err = Cat(&output, "restricted.txt")
	errOnWin = "failed to read file"
	errOnLinux = "failed to open file: open restricted.txt: permission denied"
	if err.Error() != errOnWin && err.Error() != errOnLinux {
		t.Errorf("Cat failed to return the correct error when reading a file. Expected: %q, Got: %q", errOnWin, err.Error())
	}

	permErr = os.Chmod("restricted.txt", 0755)
	if permErr != nil {
		t.Fatalf("Failed to change permissions on temporary file: %v", permErr)
	}

	permErr = os.Remove("restricted.txt")
	if permErr != nil {
		t.Fatalf("Failed to remove temporary file: %v", permErr)
	}

	// Test case 1: Single file
	err = Cat(&output, "cd.go")
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
