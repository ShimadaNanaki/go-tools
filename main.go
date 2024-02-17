package main

import (
	"fmt"
	"github.com/ShimadaNanaki/go-tools/modules"
)

func main() {
	var input string
	modules.ClearTerminal()
	fmt.Println("please select a script to run:")
	fmt.Println("1: cleaner")
	fmt.Println("2: ext_deleter")
	fmt.Println("3: name_randomer")
	fmt.Println("0: exit")
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	switch input {
	case "1":
		modules.FileCleaner()
	case "2":
		modules.ExtDeleter()
	case "3":
		modules.NameRandomer()
	case "0":
		fmt.Println("Bye.")
		return
	default:
		fmt.Println("Invalid input.")
	}
}
