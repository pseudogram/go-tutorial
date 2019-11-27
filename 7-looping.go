package looping

import (
	"fmt"
)

func basicLooping() {
	fmt.Println("Basic Looping")

	// initalizer; test; incrementer.
	for i := 0; i < 5; i++ {
		fmt.Println("simples", i)
	}

	for i, j := 0, 0; i < 5; i, j = i+1, j+2 { // take note. i++ is nt an expressin it is a statement so cant be used here...
		fmt.Println("doubles", i, j)

		// you can manipulate the counters within a loop, so be careful
		i++
	}
}

func whileLoops() {
	fmt.Println("Equivilant to while loops?")
	i := 0
	for ; i < 5; i++ { // i is no longer scoped to just the for loop.
		fmt.Println(i)
	}
	fmt.Println(i)

	// there's no such thing as while loops .... 😱
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	// equivilant to
	k := 0
	for k < 5 {
		fmt.Println(k)
		k++
	}
}

func continueBreakLabels() {

	fmt.Println("continue break labels")

	for i := 0; i < 8; i++ { // i is no longer scoped to just the for loop.
		if i%2 == 0 {
			continue // break this iteration of the loop
		}
		fmt.Println(i)
		if i == 7 {
			break // breaks out of the currently scoped loop
		}
	}

Loop: // this is a label.
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if j == 3 {
				break Loop // use labels to break out of multiple for loops
			}
		}
	}
}

func forRangeLoops() {

	fmt.Println("For Range Loops")
	s := []int{11, 22, 33, 44, 55}

	for k, v := range s {
		fmt.Println(k, v)
	}

	statePopulations := map[string]int{
		"CA": 98754243,
		"TX": 738246,
		"FL": 324623987,
		"OH": 78467826,
	}

	for k, v := range statePopulations {
		fmt.Println(k, v)
	}

	for k, v := range "A string" {
		fmt.Println(k, string(v))
	}

	// you can also use range to loop over channels (concurrent programming)
}

func RunAll() {
	basicLooping()
	whileLoops()
	continueBreakLabels()
	forRangeLoops()
}
