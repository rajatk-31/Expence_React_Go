package models

import "gopkg.in/mgo.v2/bson"

// Represents a movie, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type User struct {
	ID         bson.ObjectId `bson:"_id" json:"_id"`
	Name       string        `bson:"name" json:"name"`
	CoverImage string        `bson:"cover_image" json:"cover_image"`
	Phone      uint64        `bson:"phone" json:"phone"`
	Password   string        `bson:"password" json:"password"`
}
