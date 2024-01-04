package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type List struct {
	ID    primitive.ObjectID `json:"id" bson:"_id"`
	Title string             `json:"title"`
}
type Kanban struct {
	Title      string               `json:"title"`
	Workspace  primitive.ObjectID   `json:"workspace"`
	Lists      []List               `json:"lists"`
	Creator    primitive.ObjectID   `json:"creator"`
	Members    []primitive.ObjectID `json:"members"`
	Visibility string               `json:"visibility"`
}
type KanbanRsp struct {
	ID         primitive.ObjectID   `json:"id" bson:"_id"`
	Title      string               `json:"title"`
	Workspace  primitive.ObjectID   `json:"workspace"`
	Lists      []List               `json:"lists"`
	Creator    primitive.ObjectID   `json:"creator"`
	Members    []primitive.ObjectID `json:"members"`
	Visibility string               `json:"visibility"`
}

type KanbanWithCards struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	Title      string             `json:"title"`
	Lists      []List             `json:"lists"`
	Visibility string             `json:"visibility"`
	Cards      []struct {
		ID     primitive.ObjectID `json:"id" bson:"_id"`
		Title  string             `json:"title"`
		Status string             `json:"status"`
	} `json:"cards"`
}
