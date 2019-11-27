package somecollections

import (
	"fmt"
)

func arrays() {
	fmt.Println("basics of arrays")
	grades := [3]int{34, 64, 74}               // arrays can only hold one type of data, but it can be any type
	grades2 := [...]int{34, 64, 74, 34, 2, 72} // ... can be used to initialize an array with an arbitrary number of elements

	var students [5]string
	students[0] = "Lisa"
	students[1] = "Margo"
	students[2] = "Ahmed"

	fmt.Printf("Grades: %v\n", grades)
	fmt.Printf("Grades2: %v\n", grades2)
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Number of Potential students: %v\n", len(students))
}

func multipleDimentions() {
	fmt.Println("Multiple dimensions")

	var idMatrix [3][3]int
	idMatrix[0] = [3]int{1, 0, 0}
	idMatrix[1] = [3]int{0, 1, 0}
	idMatrix[2] = [3]int{0, 0, 1}

	fmt.Printf("Identity matrix:\n %v\n", idMatrix)
}

func passingAroundArrays() {
	fmt.Println("Passing around arrays")
	// arrays are treated like values, so assigning an array to another varialbe creates a copy of the first object, not a reference
	a := [...]int{1, 2, 3}
	b := a   // a is copied to b
	b[1] = 5 // so this won't affect the data in a

	c := &a // using ampersand means c is now a pointer to a
	c[2] = 8

	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("c: %v\n", c)

}

func slices() {
	fmt.Println("Playing with slices")
	// slices are like arrays.
	a := []int{1, 2, 3} // but you don't decalare the size
	b := a              // a is NOT copied to b but by default points to the same data
	b[1] = 5

	fmt.Printf("Type a: %T\n", a)
	fmt.Printf("Length of a: %v\n", len(a))
	fmt.Printf("Capacity of a: %v\n", cap(a))
	fmt.Printf("b: %v\n", b)

	c := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	d := c[:]   // all elements
	e := c[3:]  // 4th element onwards
	f := c[:6]  // slice the first 6 elements
	g := c[3:6] // 4th, 5th and 6th elements
	fmt.Printf("c: %v\t\tlen %v\t\tcap %v\n", c, len(c), cap(c))
	fmt.Printf("d: %v\t\tlen %v\t\tcap %v\n", d, len(d), cap(d))
	fmt.Printf("e: %v\t\tlen %v\t\tcap %v\n", e, len(e), cap(e))
	fmt.Printf("f: %v\n", f)
	fmt.Printf("g: %v\n", g)
}

func theMakeFunction() {
	fmt.Println("The Make Function")

	a := make([]int, 3, 100) // creates a slice with 3 default values, but with a max capacity of 100 (the underlying array is size 100)
	fmt.Println(a)
	fmt.Printf("Length of a: %v\n", len(a))
	fmt.Printf("Capacity of a: %v\n", cap(a))
}

func slicesPart2() {
	fmt.Println("Slices part 2")

	b := []int{} // creates an empty slice
	fmt.Println(b)
	fmt.Printf("Length of b: %v\n", len(b))
	fmt.Printf("Capacity of b: %v\n", cap(b))

	b = append(b, 1) // creates a new slice if capacity reached. doubling the size of the underlying array
	fmt.Println(b)
	fmt.Printf("Length of b: %v\n", len(b))
	fmt.Printf("Capacity of b: %v\n", cap(b))

	b = append(b, 2, 3, 4, 5, 6) // all subsequent values added individually to array
	fmt.Println(b)
	fmt.Printf("Length of b: %v\n", len(b))
	fmt.Printf("Capacity of b: %v\n", cap(b))

	// b = append(b, []int{7,8,9}) // doesnt work
	b = append(b, []int{7, 8, 9}...) // elipsis separates out the slice into individual vals
	fmt.Println(b)
	fmt.Printf("Length of b: %v\n", len(b))
	fmt.Printf("Capacity of b: %v\n", cap(b))
}

func slicesPart3() {
	fmt.Println("Slices part 2")

	a := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(a)
	fmt.Printf("Length of a: %v\n", len(a))
	fmt.Printf("Capacity of a: %v\n", cap(a))

	b := append(a[:2], a[3:]...) // b is a slice of a, and refers n-1 elements
	fmt.Println(b)
	fmt.Println(a) // so a is modified
}

func runAll() {
	arrays()
	multipleDimentions()
	passingAroundArrays()
	slices()
	theMakeFunction()
	slicesPart2()
	slicesPart3()
}
