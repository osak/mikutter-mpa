package model

import (
	"github.com/jmoiron/sqlx"
)

type Plugin struct {
	Name        string
	Version     string
	Description string
	Url         string
}

type PluginDAO interface {
	FindPlugin(name string) (Plugin, error)
}

type mysqlPluginDAO struct {
	db *sqlx.DB
}

func NewMySQLDAO(db *sqlx.DB) PluginDAO {
	dao := new(mysqlPluginDAO)
	dao.db = db
	return dao
}

func (dao *mysqlPluginDAO) FindPlugin(name string) (Plugin, error) {
	plugin := Plugin{}
	err := dao.db.Get(&plugin, `SELECT * FROM plugins WHERE name=:name`, name)
	if err != nil {
		return nil, err
	}
	return plugin, nil
}
