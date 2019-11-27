// every go application is structured in packages, you declare a files package at the begining of a file
package variables

// import built in or external packages
import (
	"fmt"
	"strconv"
)

func firstSteps() {
	// Printing to console
	fmt.Println("Hello World!")

	// Declaring a variable (three methods)
	var i int
	i = 42
	fmt.Printf("%v, %T\n", i, i)

	var j float32 = 27
	fmt.Printf("%v, %T\n", j, j)

	k := 99 // compiler infers type
	fmt.Printf("%v, %T\n", k, k)

	l := 99.
	fmt.Printf("%v, %T\n", l, l)
}

// you can also decalre varibales at the package level
var aPersonsName = "John Doe"

// you can also decalre multiple variables in a var block
var (
	actorName = "Idris Elba"
	bestRole  = "Luther"
	worstRole = "Gunslinger in Dark Tower"
)

func playingWithVarialbes() {
	fmt.Printf("\n%s's best role has been %s, and his worst role has been %s",
		actorName, bestRole, worstRole)
}

func convertVarType() {
	var i float32 = 20.2
	var j int
	j = int(i)
	fmt.Printf("var j is of type %T, with a value of %v\n")

	// converting strings requires using the strconv package
	var k string = string(j)
	var l string = strconv.Itoa(j)
	fmt.Printf("Wrap using string _ %s  -- this wont print becasue string type conversions can be wierd\nWrap using strconv _ %s", k, l)
}

func runAll() {
	firstSteps()
	playingWithVarialbes()
	convertVarType()
}

/*
Notes:

Variable names make a difference, the first letter makes the difference.

Upper case is exported (accessed outsied the package)
lowercase is package package scoped (private)
-- there's no private scope --

Naming conventions:
- Pascal/camelCase
- vars with short lives have short names

Type convertions
*/
