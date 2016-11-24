package model

import (
	"github.com/jmoiron/sqlx"
)

type Plugin struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description"`
	Url         string `json:"url"`
}

type PluginDAO interface {
	FindPlugin(name string) (Plugin, error)
}

type mysqlPluginDAO struct {
	db *sqlx.DB
}

func NewPluginMySQLDAO(db *sqlx.DB) PluginDAO {
	dao := new(mysqlPluginDAO)
	dao.db = db
	return dao
}

func (dao *mysqlPluginDAO) FindPlugin(name string) (Plugin, error) {
	plugin := Plugin{}
	err := dao.db.Get(&plugin, `SELECT * FROM plugins WHERE name=?`, name)
	if err != nil {
		return Plugin{}, err
	}
	return plugin, nil
}
