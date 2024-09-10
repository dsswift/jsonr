package main

import (
	"log/slog"
	"os"
	"os/exec"
	"strconv"
	"testing"
)

func TestGenerateFile(t *testing.T) {

	var primes = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}

	for i := 0; i < 10; i++ {

		kb := 10 * primes[i]

		slog.Info("Generating and testing a new file", "kb", kb)

		outputFile := "test_output.json"

		err := exec.Command("go", "run", "main.go", "--kb", strconv.Itoa(kb), "--file", outputFile).Run()
		if err != nil {
			t.Fatalf("Failed to run the application: %v", err)
		}

		if _, err := os.Stat(outputFile); os.IsNotExist(err) {
			t.Fatalf("Output file was not created: %s", outputFile)
		}

		fileInfo, err := os.Stat(outputFile)
		if err != nil {
			t.Fatalf("Failed to get file info: %v", err)
		}

		expectedSize := kb * 1024
		fileSize := int(fileInfo.Size())

		if fileSize < expectedSize-128 || fileSize > expectedSize+128 { // 128 bytes = 0.125 kb
			t.Fatalf("File size does not match expected size. Got %d bytes, expected ~%d bytes", fileSize, expectedSize)
		}

		os.Remove(outputFile)
	}
}
