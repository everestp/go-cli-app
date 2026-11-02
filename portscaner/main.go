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
		fmt.Println("Usage: go run main.go <host>")
		os.Exit(1)
	}

	host := os.Args[1]

	ports := make([]int, 0, 65536)
	for port := range 65536 {
		ports = append(ports, port)
	}

	// Process in batches (e.g., chunks of 100)
	for _, xs := range slices.Chunk(ports, 100) {
		var wg sync.WaitGroup

		for _, port := range xs {
			wg.Add(1)
			// Pass port as parameter to avoid closure scoping bugs
			go func(p int) {
				defer wg.Done()

				conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, p), 500*time.Millisecond)
				if err != nil {
					return // Port is closed, ignore error
				}
				defer conn.Close()

				fmt.Printf("[+] Port %d: open\n", p)
			}(port)
		}

		// Wait for this chunk of 100 to finish before moving to the next batch
		wg.Wait()
	}

	fmt.Println("\nScan completed!")
}
