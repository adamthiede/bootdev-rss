package main

import (
	"fmt"
	"time"
	"github.com/adamthiede/bootdev-rss/internal/database"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type User struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Name string `json:"name"`
}

func (db database) CreateUser(name string) (User, error) {
	newUserUUID:=uuid.New()
	newUserTime:=time.Now()
	newUser := CreateUserParams{
		ID: newUserUUID,
		CreatedAt: newUserTime,
		UpdatedAt: newUserTime,
		Name: name,
	}
	fmt.Printf("Added user id %v: %s\n", name, newUserUUID)
	return newUser, nil
}
