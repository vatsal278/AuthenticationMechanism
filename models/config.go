package models

import (
	"user_auth/db"
)

var server = ":9002"

// var server = os.Getenv("DATABASE")
// Database name
var databaseName = "users"

// Create a connection
var dbConnect = db.NewConnection(server, databaseName)
