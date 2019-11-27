// every go application is structured in packages, you declare a files package at the begining of a file
package someprimatives

// import built in or external packages
import (
	"fmt"
)

func booleans() {
	fmt.Println("Booleans")
	var j bool // default value given as false
	var k bool = true
	n := 1 == 1

	fmt.Printf("%v, %T\n", j, j)
	fmt.Printf("%v, %T\n", k, k)
	fmt.Printf("%v, %T\n", n, n)
}

func integers() {
	fmt.Println("\nIntegers")
	// int
	// there are primative ints from int8 - int64, the default depends on your system architecture
	a := 42
	var b int8 = 127

	// uint  -- unsigned ints
	// uint8-64 also exist
	var c uint = 39200

	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)

	d := 22
	// int ops
	fmt.Println(a + d)
	fmt.Println(a - int(c))
	fmt.Println(a * d)
	fmt.Println(a / d)
	fmt.Println(a % d)
}

func intBitOperations() {
	fmt.Println("Bit operations")
	a := 10 // 1010
	b := 3  // 0011

	fmt.Println(^a) // 111...11110101 == -11 --> remember ints are signed 64 bit,
	// and when flipping all bits, you're using two's complement representation

	fmt.Println(a & b)  // 0010  AND
	fmt.Println(a | b)  // 1011  OR
	fmt.Println(a ^ b)  // 1001  XOR
	fmt.Println(a &^ b) // 0010 NAND

	c := 8              // 1000  == 2^3
	fmt.Println(c << 1) // 10000 Left Shift   == 2^3 * 2^1 == 2^4
	fmt.Println(c >> 2) // 0010 Right Shift   == 2^3 / 2^1 == 2^2
}

func bigNumbers() {
	fmt.Println("Big Nubmers")
	var n float64 = 3.14
	n = 13.7e72
	n = 13.043E10
	fmt.Printf("%v, %T", n, n)
}

func complexNumbers() {
	fmt.Println("Complex Numbers")
	// complex numbers treated as first class citizens in goLang
	a := 10.2
	b := 3.7
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	// no modulus on complex numbers

	// there are also imaginary numbers !!
	var n complex64 = 1 + 2i
	fmt.Printf("%v, %T\n", n, n)
	fmt.Printf("%v, %T\n", real(n), real(n)) // extract the real part of the number
	fmt.Printf("%v, %T\n", imag(n), imag(n)) // extract the imaginary part
	// the complex number is split into 2 float32 numbers

	// similarly complex126 is split into float64 nums
	var m complex128 = 1 + 2i
	fmt.Printf("%v, %T\n", m, m)
	fmt.Printf("%v, %T\n", real(m), real(m))
	fmt.Printf("%v, %T\n", imag(m), imag(m))

	// to make a complex number use the complex function
	var o complex128 = complex(5, 12) // first param is real, second is imaginary
	fmt.Printf("%v, %T\n", o, o)
}

func strings() {
	fmt.Println("Strings") // UTF-8

	// note strings are alias' for bytes
	// thus can be treated LIKE an array and be indexed
	// chars in a string are Uint8
	s := "this is a string"
	fmt.Printf("%v, %T\n", string(s[2]), s[2])

	s2 := " another string, haha"
	fmt.Printf("%v, %T\n", s+s2, s+s2)

	b := []byte(s)
	fmt.Printf("%v, %T\n", b, b)

	// runes
	var r rune = 'a' // it is an int32 character
	// r := 'b'
	fmt.Printf("runes -> %s is %v, %T\n", r, r, r)
}

func runAll() {
	booleans()
	integers()
	intBitOperations()
	bigNumbers()
	complexNumbers()
	strings()
}
