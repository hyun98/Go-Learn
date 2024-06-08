package mongo

import (
	"eCommerce/types"
)

func (m *MService) GetUserBucket(user string) (*types.User, error) {
	if r, err := m.repository.Mongo.GetUserBucket(user); err != nil {
		return nil, err
	} else {
		return r, err
	}
}

func (m *MService) GetContent(name string) ([]*types.Content, error) {
	if r, err := m.repository.Mongo.GetContent(name); err != nil {
		return nil, err
	} else {
		return r, err
	}
}

func (m *MService) GetUserBucketHistory(user string) (*types.History, error) {
	if r, err := m.repository.Mongo.GetUserBucketHistory(user); err != nil {
		return nil, err
	} else {
		return r, err
	}
}
