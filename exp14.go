package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	// Allocate three logical processors
	runtime.GOMAXPROCS(3)

	// WaitGroup to wait for 3 goroutines
	var processTest sync.WaitGroup
	processTest.Add(3)

	// ------------------- Goroutine 1 : Print 51–100 -------------------
	go func() {
		defer processTest.Done()
		for i := 0; i < 30; i++ {
			for j := 51; j <= 100; j++ {
				fmt.Printf(" %d", j)
				if j == 100 {
					fmt.Println()
				}
			}
		}
	}()

	// ------------------- Goroutine 2 : Print A–Z -------------------
	go func() {
		defer processTest.Done()
		for i := 0; i < 10; i++ {
			for char := 'A'; char <= 'Z'; char++ {
				fmt.Printf("%c ", char)
			}
			fmt.Println()
		}
	}()

	// ------------------- Goroutine 3 : Print 0–50 -------------------
	go func() {
		defer processTest.Done()
		for i := 0; i < 30; i++ {
			for j := 0; j <= 50; j++ {
				fmt.Printf(" %d", j)
				if j == 50 {
					fmt.Println()
				}
			}
		}
	}()

	// Wait for all goroutines to complete
	processTest.Wait()
}
