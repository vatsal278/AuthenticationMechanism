package db

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/mgo.v2"
)

// DBConnection defines the connection structure
type session struct {
	session *mgo.Session
}

func NewDB() dbinterface.IDB {
	return &session{
		session: models.dbConnect(),
	}
}

// NewConnection handles connecting to a mongo database
func NewConnection(host string, dbName string) (conn *session) {
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

	sess.SetMode(mgo.Monotonic, true)
	conn = &session{sess}
	return conn
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
