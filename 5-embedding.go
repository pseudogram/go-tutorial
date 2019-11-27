package embedding

import (
	"fmt"
)

type Animal struct {
	Name   string
	Origin string
}

// you can embed the Animal struct in the Bird.
// This is NOT Inheritance.
// And the animal is NOT a field. This is Go magic.
type Bird struct {
	Animal   // Embedded
	SpeedKPH float32
	CanFly   bool
}

type Magpie struct {
	Bird
	isACunt bool
}

func introToEmbedding() {
	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Austrailia"
	b.SpeedKPH = 32
	b.CanFly = false
	fmt.Println(b) // notice when printed embedded structs print inside {}

	// You can use multiple embedding
	// But you can't initilze an embedded struct using a struct literal like so below.

	// m := Magpie{
	// 	Name:     "Billy the Magpie",
	// 	Origin:   "everywhere",
	// 	SpeedKPH: 28,
	// 	CanFly:   true,
	// 	isACunt:  true,
	// }

	m := Magpie{}
	m.Name = "Billy the Magpie"
	m.Origin = "everywhere"
	m.SpeedKPH = 28
	m.CanFly = true
	m.isACunt = true
	fmt.Println(m)

	// instead you can you must declare the indide struct

	b2 := Bird{
		Animal:   Animal{Name: "Another Emu", Origin: "Austrailia"},
		SpeedKPH: 34,
		CanFly:   false,
	}
	fmt.Println(b2)
}

func tags() {
	// to use tags you need to import the reflect package
	type Car struct {
		Make     string `required max:"10"`
		SpeedKPH float32
		wheels   int8
	}

	// tags are actually nothing but a strin of text.
	m := reflectTypeOf(Car{})
	field, _ := m.FieldByName("Name")
	fmt.Println(field.Tag)
	// so once accessed, they actually will need to be handled by another validation framework
}

func RunAll() {
	introToEmbedding()
	tags()
}
