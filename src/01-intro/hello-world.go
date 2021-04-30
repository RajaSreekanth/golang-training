// package declaration
package main // starting pint goes in main

// package imports
import fmt "fmt"

// package variables

// package init fn

// function main
func main() {
	// variable declaration
	// type1:
	// var message string
	// message = "hello world!"

	// type2:
	// var message string = "hello world!"

	// type 3:
	message := "helol worlkd"
	fmt.Println(message)

	/* type 1
	var userName string
	var greetMsg string
	userName = "sreekanth"
	greetMsg = "hava a nice day"
	*/

	/* type 2
	var userName, greetMsg string
	userName = "sreekanth"
	greetMsg = "hava a nice day"
	*/

	/* type 2
	var userName, greetMsg string = "sreekanth", "hava a nice day"
	*/

	/* type 3
	userName, greetMsg := "sreekanth", "hava a nice day"
	*/

	var (
		userName string
		greetMsg string
	)
	userName = "sreekanth"
	greetMsg = "hava a nice day"

	fmt.Println(userName, greetMsg)
}

// other types and functions
