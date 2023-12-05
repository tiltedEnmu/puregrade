package repository

import (
	"context"
	"log"

	"github.com/tiltedEnmu/puregrade_post/internal/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoPost struct {
	client *mongo.Client
}

func NewMongoPost(mongo *mongo.Client) Post {
	return &MongoPost{client: mongo}
}

func (r *MongoPost) Create(post *entities.Post) error {
	db := r.client.Database("main")

	buf, err := bson.Marshal(post)
	if err != nil {
		return err
	}

	_, err = db.Collection("posts").InsertOne(context.TODO(), buf)

	return err
}

func (r *MongoPost) Get(id string) (*entities.Post, error) {
	db := r.client.Database("main")
	coll := db.Collection("posts")
	filter := bson.D{{"_id", id}}

	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		log.Println("Mongo Get(): Find post error: ", err.Error())
		return nil, err
	}

	var post *entities.Post
	if err = cursor.Decode(&post); err != nil {
		log.Println("Mongo Get(): Decode post error: ", err.Error())
	}

	return post, err
}

func (r *MongoPost) Delete(id string) error {
	coll := r.client.Database("main").Collection("posts")
	filter := bson.D{{"_id", id}}

	_, err := coll.DeleteOne(context.TODO(), filter)

	return err
}
