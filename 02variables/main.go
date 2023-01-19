package main

import "fmt"

// if first character of variable is in Caps then it is a public variable, access by all

const LoginToken string = "ladfkhasdfjklb12jkb132j3b12j31j3b1l"

func main() {
	var username string = "ajay"
	fmt.Println(username)
	fmt.Printf("variable of type : %T \n",username)

	var isLoggedIn bool= true
	fmt.Println(isLoggedIn)
	fmt.Printf("variable of type : %T \n",isLoggedIn)
	
	var smallVal uint8= 255
	fmt.Println(smallVal)
	fmt.Printf("variable of type : %T \n",smallVal)
	
	var smallFloat float64= 255.3246928374692734
	fmt.Println(smallFloat)
	fmt.Printf("variable of type : %T \n",smallFloat)

	// default value and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type %T \n", anotherVariable)

	// implicit type/way to declare the variable
	var website = "try.com" //lexar declare the type according to the value store in variable
	fmt.Println(website)

	// no var way
	numberOfUser := 10 
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
}