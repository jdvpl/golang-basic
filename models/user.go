package models

import "time"

type User struct {
	Id        int
	Name      string
	LastName  string
	Email     string
	BirthDate time.Time
	Status    bool
	CreatedAt time.Time
}

func (u *User) AddUser(user User) {
	u.Id = user.Id
	u.Name = user.Name
	u.LastName = user.LastName
	u.Email = user.Email
	u.BirthDate = user.BirthDate
	u.Status = user.Status
	u.CreatedAt = user.CreatedAt
}
