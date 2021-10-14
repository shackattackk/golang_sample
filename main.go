package main

import (
	"fmt"
	"math"
)

/*
this is
a multi line comment

*/

// singe line comment

func main() {

	//strings,digits and booleans
	fmt.Println("Simple string") // you CANNOT use single quotes for strings '<string>'
	fmt.Println(`This is a multi-line. 
String`)
	fmt.Println('L')
	fmt.Println("Addition:", 1+3)
	fmt.Println("Subtraction:", 1-5)
	fmt.Println("Multiplication:", 2*11)
	fmt.Println("Division:", 20.0/3)
	fmt.Println("Exponents:", math.Pow(20.0, 3))
	fmt.Println("Greater than:", 1 > 2)
	fmt.Println("Less than or equal to:", 1 <= 2)
	fmt.Println("Equivalent:", "yay" == "yays")
	fmt.Println("Not Equivalent", 3 != 4)

	//variables

	var name string //declare variable with data type
	var val, ok = "yes", true

	myInt := 15 //shorthand syntax of declaring an variable
	name = "iu"

	fmt.Println("myInt is:", myInt)
	fmt.Println("myInt times two:", myInt*2)
	fmt.Println("val is:", val)
	fmt.Println("ok is:", ok)
	fmt.Println("name is:", name)

	//arrays

	names := [3]string{"John", "Powers", "Hey"} //declare array of strings with 3 elements called name
	var names2 [3]string

	names2[0] = "Keith"
	names2[1] = "Kyron"
	names2[2] = "Kevin"
	fmt.Println(names)
	fmt.Println(names2)

	//slices
	cities := []string{}
	cities = append(cities, "Manila")
	cities = append(cities, "London", "San Francisco", "Los Angeles")

	fmt.Println(cities)
	fmt.Println(cities[2])

	//maps
	birthdays := map[string]string{
		"Keith": "02/01/1990",
		"John":  "01/15/2003",
		"Zen":   "10/10/2010",
	}
	fmt.Println(birthdays)
	fmt.Println("Birthday of Zen:", birthdays["Zen"])

	ages := map[string]int{}
	ages["Keith"] = 28
	ages["John"] = 12
	ages["Zen"] = 5
	delete(ages, "John") //delete an element from a map
	fmt.Println(ages)

	for _, age := range ages {
		if age > 21 {
			fmt.Println("Old")
		} else if age >= 10 && age <= 18 {
			fmt.Println("Teenager")
		} else {
			fmt.Println("young")
		}
	}
}
