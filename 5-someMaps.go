package someMaps

import (
	"fmt"
)

func introMaps() {
	fmt.Println("maps")

	// effectively just key value pairs
	statePopulations := map[string]int{
		"CA": 98754243,
		"TX": 738246,
		"FL": 324623987,
		"OH": 78467826,
	}

	fmt.Println(statePopulations)
	statePopulations["MI"] = 6372837

	fmt.Println(statePopulations, "Add MI!")
	delete(statePopulations, "FL") // can only delete 1 at a time
	fmt.Println(statePopulations, "Delete FL")

	// update a value using
	statePopulations["CA"] = 34570000
	fmt.Println(statePopulations, "update CA")

	// // you can use anything as a key
	// fmt.Println("You can use anything as a key!")
	// m := map[[3]int]string{} // default empty map with
	// fmt.Println(m)
	// fmt.Printf("Length of m: %v\n", len(m))

	// m[[3]int{1, 2, 4}] = "hello"
	// fmt.Println(m)
	// fmt.Printf("Length of m: %v\n", len(m))
}

func theOkaySyntax() {

	// effectively just key value pairs
	statePopulations := map[string]int{
		"CA": 98754243,
		"TX": 738246,
		"FL": 324623987,
		"OH": 78467826,
	}
	pop, ok := statePopulations["Ca"]
	_, okay := statePopulations["Ca"] // use _ if you just want to check presence
	pop1, ok1 := statePopulations["CA"]

	fmt.Println("Ca", pop, ok, okay)
	fmt.Println("CA", pop1, ok1)

	// maps are pass by reference, so if you assign another var to a map, you modify both
	sp := statePopulations
	delete(sp, "CA")
	pop1, ok1 = statePopulations["CA"]
	fmt.Println("CA", pop1, ok1)
}

func runAll() {
	introMaps()
	theOkaySyntax()
}
