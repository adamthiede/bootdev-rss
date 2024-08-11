package main

import (
	"time"
	"github.com/adamthiede/bootdev-rss/internal/database"
	"github.com/google/uuid"
	"fmt"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
}

func CreateUser(db database.DBTX, name string) (User, error) {
	newUserUUID := uuid.New()
	newUserTime := time.Now()
	newUser := User{
		ID:        newUserUUID,
		CreatedAt: newUserTime,
		UpdatedAt: newUserTime,
		Name:      name,
	}
	fmt.Printf("Added user id %v: %s\n", name, newUserUUID)
	return newUser, nil
}
