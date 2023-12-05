package entities

import "time"

type Post struct {
	ID        string    `bson:"_id" json:"id"`
	AuthorID  string    `bson:"authorID" json:"authorID"`
	Title     string    `bson:"title" json:"title"`
	Body      string    `bson:"body" json:"body"`
	Tags      []string  `bson:"tags" json:"tags"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
}
