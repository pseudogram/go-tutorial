package someconstants

// import built in or external packages
import (
	"fmt"
)

func constants() {
	fmt.Println("Constants")
	const j int = 42 // default value given as false
	// j = 78           // not allowed, obs

	// const k float64 = math.Sin(1.57) // not allowed because value is only set at runtime, consts must be set at runtime

	fmt.Printf("%v, %T\n", j, j)
	// fmt.Printf("%v, %T\n", k, k)
	// fmt.Printf("%v, %T\n", n, n)
}

// Colletion types can't be constants (arrays and maps)

const s int16 = 27

func shadowingPackageConsts() {
	fmt.Println("Shadowing Constants")
	const s int64 = 2.3E15
	fmt.Printf("%v %T\n", a, a)
}

func implicitConstants() {
	fmt.Println("Implicit Constants")
	const a = 42 // Is this an int or what?
	var b int16 = 27
	fmt.Printf("%v %T\n", a, a)
	fmt.Printf("%v %T\n", a+b, a+b) // effectively every use of 'a' is just replaced with the number 42
}

const a = iota // iota is a special symbol that is a counter...

const (
	b = iota
	c = iota
	d = iota
)

const ( // iota is scoped to a constant block an the compliler implictly figures out following values
	e = iota
	f
	g
)

const (
	errorSpecialist = iota // its good to have a default value / error value
	catSpecialist
	dogSpecialist
	rabbitSpecialist
)

const (
	_ = iota + 5 // _ use underscore when the value doesn matter
	bearSpecialist
	emuSpecialist
)

func incrementingUsingIota() {
	fmt.Println("incrementing using Iota")
	fmt.Printf("%v %T\n", a, a)

	fmt.Println("const block 1")
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)

	fmt.Println("const block 2")
	fmt.Printf("%v \n", e)
	fmt.Printf("%v \n", f)
	fmt.Printf("%v \n", g)

	fmt.Println("domestic specialists")
	fmt.Printf("%v \n", errorSpecialist)
	fmt.Printf("%v \n", catSpecialist)
	fmt.Printf("%v \n", dogSpecialist)
	fmt.Printf("%v \n", rabbitSpecialist)

	fmt.Println("exotic specialists")
	// fmt.Printf("%v \n", _) // this will not run as _ is a value that is thrown away.
	fmt.Printf("%v \n", bearSpecialist)
	fmt.Printf("%v \n", emuSpecialist)
}

const (
	_  = iota // ignore first value by assigning to black identier
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func bitShiftingAndIota() {
	fmt.Println("Bit shifting and Iota")

	fmt.Printf("KB = %v\nGB = %v\n", KB, GB)
	fileSize := 4000000000.
	fmt.Printf("%.2fGB\n", fileSize/GB)
}

const (
	isAdmin = 1 << iota
	isHeadquaters
	canSeeFinancials

	canSeeAfrica
	canSeeAisia
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica
)

func bitShiftIotaPro() {
	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("%b\n", roles)
	fmt.Printf("Is Admin? %v\n", isAdmin&roles == isAdmin)
	fmt.Printf("Is HQ? %v\n", isHeadquaters&roles == isHeadquaters)
}

func runAll() {
	constants()
	shadowingPackageConsts()
	implicitConstants()
	incrementingUsingIota()
	bitShiftingAndIota()
	bitShiftIotaPro()
}
