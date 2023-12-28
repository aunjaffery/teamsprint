package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Workspace struct {
	Name    string               `json:"name,omitempty"`
	Creator primitive.ObjectID   `json:"creator"`
	Members []primitive.ObjectID `json:"members"`
}
type WsRsp struct {
	ID      primitive.ObjectID   `json:"id" bson:"_id"`
	Name    string               `json:"name,omitempty"`
	Creator primitive.ObjectID   `json:"creator"`
	Members []primitive.ObjectID `json:"members"`
}
