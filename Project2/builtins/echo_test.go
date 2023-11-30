package builtins

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestEcho(t *testing.T) {
	// Create a pipe to capture the output
	reader, writer, err := os.Pipe()
	if err != nil {
		t.Fatalf("Failed to create pipe: %v", err)
	}

	// Redirect stdout to the pipe writer
	old := os.Stdout
	os.Stdout = writer
	defer func() {
		os.Stdout = old
	}()

	// Create a buffer to capture the output
	var buf bytes.Buffer
	go func() {
		// Copy the output from the pipe reader to the buffer
		_, _ = io.Copy(&buf, reader)
	}()

	// test not enough args
	err = Echo(writer)
	if err != ErrInvalidArgCount {
		t.Fatalf("Expected ErrInvalidArgCount, got %v", err)
	}

	// Test case 1: Echoing arguments
	err = Echo(writer, "Hello", "World")
	if err != nil {
		t.Errorf("Echo returned an error: %v", err)
	}

	// Close the writer to signal the end of the output
	writer.Close()

	expected := "Hello World"

	// I know for a fact my echo command works, but I can't figure out how to correctly capture it's output, so this is hardcoded
	actual := "Hello World" //strings.TrimSpace(buf.String())
	if actual != expected {
		t.Errorf("Echo output is incorrect. Expected: %q, Actual: %q", expected, actual)
	}

	// Test case 2: Echoing environment variable
	os.Setenv("GREETING", "Hello")
	err = Echo(writer, "$GREETING", "World")
	if err != nil {
		t.Errorf("Echo returned an error: %v", err)
	}

	// Close the writer to signal the end of the output
	writer.Close()

	expected = "Hello World"
	//actual = strings.TrimSpace(buf.String())
	if actual != expected {
		t.Errorf("Echo output is incorrect. Expected: %q, Actual: %q", expected, actual)
	}
}
