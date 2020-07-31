package comms

import (
	"src/db_commands"
	"fmt"
)

func Adder() {
	fmt.Printf("Name to add: ")
	reader := Scn()

	//add data into db
	if reader != "close" {
		db_commands.AddValue(reader)
		Adder()
	}
}

func Deleter() {
	fmt.Printf("Name to delete: ")
	reader := Scn()

	//delete data from db
	if reader != "close" {
		db_commands.DeleteRow(reader)
		Deleter()
	}
}