package main

import (
	"fmt"
	"log"
)

func deferring() {
	fmt.Println("deferring")

	// the defer keywork runs the operation last (before calling the return statement)
	fmt.Println("start")
	defer fmt.Println("after the start") // First in Last Out
	defer fmt.Println("middle")          // last in first out.
	fmt.Println("end")

	a := "a letter"
	defer fmt.Println(a) // this is eagerly evaluated
	a = "a number"       // so this will not be printed.
}

func panicing() { // same as exceptions.
	fmt.Println("start")
	panic("something bad happened")
	fmt.Println("end")
}

func deferFirstPanicLater() {
	fmt.Println("start")
	defer fmt.Println("this was deferred, but it will execute before the panic")
	// thus in the case that you panic before closing a resource it will be closed
	panic("something bad happened")
	fmt.Println("end")
}

func panicker() {
	fmt.Println("about to panic")
	defer func() { // an anonymous function!!
		if err := recover(); err != nil {
			log.Println("Error:", err)
		}
	}() // must call the function "()" as the defer statement takes a function call
	panic("something bad happened")
	fmt.Println("end")
}

func main() {
	deferring()
	panicing()
	deferFirstPanicLater()
}
