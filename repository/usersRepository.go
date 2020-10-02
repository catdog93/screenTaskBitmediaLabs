package repository

import "os"

var (
	usersCollection = os.Getenv("USERS_COLLECTION")
)
