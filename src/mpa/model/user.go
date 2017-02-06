package model

import (
	"errors"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	ErrNoEntry = errors.New("mpa/model: No user entry exists")
)

type User struct {
	Login      string `bson:"login" json:"login"`
	Name       string `bson:"name" json:"name"`
	LoginToken string

	id    bson.ObjectId
	exact bool
}

// Exact returns true iff the object is based on a DB entry.
// An object is considered based on a DB entry when any of following conditions
// are satisfied:
// - The object is read from DB
// - The object has been written out into DB
func (u *User) Exact() bool {
	return u.exact
}

// SameUser returns true iff u1 and u2 refers the same user.
func SameUser(u1, u2 User) bool {
	return u1.id == u2.id
}

type UserDAO interface {
	FindByLogin(login string) (User, error)
	Create(user *User) (*User, error)
	Fill(user *User) error
}

type MongoUserDAO struct {
	Collection *mgo.Collection
}

type mongoUser struct {
	Id         bson.ObjectId `bson:"_id"`
	Login      string
	Name       string
	LoginToken string
}

func (mu mongoUser) buildUser() User {
	return User{
		Login:      mu.Login,
		Name:       mu.Name,
		LoginToken: mu.LoginToken,
		id:         mu.Id,
		exact:      true,
	}
}

func (u User) buildMongoUser() mongoUser {
	mu := mongoUser{
		Login:      u.Login,
		Name:       u.Name,
		LoginToken: u.LoginToken,
	}
	if u.Exact() {
		mu.Id = u.id
	}
	return mu
}

func userFromMongoId(id bson.ObjectId) User {
	return User{
		id: id,
	}
}

func (dao *MongoUserDAO) FindByLogin(login string) (User, error) {
	mu := mongoUser{}
	err := dao.Collection.Find(bson.M{
		"login": login,
	}).One(&mu)
	switch {
	case err == mgo.ErrNotFound:
		return User{}, ErrNoEntry
	case err != nil:
		return User{}, err
	}
	return mu.buildUser(), nil
}

func (dao *MongoUserDAO) Create(user *User) (*User, error) {
	if _, err := dao.FindByLogin(user.Login); err != ErrNoEntry {
		if err == nil {
			return user, fmt.Errorf("mpa/user: Duplicate user %v", user.Login)
		} else {
			return user, err
		}
	}

	mu := user.buildMongoUser()
	info, err := dao.Collection.Upsert(bson.M{"login": user.Login}, &mu)
	if err != nil {
		return user, err
	}
	if bsonId, ok := info.UpsertedId.(bson.ObjectId); ok {
		user.id = bsonId
		user.exact = true
		return user, nil
	} else {
		return nil, fmt.Errorf("mpa/model: Unexpected id type: %T", info.UpsertedId)
	}
}

func (dao *MongoUserDAO) Fill(user *User) error {
	if user.Exact() {
		return nil
	}
	mu := mongoUser{}
	err := dao.Collection.Find(bson.M{
		"_id": user.id,
	}).One(&mu)
	if err != nil {
		return err
	}
	*user = mu.buildUser()
	return nil
}
