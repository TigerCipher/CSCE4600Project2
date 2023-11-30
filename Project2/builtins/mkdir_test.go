package builtins

import (
	"errors"
	"os"
	"testing"
)

// testing "return err" part of mkdir.go to achieve 100% was aided by chatgpt

var mockMkdir func(string, os.FileMode) error

func TestMakeDirectory(t *testing.T) {
	// Test case 1: Single argument
	err := MakeDirectory("testdir")
	if err != nil {
		t.Errorf("MakeDirectory returned an error: %v", err)
	}
	os.Remove("testdir")

	// Test case 2: Multiple arguments
	err = MakeDirectory("dir1", "dir2")
	if err != ErrInvalidArgCount {
		t.Errorf("MakeDirectory did not return the expected error for multiple arguments. Expected: %v, Got: %v", ErrInvalidArgCount, err)
	}

	// Test case 3: Verify directory creation
	err = MakeDirectory("testdir")
	if err != nil {
		t.Errorf("MakeDirectory returned an error: %v", err)
	}
	//os.Remove("testdir")

	// Test case 4: Verify error handling
	expectedErr := errors.New("mkdir testdir: Cannot create a file when that file already exists.")
	expectedErrUnix := errors.New("mkdir testdir: file exists")
	mockMkdir = func(path string, perm os.FileMode) error {
		return expectedErr
	}
	err = MakeDirectory("testdir")
	if err.Error() != expectedErr.Error() && err.Error() != expectedErrUnix.Error() {
		t.Errorf("MakeDirectory did not return the expected error. Expected: %v, Got: %v", expectedErr, err)
	}
	mockMkdir = nil
	os.Remove("testdir")
}
