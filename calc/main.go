package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	log.SetFlags(0)

	if len(os.Args) < 2 {
		log.Fatalln("missing subcommand")
	}
	subCommnd := os.Args[1]
	switch subCommnd {
	case "add":
		addCmd(os.Args[2:])
	case "sub":
		subCmd(os.Args[2:])
	default:
		log.Fatalln("invalid command")

	}
}

func addCmd(args []string) {

	if len(args) != 2 {
		log.Fatalln("incorrect arguments for add cmd")
	}

	num1, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		log.Fatalf("invalid number", err)
	}
	num2, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		log.Fatalf("invalid number", err)
	}
	sum := num1 + num2
	fmt.Println(sum)
}
func subCmd(args []string) {

	if len(args) != 2 {
		log.Fatalln("incorrect arguments for add cmd")
	}

	num1, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		log.Fatalf("invalid number", err)
	}
	num2, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		log.Fatalf("invalid number", err)
	}
	sub := num1 - num2
	fmt.Println(sub)
}
