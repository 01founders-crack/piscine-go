package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		for _, v := range os.Args[1:] {
			if v == "01" || v == "galaxy" || v == "galaxy 01" {
				fmt.Println("Alert!!!")
				return
			}
		}
	}
}
