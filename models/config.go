package models

import (
	"user_auth/db"
)

var server = "172.18.0.2:27017"

// var server = os.Getenv("DATABASE")
// Database name
var databaseName = "users"

// Create a connection
var dbConnect = db.NewConnection(server, databaseName)
