package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatalf("failed to connect to tcp server", err)

	}
	defer conn.Close()
	go io.Copy(conn, os.Stdin)
	io.Copy(os.Stdout, conn)

}
