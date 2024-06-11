package mongo

import (
	"errors"
	"fmt"
)

func (m *MService) PostCreateUser(user string) error {

	if err := m.repository.Mongo.CreateUser(user); err != nil {
		fmt.Println(err)
		return err
	} else {
		return err
	}
}

func (m *MService) PostCreateContent(name string, price int64) error {
	if err := m.repository.Mongo.CreateContent(name, price); err != nil {
		fmt.Println(err)
		return err
	} else {
		return err
	}
}

func (m *MService) PostBuy(user string) error {
	/*
		기존 user 장바구니 데이터를 History 에 넣어준다.
	*/
	if u, err := m.repository.Mongo.GetUserBucket(user); err != nil {
		fmt.Println("GetUserBucket err: ", err)
		return err
	} else if len(u.Bucket) == 0 {
		return errors.New("장바구니에 데이터가 없습니다")
	} else {
		// history 컬렉션에 데이터 넣어주기
		if err = m.repository.Mongo.UpsertHistory(user, u.Bucket); err != nil {
			fmt.Println("UpsertHistory err: ", err)
			return err
		}
		// 장바구니 비우기
		if err = m.repository.Mongo.RemoveUserBucket(user); err != nil {
			fmt.Println("RemoveUserBucket err: ", err)
			return err
		}
		return nil
	}
}

func (m *MService) PostCreateBucket(user, content string) error {
	// 1. User 가 존재 하는지
	// 2. Content 가 존재 하는지
	if c, err := m.repository.Mongo.GetContent(content); err != nil {
		fmt.Println("GetContent err: ", err)
		return err
	} else if _, err = m.repository.Mongo.GetUserBucket(user); err != nil {
		fmt.Println("GetUserBucket err: ", err)
		return err
	} else if len(c) == 0 {
		return errors.New("content 없음")
	} else if err = m.repository.Mongo.CreateBucket(user, content); err != nil {
		fmt.Println("CreateBucket err: ", err)
		return err
	} else {
		return nil
	}
}
