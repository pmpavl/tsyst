package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `json:"id" bson:"_id"`
	Email         string             `json:"email" bson:"email"`
	EmailVerified bool               `json:"emailVerified" bson:"emailVerified"`
	Password      string             `json:"password" bson:"password"`
	Salt          string             `json:"salt" bson:"salt"`
	Meta          *Meta              `json:"meta,omitempty" bson:"meta,omitempty"`
	Tokens        *Tokens            `json:"tokens,omitempty" bson:"tokens,omitempty"`
}

type Meta struct {
	FirstName string `json:"firstname" bson:"firstname"`
	LastName  string `json:"lastname" bson:"lastname"`
	School    uint64 `json:"school" bson:"school"`
	Class     uint64 `json:"class" bson:"class"`
}

type Tokens struct {
	Access  string `json:"access" bson:"access"`
	Refresh string `json:"refresh" bson:"refresh"`
}
