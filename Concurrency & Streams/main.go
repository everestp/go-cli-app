package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {

		defer wg.Done()
		DoSOmethingSlowly("eveest")
	}()
	go func() {

		defer wg.Done()
		DoSOmethingSlowly("paudel")
	}()

	//wait here for both goroutines
	wg.Wait()

}

func DoSOmethingSlowly(name string) {
	fmt.Println("DOing something slowly", name)
	time.Sleep(time.Second * 4)
	fmt.Println("Finsihed", name)
}
