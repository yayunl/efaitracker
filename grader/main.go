package main

import (
	"bufio"
	"efaitracker/greeting"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Type something here: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Print("Error")
	}

	fmt.Println("Greeting: ", greeting.GreetingWord)
	fmt.Println("User input: ", input)
}
