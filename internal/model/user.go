package model

import (
	"GobbleGame/internal/vars"
)

type User struct {
	Id                int
	Username          string
	Email             string
	Password          string
	EncryptedPassword string
}

func (u *User) ToResponse() *vars.Response {
	return &vars.Response{
		Type: "user",
		Id:   u.Id,
		Attributes: vars.UserAttributes{
			Username: u.Username,
			Email:    u.Email,
		},
	}
}
