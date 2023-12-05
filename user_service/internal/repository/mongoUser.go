package repository

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/spf13/viper"
	"github.com/tiltedEnmu/puregrade-user/internal/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoUser struct {
	client *mongo.Client
}

func NewMongoUser(client *mongo.Client) *MongoUser {
	return &MongoUser{client: client}
}

func (r *MongoUser) Create(user entities.User) error {
	coll := r.client.Database(viper.GetString("mongo.db")).Collection("users")

	doc := bson.D{
		{"_id", user.Id},
		{"email", user.Email},
		{"username", user.Username},
		{"password", user.Password},
		{"avatar", user.Avatar},
		{"createdAt", user.CreatedAt},
		{"roles", user.Roles},
	}
	_, err := coll.InsertOne(context.TODO(), doc)

	return err
}

func (r *MongoUser) Get(username string) (entities.User, error) {
	coll := r.client.Database(viper.GetString("mongo.db")).Collection("users")
	filter := bson.D{{"username", username}}

	var user entities.User

	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		log.Println("Mongo Get(): Find user error: ", err.Error())
		return user, err
	}

	if err = cursor.Decode(&user); err != nil {
		log.Println("Mongo Get(): Decode user error: ", err.Error())
	}

	return user, err
}

func (r *MongoUser) GetById(id int64) (entities.User, error) {
	coll := r.client.Database(viper.GetString("mongo.db")).Collection("users")
	filter := bson.D{{"_id", id}}

	var user entities.User

	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		log.Println("Mongo Get(): Find user error: ", err.Error())
		return user, err
	}

	if err = cursor.Decode(&user); err != nil {
		log.Println("Mongo Get(): Decode user error: ", err.Error())
	}

	return user, err
}

func (r *MongoUser) CheckUserRole(id, role int64) (bool, error) {
	coll := r.client.Database(viper.GetString("mongo.db")).Collection("users")
	filter := bson.M{"_id": id, "roles": bson.M{"$in": role}}
	opts := options.Count().SetLimit(1)

	isExists, err := coll.CountDocuments(context.TODO(), filter, opts)
	if err != nil {
		log.Println("Mongo CheckUserRole(): Find user error: ", err.Error())
		return false, err
	}

	return isExists != 0, err
}

func (r *MongoUser) AddFollower(id, publisherId int64) error {
	coll := r.client.Database(viper.GetString("mongo.db")).Collection("users")
	filter := bson.D{{"_id", publisherId}}
	update := bson.D{{"$push", bson.M{"followers": id}}}

	_, err := coll.UpdateOne(context.TODO(), filter, update)

	return err
}

func (r *MongoUser) DeleteFollower(id, publisherId int64) error {
	coll := r.client.Database(viper.GetString("mongo.db")).Collection("users")
	filter := bson.D{{"_id", publisherId}}
	update := bson.D{{"$pull", bson.M{"followers": id}}}

	_, err := coll.UpdateOne(context.TODO(), filter, update)

	return err
}

func (r *MongoUser) Delete(id int64) error {
	coll := r.client.Database(viper.GetString("mongo.db")).Collection("users")
	filter := bson.D{{"_id", id}}

	_, err := coll.DeleteOne(context.TODO(), filter)

	return err
}
