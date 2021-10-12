package main

import "fmt"

/*
this is
a multi line comment

*/

// singe line comment

func main() {

	fmt.Println("Simple string") // you CANNOT use single quotes for strings '<string>'
	fmt.Println(`
This is a multi-line. \n
String
`)
	fmt.Println('L')
}
