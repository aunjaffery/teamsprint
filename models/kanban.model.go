package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Kanban struct {
	Title      string               `json:"title"`
	Workspace  primitive.ObjectID   `json:"workspace"`
	Creator    primitive.ObjectID   `json:"creator"`
	Members    []primitive.ObjectID `json:"members"`
	Visibility string               `json:"visibility"`
}
type KanbanRsp struct {
	ID         primitive.ObjectID   `json:"id" bson:"_id"`
	Title      string               `json:"title"`
	Workspace  primitive.ObjectID   `json:"workspace"`
	Creator    primitive.ObjectID   `json:"creator"`
	Members    []primitive.ObjectID `json:"members"`
	Visibility string               `json:"visibility"`
}
