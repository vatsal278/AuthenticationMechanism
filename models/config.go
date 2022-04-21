package models

import (
	"os"
	"user_auth/db"
)

var server = os.Getenv("DATABASE")

// Database name
var databaseName = "users"

// Create a connection
var dbConnect = db.NewConnection(server, databaseName)
