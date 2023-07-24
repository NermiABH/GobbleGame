package jwt

import (
	"GobbleGame/internal/model"
	"GobbleGame/internal/vars"
)

var secretKey = []byte("pusinu48")

type Tokens struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

func (t *Tokens) ToResponse() *vars.Response {
	return &vars.Response{
		Type:       "jwt",
		Attributes: t,
	}
}

func New(user *model.User) *Tokens {
	return &Tokens{
		Access:  newAccess(user),
		Refresh: newRefresh(),
	}
}

func newAccess(user *model.User) string {
	return ""
}

func newRefresh() string {
	return ""
}

func VerifyAccess(access string) error {
	return nil
}

func RecreateTokens(refresh string) (*Tokens, error) {
	return nil, nil
}

func verifyRefresh(refresh string) error {
	return nil
}
