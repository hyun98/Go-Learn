package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// Document 가 있으면 Update, 없으면 Insert
var upsertOption = options.Update().SetUpsert(true)

func (m *Mongo) CreateUser(user string) error {
	filter := bson.M{
		"user": user,
	}

	update := bson.M{
		"$set":         bson.M{"user": user},
		"$setOnInsert": bson.M{"bucket": []string{}},
	}

	_, err := m.user.UpdateOne(context.Background(), filter, update, upsertOption)

	return err
}

func (m *Mongo) CreateContent(name string, price int64) error {
	filter := bson.M{
		"name": name,
	}

	currentTime := time.Now().Unix()

	update := bson.M{
		"$set":         bson.M{"name": name, "price": price, "updatedAt": currentTime},
		"$setOnInsert": bson.M{"createdAt": currentTime},
	}

	_, err := m.content.UpdateOne(context.Background(), filter, update, upsertOption)

	return err
}

func (m *Mongo) CreateBucket(user, content string) error {
	// $push, $addToSet
	filter := bson.M{"user": user}
	update := bson.M{"$push": bson.M{"bucket": content}}

	_, err := m.user.UpdateOne(context.Background(), filter, update, options.Update())
	return err
}

func (m *Mongo) UpsertHistory(user string, bucketList []string) error {
	filter := bson.M{"user": user}
	update := bson.M{
		"$set":         bson.M{"user": user},
		"$push":        bson.M{"contentList": bson.M{"$each": bucketList}},
		"$setOnInsert": bson.M{"createdAt": time.Now().Unix()},
	}
	_, err := m.history.UpdateOne(context.Background(), filter, update, upsertOption)
	return err
}

func (m *Mongo) RemoveUserBucket(user string) error {
	filter := bson.M{"user": user}
	update := bson.M{"$set": bson.M{"bucket": []string{}}}

	_, err := m.user.UpdateOne(context.Background(), filter, update)
	return err
}
