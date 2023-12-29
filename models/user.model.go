package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Name     string             `json:"name,omitempty" bson:"name"`
	Email    string             `json:"email,omitempty" bson:"email"`
	Password string             `json:"-" bson:"password"`
}
type Signup struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type Ws struct {
	ID      primitive.ObjectID `json:"id" bson:"_id"`
	Name    string             `json:"name,omitempty" bson:"name"`
	Creator primitive.ObjectID `json:"creator" bson:"creator"`
}
type LoginRsp struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Name     string             `json:"name,omitempty" bson:"name"`
	Email    string             `json:"email,omitempty" bson:"email"`
	Password string             `json:"-" bson:"password"`
	Ws       []Ws               `json:"ws" bson:"ws"`
}
