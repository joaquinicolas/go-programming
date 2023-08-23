package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

func main() {

	var wg sync.WaitGroup
	words := []string{
		"alpha",
		"beta",
		"gamma",
		"delta",
		"epsilon",
		"zeta",
		"eta",
		"theta",
		"iota",
		"rho",
		"sigma",
		"tau",
		"upsilon",
		"phi",
	}

	wg.Add(len(words))

	for i, v := range words {
		go printSomething(fmt.Sprintf("Word %d: %s", i+1, v), &wg)
	}

	wg.Wait()
	wg.Add(1)
	printSomething("This is the second thing to be printed!", &wg)

}
