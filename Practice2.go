package main

import "fmt"

func main() {
	// Data Types

	//var variable_list optional_data_type;
	var Age = 5.5 // Static
	var x float64 = 10

	Name := "Rudra" // Dynamic := meaning is var in shorthand
	a, b, c := 1, true, "A new day"

	fmt.Println(Age)
	fmt.Printf("Age type is: %T\n", Age)


	fmt.Println(Name)
	fmt.Printf("Name type is: %T\n", Name)

	fmt.Printf("a type is %T\n", a)
	fmt.Printf("b type is %T\n", b)
	fmt.Printf("c type is %T\n", c)

	fmt.Printf("x type is %T\n", x)

}

// Note: I have confussion about "Derived types".
/*
They include:
	(a) Pointer types
	(b) Array types
	(c) Structure types
	(d) Union types
	(e) Function types
	f) Slice types
	g) Interface types
	h) Map types
	i) Channel Types
*/