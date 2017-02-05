package model

import (
	"fmt"
	"github.com/satori/go.uuid"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Plugin struct {
	Author      User
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description"`
	Url         string `json:"url"`
	Uuid        uuid.UUID
	Slug        string

	id    bson.ObjectId
	exact bool
}

type PluginDAO interface {
	FindBySlug(name string) (Plugin, error)
	FindBySlugAndVersion(name, version string) (Plugin, error)
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
	Uuid        string
	Slug        string
}

func (mp mongoPlugin) buildPlugin() Plugin {
	return Plugin{
		Author:      userFromMongoId(mp.UserId),
		Name:        mp.Name,
		Version:     mp.Version,
		Description: mp.Description,
		Url:         mp.Url,
		Uuid:        uuid.FromStringOrNil(mp.Uuid),
		Slug:        mp.Slug,
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
		Uuid:        p.Uuid.String(),
		Slug:        p.Slug,
	}
	if p.exact {
		mp.Id = p.id
	}
	return mp, nil
}

// FindBySlug returns the latest version of the plugin that has specified slug.
func (dao *MongoPluginDAO) FindBySlug(slug string) (Plugin, error) {
	mp := mongoPlugin{}
	err := dao.Collection.Find(bson.M{
		"slug": slug,
	}).Sort("-version").One(&mp)
	if err != nil {
		return Plugin{}, err
	}
	return mp.buildPlugin(), nil
}

// FindBySlugAndVersion returns a plugin that has specified slug and version.
func (dao *MongoPluginDAO) FindBySlugAndVersion(slug, version string) (Plugin, error) {
	mp := mongoPlugin{}
	err := dao.Collection.Find(bson.M{
		"slug":    slug,
		"version": version,
	}).One(&mp)
	if err != nil {
		return Plugin{}, err
	}
	return mp.buildPlugin(), nil
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
