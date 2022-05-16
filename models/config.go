package models

import (
	"os"
	"user_auth/db"
)

var server = os.Getenv("DBADDRESS")
var databaseName = os.Getenv("DBNAME")

func dbConnect() *db.session {
	param := db.NewConnection(server, databaseName)
	return param
}
