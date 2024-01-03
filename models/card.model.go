package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Card struct {
	Title       string              `json:"title"`
	Description string              `json:"description"`
	Kanban      primitive.ObjectID  `json:"kanban"`
	Status      string              `json:"status"`
	Assignee    *primitive.ObjectID `json:"assignee"`
}
type CreateCard struct {
	Title       string              `json:"title"`
	Kanban      primitive.ObjectID  `json:"kanban"`
	Status      string              `json:"status"`
}
type CardRsp struct {
	ID          primitive.ObjectID  `json:"id" bson:"_id"`
	Title       string              `json:"title"`
	Description string              `json:"description"`
	Kanban      primitive.ObjectID  `json:"kanban"`
	Status      string              `json:"status"`
	Assignee    *primitive.ObjectID `json:"assignee"`
}
