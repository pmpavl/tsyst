package models

import "github.com/google/uuid"

type UserTokens struct {
	AccessToken  string `json:"accessToken" bson:"accessToken"`   // ACCESS_TOKEN
	RefreshToken string `json:"refreshToken" bson:"refreshToken"` // REFRESH_TOKEN
}

func NewUserTokens() *UserTokens {
	return &UserTokens{
		AccessToken:  uuid.NewString(),
		RefreshToken: uuid.NewString(),
	}
}

func (t *UserTokens) Refresh() { t.AccessToken = uuid.NewString() }
