package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	ID       *primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title    string              `json:"title,omitempty"`
	Subtitle string              `json:"subtitle,omitempty"`
	Author   string              `json:"author,omitempty"`
}

type UpdateBook struct {
	ID       *primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title    string              `json:"title,omitempty"`
	Subtitle string              `json:"subtitle,omitempty"`
	Author   string              `json:"author,omitempty"`
}
