package commander

import (
	"src/commander/comms"
	"fmt"
	"os"
)

func validator(val string) {
	input    := val
	commands := []string{"help","delete", "close", "add"}
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
	commands := []string{"help","delete", "close", "add"}
	
	//makes sure if the input is valid or not
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
}