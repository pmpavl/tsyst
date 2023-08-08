package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `json:"-" bson:"_id,omitempty"`             // ID в mongoDB
	Email         string             `json:"email" bson:"email"`                 // Почта
	EmailVerified bool               `json:"emailVerified" bson:"emailVerified"` // Подтверждена ли почта
	Password      string             `json:"-" bson:"password"`                  // Пароль hash(password + salt)
	PasswordSalt  string             `json:"-" bson:"passwordSalt"`              // Соль пароля
	Tokens        *UserTokens        `json:"-" bson:"tokens,omitempty"`          // Токены

	CreatedAt time.Time `json:"-" bson:"createdAt"`
	UpdatedAt time.Time `json:"-" bson:"updatedAt"`
	DeletedAt time.Time `json:"-" bson:"deletedAt,omitempty"`
}

func NewUser(email, password, passwordSalt string) *User {
	now := time.Now()

	return &User{
		Email:         email,
		EmailVerified: true, // TODO: Должен быть false, true проставляется только после подтверждения почты
		Password:      password,
		PasswordSalt:  passwordSalt,
		CreatedAt:     now,
		UpdatedAt:     now,
	}
}
