package builtins

import (
	"os"
	"testing"
)

func TestListFiles(t *testing.T) {
	// Create a temporary directory for testing
	tempDir := "testdir"
	err := os.Mkdir(tempDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create temporary directory: %v", err)
	}
	defer os.Remove(tempDir)

	// Create some test files
	testFiles := []string{"file1.txt", "file2.txt", "file3.txt"}
	for _, filename := range testFiles {
		file, err := os.Create(tempDir + "/" + filename)
		if err != nil {
			t.Fatalf("Failed to create test file: %v", err)
		}
		file.Close()
		defer os.Remove(tempDir + "/" + filename)
	}

	// Test listing files in the directory
	err = ListFiles(tempDir)
	if err != nil {
		t.Fatalf("Failed to list files: %v", err)
	}

	// Test listing files in a non-existent directory
	err = ListFiles("nonexistentdir")
	if err == nil {
		t.Fatalf("Expected error when listing files in non-existent directory")
	}

	// Test listing files with reverse order
	err = ListFiles("-r", tempDir)
	if err != nil {
		t.Fatalf("Failed to list files with reverse order: %v", err)
	}

	// Test listing files with long format
	err = ListFiles("-l", tempDir)
	if err != nil {
		t.Fatalf("Failed to list files with long format: %v", err)
	}

	// Test listing files with hidden files shown
	err = ListFiles("-a", tempDir)
	if err != nil {
		t.Fatalf("Failed to list files with hidden files shown: %v", err)
	}

	// Test listing files sorted by size
	err = ListFiles("-S", tempDir)
	if err != nil {
		t.Fatalf("Failed to list files sorted by size: %v", err)
	}

	// Test listing files with human-readable sizes
	err = ListFiles("-h", "-l", tempDir)
	if err != nil {
		t.Fatalf("Failed to list files with human-readable sizes: %v", err)
	}

}

func TestPrintFileInfo(t *testing.T) {
	// Create a temporary file for testing
	tempFile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	// Get the file info
	fileInfo, err := tempFile.Stat()
	if err != nil {
		t.Fatalf("Failed to get file info: %v", err)
	}

	// Call the printFileInfo function
	printFileInfo(fileInfo, tempFile.Name(), true)
}

func TestFormatSize(t *testing.T) {
	testCases := []struct {
		size         int64
		expectedSize string
	}{
		{0, "0 B"},
		{1023, "1023 B"},
		{1024, "1.0 KB"},
		{1048576, "1.0 MB"},
		{1073741824, "1.0 GB"},
		{1099511627776, "1.0 TB"},
		{1125899906842624, "1.0 PB"},
	}

	for _, tc := range testCases {
		actualSize := formatSize(tc.size)
		if actualSize != tc.expectedSize {
			t.Errorf("Unexpected size.\nInput: %d\nExpected: %s\nActual: %s", tc.size, tc.expectedSize, actualSize)
		}
	}
}
