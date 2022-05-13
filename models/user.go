package models

import (
	"user_auth/Credentials"
	"user_auth/helpers"

	"gopkg.in/mgo.v2/bson"
)

// User defines user object structure
type User struct {
	Name     string `json:"name" bson:"name"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

// UserModel defines the model structure
type UserModel struct{}

// Signup handles registering a user
func (u *UserModel) Signup(cred Credentials.SignUpCredentials) error {
	// Connect to the user collection
	dbconnect := dbConnect()
	collection := dbconnect.Use(databaseName, "user")
	// Assign result to error object while saving user
	err := collection.Insert(bson.M{
		"name":     cred.Name,
		"email":    cred.Email,
		"password": helpers.GeneratePasswordHash([]byte(cred.Password)),
	})

	return err
}

// GetUserByEmail handles fetching user by email
func (u *UserModel) GetUserByEmail(email string) (user User, err error) {
	// Connect to the user collection
	dbconnect := dbConnect()
	collection := dbconnect.Use(databaseName, "user")
	// Assign result to error object while saving user
	err = collection.Find(bson.M{"email": email}).One(&user)
	return user, err
}
