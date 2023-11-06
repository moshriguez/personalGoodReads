package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// add omitempty?
type Book struct {
	ID         primitive.ObjectID `bson:"_id" json:"id"`
	Title      string             `bson:"title" json:"title"`
	Author     string             `bson:"author" json:"author"`
	Read       bool               `bson:"read" json:"read"`
	ReadOn     time.Time          `bson:"read_on" json:"read_on"`
	LostInMove bool               `bson:"lost_in_move" json:"lost_in_move"`
	Deleted    bool               `bson:"deleted" json:"deleted"`
}
