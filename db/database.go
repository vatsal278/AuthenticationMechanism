package db

import (
	"fmt"
	"os"
	"time"
	"user_auth/models"

	"gopkg.in/mgo.v2"
)

// DBConnection defines the connection structure
type session struct {
	session *mgo.Session
}

func NewDB() models.IUserModel {
	return &models.UserModel{
		Db: dbConnect().DB(os.Getenv("DBNAME")),
	}
}

func dbConnect() *mgo.Session {
	var server = os.Getenv("DBADDRESS")
	var databaseName = os.Getenv("DBNAME")
	param := NewConnection(server, databaseName)
	return param
}

// NewConnection handles connecting to a mongo database
func NewConnection(host string, dbName string) (sess *mgo.Session) {
	info := &mgo.DialInfo{
		// Address if its a local db then the value host=localhost
		Addrs: []string{host},
		// Timeout when a failure to connect to db
		Timeout: 60 * time.Second,
		// Database name
		Database: dbName,
		// Database credentials if your db is protected
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PWD"),
	}
	fmt.Printf("%+v\n", *info)
	sess, err := mgo.DialWithInfo(info)

	if err != nil {
		panic(err)
	}

	sess.SetMode(mgo.Strong, true)

	return sess
}

// Use handles connect to a certain collection
func (conn *session) Use(dbName, tableName string) (collection *mgo.Collection) {
	// This returns method that interacts with a specific collection and table
	return conn.session.DB(dbName).C(tableName)
}

// Close handles closing a database connection
func (conn *session) Close() {
	// This closes the connection
	conn.session.Close()
}
