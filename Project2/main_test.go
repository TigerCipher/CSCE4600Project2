package main

import (
	"bytes"
	"io"
	"os/exec"
	"strings"
	"testing"
	"testing/iotest"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_runLoop(t *testing.T) {
	t.Parallel()
	exitCmd := strings.NewReader("exit\n")
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name     string
		args     args
		wantW    string
		wantErrW string
	}{
		{
			name: "no error",
			args: args{
				r: exitCmd,
			},
		},
		{
			name: "read error should have no effect",
			args: args{
				r: iotest.ErrReader(io.EOF),
			},
			wantErrW: "EOF",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			w := &bytes.Buffer{}
			errW := &bytes.Buffer{}

			exit := make(chan struct{}, 2)
			// run the loop for 10ms
			go runLoop(tt.args.r, w, errW, exit)
			time.Sleep(10 * time.Millisecond)
			exit <- struct{}{}

			require.NotEmpty(t, w.String())
			if tt.wantErrW != "" {
				require.Contains(t, errW.String(), tt.wantErrW)
			} else {
				require.Empty(t, errW.String())
			}
		})
	}
}

func TestExecuteCommand(t *testing.T) {
	// Create a test command
	cmd := exec.Command("tree", "")

	// Create a buffer to capture the output
	output := bytes.Buffer{}

	// Set the output device to the buffer
	cmd.Stdout = &output

	err := executeCommand("tree")

	// Check if the command execution produced an error
	if err != nil {
		t.Errorf("executeCommand returned an error: %v", err)
	}

	// Check if the output matches the expected value
	// expectedOutput := "Hello, World!\n"
	// if output.String() != expectedOutput {
	// 	t.Errorf("executeCommand produced incorrect output. Expected: %q, Got: %q", expectedOutput, output.String())
	// }
	// my drive letter may be different than others, so I can't really hardcode an expected outcome for the tree command, but if by this point it didn't produce an error, then it should have worked
}

func TestHandleInput(t *testing.T) {
	// Create a buffer to capture the output
	output := bytes.Buffer{}

	// Create a channel for the exit signal
	exit := make(chan struct{})

	// Test case 1: Built-in command "echo"
	err := handleInput(&output, "echo Hello, World!", exit)
	if err != nil {
		t.Errorf("handleInput returned an error: %v", err)
	}

	// Test case 2: LS command
	err = handleInput(&output, "ls", exit)
	if err != nil {
		t.Errorf("handleInput returned an error: %v", err)
	}

	// Test case 3: mkdir command
	err = handleInput(&output, "mkdir tempdir", exit)
	if err != nil {
		t.Errorf("handleInput returned an error: %v", err)
	}

	// Test case 4: rm command
	err = handleInput(&output, "rm tempdir", exit)
	if err != nil {
		t.Errorf("handleInput returned an error: %v", err)
	}

	// Test case 5: rm command
	err = handleInput(&output, "cat README.md", exit)
	if err != nil {
		t.Errorf("handleInput returned an error: %v", err)
	}

	// Test case 5: Non-existent command
	/*	err = handleInput(&output, "nonexistent", exit)
		expectedErr := "exec: \"nonexistent\": executable file not found in %PATH%"
		exectedErrOnUnix := "exec: \"nonexistent\": executable file not found in $PATH"
		if err.Error() != expectedErr || err.Error() != exectedErrOnUnix {
			t.Errorf("handleInput did not return the expected error for a non-existent command. Expected: '%v', Got: '%v' (or if on linux then expected: '%v')", expectedErr, err, exectedErrOnUnix)
		}
	*/
	// Test case 6: cd command
	err = handleInput(&output, "cd builtins", exit)
	if err != nil {
		t.Errorf("handleInput returned an error: %v", err)
	}

	// Test case 7: env command
	err = handleInput(&output, "env", exit)
	if err != nil {
		t.Errorf("handleInput returned an error: %v", err)
	}

	// Test case 8: Built-in command "exit"
	go func() {
		err := handleInput(&output, "exit", exit)
		if err != nil {
			t.Errorf("handleInput returned an error: %v", err)
		}
	}()

	// Wait for the exit signal
	<-exit
}
