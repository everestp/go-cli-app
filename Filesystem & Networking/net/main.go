package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	ls, err := net.Listen("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatalln("failed to lsiten on tcp 8081")
	}

	fmt.Println("server listening on  port 8081")
	for {
		conn, err := ls.Accept()
		if err != nil {
			fmt.Println("failed to connect")
			continue

		}

		defer conn.Close()

		fmt.Println("client connected :", conn.RemoteAddr().String())
		go func() {
			w := io.MultiWriter(conn, os.Stdout)
			scanner := bufio.NewScanner(conn)
			for scanner.Scan() {
				fmt.Fprintln(w, "echo:", scanner.Text())
			}

			io.Copy(w, conn)
		}()
	}
}
