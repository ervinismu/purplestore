package model

type User struct {
	ID             int    `db:"id"`
	Username       string `db:"username"`
	Email          string `db:"email"`
	HashedPassword string `db:"hashed_password"`
}
