package response

import "go.mongodb.org/mongo-driver/bson/primitive"

type CreateResponse struct {
	ID primitive.ObjectID `json:"id"`
}
