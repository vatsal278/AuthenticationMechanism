package models

import (
	"os"
	"user_auth/db"
)

var server = os.Getenv("DBADDRESS")

// Database name
var databaseName = os.Getenv("DBNAME")

// Create a connection
var dbConnect = db.NewConnection(server, databaseName)
