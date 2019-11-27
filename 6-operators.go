package SomeOperators

import (
	"fmt"
	"math"
)

func ifStatements() {
	fmt.Println("if Statements")

	if true {
		fmt.Println("The test is true")
	} // always add curly braces

	// initialiser syntax
	statePopulations := map[string]int{
		"CA": 98754243,
		"TX": 738246,
		"FL": 324623987,
		"OH": 78467826,
	}

	if pop, ok := statePopulations["OH"]; ok { // assign pop and ok, check ok, use pop in statement. Dope.
		fmt.Println(pop)
	}

}

func equalityOperators() {
	fmt.Println("Equality Operators")
	number := 50
	guess := 0

	fmt.Println(number < guess, number > guess, number <= guess, number >= guess, number == guess, number != guess)

	// Short Circuiting example below. (set guess to 0) -> works it || and &&
	if guess < 1 || returnTrue() || guess > 100 { // OR,, Aslo go short circuits, and only evaluates the first statement if there is an or test and it returns true
		fmt.Println("Guess must be between 1 and 100!")
	} else { // AND

		if guess < number {
			fmt.Println("Too Low")
		} else if guess > number {
			fmt.Println("Too High")
		} else {
			fmt.Println("Bang on!")
			fmt.Println(!false) // will print true
		}

	}
}

func carefulWithFloats() {
	fmt.Println("Careful with Floats")
	num := 0.123456789

	compare := num == math.Pow(math.Sqrt(num), 2)
	fmt.Printf("%v, squaring and rooting a number does not alwasy return the same number\n", compare)
	// careful when using floating point numbers in conditional statements.
	maxError := 0.001
	if math.Abs(num/math.Pow(math.Sqrt(num), 2)-1) < maxError {
		fmt.Printf("These numbers are the same, with max error of %v\n", maxError)
	}
}

func returnTrue() bool {
	fmt.Println("Returning True!")
	return true
}

func switchStatements() {
	fmt.Println("Switch Statements")
	switch 5 {
	case 7:
		fmt.Println("Case 7")
	case 1, 2, 3: // no such thing as falling through, but there are multiple cases
		fmt.Println("Case 1,2,3")
	case 4, 6: // cant add the same case twice in a switch statement, here e.g. 7
		fmt.Println("Case 4,6")
	default:
		fmt.Println("There's no case for you...")
	}

	switch d := 2 + 2; d {
	case 5, 6, 7:
		fmt.Println("Big number bro")
	case 1, 2, 3:
		fmt.Println("You simmered down too much")
		// default:
		// 	fmt.Println("not been thought about")     // you don't need a default., and this will just be ignored
	}

	// the tagless Syntax!!!
	i := 4 * 2
	switch {
	case i < 20:
		fmt.Println("Com'on, bru")
		fallthrough // break thwas the cause of so many errors that each case breaks by default
		// fallthough has been added to the switch if you want to hit multiple cases
	case i < 10:
		fmt.Println("I get that you're trying")
	case i > 50:
		fmt.Println("Simmer down.")
	default:
		fmt.Println("You're alright.")
	}
}

func typeSwitching() {
	var i interface{} = [3]int{} // an interface can be any type...
	switch i.(type) {
	case int:
		fmt.Println("i is an int")
	case float32:
		fmt.Println("i is an float32")
	case string:
		fmt.Println("i is a string")
	case [3]int:
		fmt.Println("i is a [3]int, which is different to a [2]int")
		break // breaks still exist.
		fmt.Println("this wont print.")
	default:
		fmt.Println("i is another type")

	}
}

func RunAll() {
	ifStatements()
	equalityOperators()
	carefulWithFloats()
	switchStatements()
	typeSwitching()
}
