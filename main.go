package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

const (
	maxKeyLength   = 16
	maxValueLength = 128
)

func main() {

	kb := flag.Int("kb", 0, "size in KB to generate (required)")
	filePath := flag.String("file", "", "output file path (required)")
	flag.Parse()

	if *kb <= 0 || *filePath == "" {
		flag.Usage()
		os.Exit(1)
	}

	if !strings.HasSuffix(*filePath, ".json") {
		*filePath += ".json"
	}

	data := generateRandomJSON(*kb * 1024)

	err := os.WriteFile(*filePath, data, 0644)
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
		os.Exit(1)
	}
}

func generateRandomJSON(sizeInBytes int) []byte {
	data := make(map[string]string)
	currentSize := 0

	for currentSize < sizeInBytes {
		key := randomString(maxKeyLength)
		value := randomString(maxValueLength)

		data[key] = value

		// Check current size of the JSON data
		jsonData, err := json.Marshal(data)
		if err != nil {
			fmt.Println("Error generating JSON:", err)
			os.Exit(1)
		}

		currentSize = len(jsonData)
		if currentSize >= sizeInBytes {
			break
		}
	}

	finalJSON, _ := json.Marshal(data)
	return finalJSON
}

func randomString(maxLen int) string {
	length := rand.Intn(maxLen) + 1
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
