package db_commands

import (
	_ "github.com/go-sql-driver/mysql"
	"src/db_commands/idMaker"
	"src/dbConnection"
)

func AddValue(val string) {
	dbc := dbConnection.Connector()

	//create variable and send data
	pickup, err := dbc.Query("INSERT INTO user_data (user_name) VALUES ('"+val+"')")

	if err != nil {
		panic(err.Error())
	}

	idMaker.Maker(val)
	
	defer pickup.Close()
}