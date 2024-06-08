package mongo

import (
	"context"
	"eCommerce/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
	"strings"
)

type Mongo struct {
	config *config.Config
	client *mongo.Client
	db     *mongo.Database

	// TODO: 사용할 컬렉션
	user    *mongo.Collection
	content *mongo.Collection
	history *mongo.Collection
}

func NewMongo(config *config.Config) (*Mongo, error) {
	m := &Mongo{
		config: config,
	}

	ctx := context.Background()
	var err error

	if m.client, err = mongo.Connect(ctx, options.Client().ApplyURI(config.Mongo.Uri)); err != nil {
		panic(err)
	} else if err = m.client.Ping(ctx, nil); err != nil {
		panic(err)
	} else {
		m.db = m.client.Database(config.Mongo.Db)

		m.user = m.db.Collection("user")
		m.content = m.db.Collection("content")
		m.history = m.db.Collection("history")

		//if err := createIndex(m.db.Collection("test"), []string{"key"}, []string{}); err != nil {
		//	panic(err)
		//}

		if err = createIndex(m.user, []string{"user"}, []string{"user"}); err != nil {
			panic(err)
		} else if err = createIndex(m.content, []string{"name"}, []string{"name"}); err != nil {
			panic(err)
		} else if err = createIndex(m.history, []string{"user"}, []string{}); err != nil {
			panic(err)
		}

	}

	return m, nil
}

// Single Field Index
// 1. 원하는 Field 선택 (indexes []string)
// 2. 단일 필드 인덱스에서 1은 오름차순 -1은 내림차순을 의미
// 하지만 단일 필드 인덱스에 오름차순, 내림차순은 크게 중요하지 않다. 어떤 방향으로 가도 동일하게 접근하기 때문이다.
// 해당 필드에 대해 인덱스가 없는 경우에만 생성한다.
func createIndex(collection *mongo.Collection, indexes, uniques []string) error {

	type indexOptions struct {
		key    string
		order  int64
		unique bool
	}

	var indexesOpt []indexOptions

	for _, field := range indexes {
		noU := false

		for _, unique := range uniques {
			if field == unique {
				indexesOpt = append(indexesOpt, indexOptions{key: field, order: -1, unique: true})
				noU = true
				break
			}
		}

		if !noU {
			indexesOpt = append(indexesOpt, indexOptions{key: field, order: -1, unique: false})
		}
	}

	ctx := context.Background()

	doNotCreate := make(map[string]indexOptions)

	if indexCursor, err := collection.Indexes().List(ctx); err != nil {
		return err
	} else {
		defer indexCursor.Close(ctx)

		for indexCursor.Next(ctx) {
			if v, ok := indexCursor.Current.Lookup("name").StringValueOK(); !ok || v == "_id_" {
				continue
			} else {
				splits := strings.Split(v, "_")

				if len(splits) == 2 {
					if order, err := strconv.Atoi(splits[1]); err != nil {
						return err
					} else {
						if order == 1 || order == -1 {
							doNotCreate[splits[0]] = indexOptions{key: splits[0], order: int64(order), unique: false}
						}
					}
				}
			}
		}
	}

	for _, newIndex := range indexesOpt {
		if _, ok := doNotCreate[newIndex.key]; !ok {
			opt := options.Index()

			if newIndex.unique {
				opt.SetUnique(true)
			}

			m := mongo.IndexModel{
				Keys:    bson.D{{Key: newIndex.key, Value: 1}},
				Options: opt,
			}

			if _, err := collection.Indexes().CreateOne(ctx, m); err != nil {
				return err
			}
		}
	}

	return nil
}
