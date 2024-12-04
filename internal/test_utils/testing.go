package test_utils

import (
	"os"
	"testing"
)

func RunMultiWithoutSTDOUT(t *testing.T, testFn func(t *testing.T), iterations int) {
	// Save the original stdout
	origStdout := os.Stdout
	defer func() { os.Stdout = origStdout }()

	// Redirect stdout
	_, wOut, _ := os.Pipe()
	os.Stdout = wOut

	for i := 0; i < iterations; i++ {
		testFn(t)
	}
}
