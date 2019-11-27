package someStructs

import (
	"fmt"
)

type Doctor struct { // if the struct is to be public ALL fields must be capitalised as well.
	number     int
	actorName  string
	companions []string
	episodes   []string
}

// type PublicDoctor struct {
// 	Number     int
// 	ActorName  string
// 	Companions []string
// 	Episodes   []string
// }

func aStruct() {
	fmt.Println("a Struct")

	aDoctor := Doctor{
		number:    3,
		actorName: "Jon Pertwee",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}

	// can also use positional initialization, but will is prone to breaking when editing
	// see below, episodes is missing and can't be inferred

	// anotherDoctor := Doctor{
	// 	3,
	// 	"Jon Pertwee",
	// 	[]string{
	// 		"Liz Shaw",
	// 		"Jo Grant",
	// 		"Sarah Jane Smith",
	// 	},
	// }

	fmt.Println(aDoctor.companions[1])
	// fmt.Println(anotherDoctor)
}

func anonymousStructs() {
	fmt.Println("anonymous structs")
	aCompanion := struct{ name string }{name: "Sarah Jane Smith"}
	fmt.Println(aCompanion)

	anotherCompanion := aCompanion
	anotherCompanion.name = "Judy Dench"
	fmt.Println(anotherCompanion)

	yetAnotherCompanion := &aCompanion
	yetAnotherCompanion.name = "Julie Ann Moore"
	fmt.Println(yetAnotherCompanion)
	fmt.Println(aCompanion)

}

func RunAll() {
	aStruct()
	anonymousStructs()
}
