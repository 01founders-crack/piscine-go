package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		data, err := os.ReadFile(os.Args[1])
		if err != nil {
			fmt.Print("ERROR: " + err.Error())
		}
		fmt.Print(string(data))
	} else if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
	} else if len(os.Args) < 2 {
		fmt.Println("File name missing")
	}
}
