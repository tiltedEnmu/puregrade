package entities

import "time"

type Post struct {
	Id        string    `bson:"_id" json:"id"`
	AuthorId  string    `bson:"authorId" json:"authorId"`
	Title     string    `bson:"title" json:"title"`
	Body      string    `bson:"body" json:"body"`
	Tags      []string  `bson:"tags" json:"tags"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
}
