package models

import "gopkg.in/mgo.v2/bson"

// Represents a movie, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type Expense struct {
	ID         bson.ObjectId       `bson:"_id" json:"_id"`
	Title      string              `bson:"title" json:"title"`
	Amount     int32               `bson:"amount" json:"amount"`
	Date       bson.MongoTimestamp `bson:"date" json:"date"`
	Created_At bson.MongoTimestamp `bson:"created_at" json:"created_at"`
}
