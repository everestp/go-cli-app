package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	addr, err := net.LookupAddr(os.Args[1])
	if err != nil {
		log.Fatalln("Failed to lookup", err)
	}
	fmt.Println(addr)
}
