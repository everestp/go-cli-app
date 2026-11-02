package main

import (
	"fmt"
	"net"
	"os"
	"slices"
	"sync"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <host>")
		os.Exit(1)
	}

	host := os.Args[1]

	ports := make([]int, 0, 65536)
	for port := range 65536 {
		ports = append(ports, port)
	}

	// Process in batches of 100 (only one variable for iter.Seq)
	for chunk := range slices.Chunk(ports, 100) {
		var wg sync.WaitGroup

		for _, port := range chunk {
			wg.Add(1)
			go func(p int) {
				defer wg.Done()

				conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, p), 500*time.Millisecond)
				if err != nil {
					return
				}
				defer conn.Close()

				fmt.Printf("[+] Port %d: open\n", p)
			}(port)
		}

		wg.Wait()
	}

	fmt.Println("\nScan completed!")
}
