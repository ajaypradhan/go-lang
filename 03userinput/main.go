package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input program"
	fmt.Println(welcome)


	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Something")
	
	// comma OK || err OK
	input, _ := reader.ReadString('\n');
	fmt.Println("thanks for entering data, ", input)
}