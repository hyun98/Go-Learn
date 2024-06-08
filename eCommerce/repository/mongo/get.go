package mongo

import (
	"context"
	"eCommerce/types"
	"go.mongodb.org/mongo-driver/bson"
)

func (m *Mongo) GetUserBucket(user string) (*types.User, error) {
	// find, findOne, aggregate
	filter := bson.M{"user": user}

	var u types.User
	if err := m.user.FindOne(context.Background(), filter).Decode(&u); err != nil {
		return nil, err
	} else {
		return &u, nil
	}
}

func (m *Mongo) GetContent(name string) ([]*types.Content, error) {
	filter := bson.M{}

	if name == "" {
		filter["user"] = name
	}

	ctx := context.Background()
	if cursor, err := m.content.Find(ctx, filter); err != nil {
		return nil, err
	} else {
		defer cursor.Close(ctx)
		var res []*types.Content
		if err := cursor.All(ctx, res); err != nil {
			return nil, err
		} else {
			return res, nil
		}
	}
}

func (m *Mongo) GetUserBucketHistory(user string) (*types.History, error) {
	filter := bson.M{"user": user}

	var u types.History
	if err := m.history.FindOne(context.Background(), filter).Decode(&u); err != nil {
		return nil, err
	} else {
		return &u, nil
	}
}
