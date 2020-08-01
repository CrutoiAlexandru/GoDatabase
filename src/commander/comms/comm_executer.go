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

func Deleter() int{
	fmt.Printf("Name to delete: ")
	reader := Scn()

	//delete all data then return as there are no more commands to be given
	if reader == "-all" {
		db_commands.DeleteAll()
		return 0
	}

	//delete data from db
	if reader == "close" {
		return 0
	}

	//end of function where if there is no other command given a specific row of db will be deleted
	db_commands.DeleteRow(reader)
	Deleter()
	return 0
}

func Shower() {
	fmt.Println("These are the names: ")

	db_commands.ShowNames()
}