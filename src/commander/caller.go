package commander

import (
	"src/commander/comms"
	"fmt"
	"os"
)

var commands []string = []string{"help", "delete", "close", "add", "show", "shutdown"}

//makes sure if the input is valid or not
func validator(val string) {
	input    := val
	valid    := false

	for i:=0; i<len(commands); i++ {
		if(input != commands[i]){
			valid = false
		}else {
			valid = true
			break
		} 
	}

	if valid == false {
		fmt.Println("Wrong input! Try help command")
		CallComms()
	}
}

func CallComms() {
	fmt.Printf("Input command: ")
	input    := comms.Scn()

	validator(input)

	if input == "help" {
		//implement list
		fmt.Println("Main functions are:")
		for i:=0; i<len(commands); i++ {
			fmt.Println("	"+commands[i])
		}
		CallComms()
	}

	if input == "add" {
		comms.Adder()
		CallComms()
	}

	if input == "close" {
		fmt.Println("Application clossed!")
		os.Exit(0)
	}

	if input == "delete" {
		comms.Deleter()
		CallComms()
	}

	if input == "show" {
		comms.Shower()
		CallComms()
	}
}