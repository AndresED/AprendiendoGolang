package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(4)
	var wg sync.WaitGroup
	numbers := []uint32{1245445, 564564, 6546546, 4564546, 454564, 23132, 654, 321321, 459649, 1111111}
	wg.Add(len(numbers))
	fmt.Println("Comienza Proceso")
	for _, v := range numbers {
		go func(a uint32) {
			defer wg.Done()
			fmt.Println(a, esPrimo(a))
		}(v)

	}
	wg.Wait()
	fmt.Println("Termina Proceso")
}

func esPrimo(a uint32) bool {
	c := 0
	var i uint32
	for i = 1; i <= a; i++ {
		if a%1 == 0 {
			c++
		}
	}
	if c == 2 {
		return true
	}
	return false
}
