package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	filename := "./lot.txt"
	log.SetFlags(0)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln("failed to read file :", err)
	}
	PrintFileContent(file)
}

func CountWords(data []byte) int {
	words := bytes.Fields(data)
	return len(words)
}

func PrintFileContent(file *os.File) {
	const bufferSize = 4096
	buffer := make([]byte, bufferSize)

	totalSize := 0

	for {
		size, err := file.Read(buffer)
		if err != nil {
			break
		}
		totalSize += size

		sunBuffer := buffer[:size]

		fmt.Print(string(sunBuffer), totalSize)

	}
}
