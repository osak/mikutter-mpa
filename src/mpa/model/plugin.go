package model

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Plugin struct {
	Author      User
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description"`
	Url         string `json:"url"`

	id    bson.ObjectId
	exact bool
}

type PluginDAO interface {
	FindByName(name string) (Plugin, error)
	FindByKeyword(keyword string) ([]Plugin, error)
	Create(plugin *Plugin) error
}

type MongoPluginDAO struct {
	Collection *mgo.Collection
}

type mongoPlugin struct {
	Id          bson.ObjectId `bson:"_id,omitempty"`
	UserId      bson.ObjectId `bson:"user_id"`
	Name        string
	Version     string
	Description string
	Url         string
}

func (mp mongoPlugin) buildPlugin() Plugin {
	return Plugin{
		Author:      userFromMongoId(mp.UserId),
		Name:        mp.Name,
		Version:     mp.Version,
		Description: mp.Description,
		Url:         mp.Url,
	}
}

func (p Plugin) buildMongoPlugin() (mongoPlugin, error) {
	if !p.Author.Exact() {
		return mongoPlugin{}, fmt.Errorf("mpa/model: Plugin.Author author must be exact")
	}
	mp := mongoPlugin{
		UserId:      p.Author.id,
		Name:        p.Name,
		Version:     p.Version,
		Description: p.Description,
		Url:         p.Url,
	}
	if p.exact {
		mp.Id = p.id
	}
	return mp, nil
}

func (dao *MongoPluginDAO) FindByName(name string) (Plugin, error) {
	mp := mongoPlugin{}
	err := dao.Collection.Find(bson.M{
		"name": name,
	}).One(&mp)
	if err != nil {
		return Plugin{}, err
	}
	return mp.buildPlugin(), nil
}

func (dao *MongoPluginDAO) FindByKeyword(keyword string) ([]Plugin, error) {
	mps := []mongoPlugin{}
	err := dao.Collection.Find(bson.M{
		"name": bson.M{
			"$regex": ".*" + keyword + ".*",
		},
	}).All(&mps)
	if err != nil {
		return nil, err
	}
	plugins := make([]Plugin, len(mps))
	for i, mp := range mps {
		plugins[i] = mp.buildPlugin()
	}
	return plugins, nil
}

func (dao *MongoPluginDAO) Create(p *Plugin) error {
	mp, err := p.buildMongoPlugin()
	if err != nil {
		return err
	}
	info, err := dao.Collection.Upsert(bson.M{"name": mp.Name}, mp)
	if err != nil {
		return err
	}
	if objId, ok := info.UpsertedId.(bson.ObjectId); ok {
		p.id = objId
	} else {
		return fmt.Errorf("mpa/model: Upserted Id is expected to be ObjectId but got %T", info.UpsertedId)
	}
	p.exact = true
	return nil
}
