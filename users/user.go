package users

import (
	"fmt"
	"time"

	"go.mod/models"
)

func ShowUsers() {
	now := time.Now()
	birthDate := now.AddDate(-29, 0, 0)

	u := models.User{
		Name:      "Juan",
		LastName:  "Daniel",
		Email:     "juanda55@gmail.com",
		BirthDate: birthDate,
		Status:    true,
		CreatedAt: now,
		Id:        5}

	user := new(models.User)
	user.AddUser(u)
	fmt.Println(user)

	user.AddUser(models.User{
		Email:     "jiren@gmail.com",
		Name:      "Jiren",
		LastName:  "Toppo",
		BirthDate: now.AddDate(-30, 0, 0),
		Status:    true,
		CreatedAt: now,
		Id:        6})
	fmt.Println(user)
}
