package main

import "fmt"

func main() {

	// can't be changed
	const pi float64 = 3.14

	// multiple variables
	var (
		x = 2
		y = 3
	)
	// or use -> x, y := 2, 3

	fmt.Println(x, y)

	// Strings
	var name string = "Samuel"
	var name2 string = `Souza`

	fmt.Println(name, name2)

	// String length
	fmt.Println(len(name))

	// Data types

	// Numeric			Boolean			String		Derived
	// 	int						true				""				Pointer
	// 	float					false				``				Array
	// 																			Structure
	// 																			Map
	// 																			Interface

	// works
	a, b := 1, 2
	fmt.Println(a, ",", b)

	// not works
	// c, d = 3, 4
	// fmt.Println(c, ",", d)

	fmt.Println()
	fmt.Println(`*********************`)
	fmt.Println()

	// ****************************** OPERATORS ****************************** //
	var1, var2 := 5, 6

	fmt.Println("var1 + var2 =", var1+var2)
	fmt.Println("var1 - var2 =", var1-var2)
	fmt.Println("var1 * var2 =", var1*var2)
	fmt.Println("var1 / var2 =", var1/var2)
	fmt.Println("var1 % var2 =", var1%var2)

	fmt.Println()

	love := true
	hate := false

	fmt.Println("AND ->", love && hate)
	fmt.Println("OR  ->", love || hate)
	fmt.Println("NOT ->", !love, !hate)

	fmt.Println()
	fmt.Println(`*********************`)
	fmt.Println()

	// ****************************** POINTERS ****************************** //

	myPointer := 10

	fmt.Println(myPointer)  // Value of myPointer
	fmt.Println(&myPointer) // Address of myPointer

	fmt.Println()

	// change value of myPointer via pointer
	changeValue(&myPointer)
	fmt.Println(myPointer)

	fmt.Println()
	fmt.Println(`*********************`)
	fmt.Println()

	// ****************************** LOOPs and IFs ****************************** //

	// conventional for
	for i := 1; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println()

	// another way
	i := 1

	for i <= 10 {
		fmt.Println(i)
		i++
	}

	fmt.Println()
	fmt.Println(`*********************`)
	fmt.Println()

	var age int = 18

	if age > 18 {
		fmt.Println("You can vote")
	} else {
		fmt.Println("You can't vote")
	}

	fmt.Println()

	switch age {
	case 16:
		fmt.Println("16!")
	case 18:
		fmt.Println("18!")
	case 20:
		fmt.Println("20!")
	default:
		fmt.Println("Default case")

	}

}

func changeValue(myPointer *int) {
	*myPointer = 7
}
