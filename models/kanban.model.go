package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Kanban struct {
	Title      string               `json:"title"`
	Workspace  primitive.ObjectID   `json:"workspace"`
	Members    []primitive.ObjectID `json:"members"`
	Visibility string               `json:"visibility"`
}
