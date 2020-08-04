package idMaker

import (
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"strconv"
	"src/dbConnection"
)

func Maker(user_name string) {
	dbc := dbConnection.Connector()

	//gets the last id row
	last_id, err := dbc.Query("SELECT user_id, LAST_VALUE(user_id) OVER (ORDER BY user_id) last_id FROM user_data")

	if err != nil {
		panic(err.Error())
	}

	defer last_id.Close()

	var uid int
	var xp int
	//transform row to int, not sure how the 2 args work!!!
	for last_id.Next() {
		err = last_id.Scan(&uid, &xp)
	}

	//generates the id number compared to the rest in the db, NEEDS FIX FOR DELETED USERS!!!
	if xp == 0 {
		xp = 1
	} else {
		xp++
	}

	var nr string = strconv.Itoa(xp)
	fmt.Println(nr)

	id, err := dbc.Query("UPDATE user_data SET user_id = "+nr+" WHERE user_name = '"+user_name+"'")

	if err != nil {
		panic(err.Error())
	}

	defer id.Close()
}