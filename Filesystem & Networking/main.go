package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("data.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalln("failed to create file:", err)
	}
	defer file.Close()
	fmt.Println(file, os.Args[1])

}
