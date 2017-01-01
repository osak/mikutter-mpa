package plugin

import (
	"github.com/jmoiron/sqlx"
)

type Plugin struct {
	Id          int    `json:"id,omitempty" db:"id"`
	UserId      int    `json:"userId,omitempty" db:"user_id"`
	Name        string `json:"name" db:"name"`
	Version     string `json:"version" db:"version"`
	Description string `json:"description" db:"description"`
	Url         string `json:"url" db:"url"`
}

type PluginDAO interface {
	FindPlugin(name string) (Plugin, error)
	FindPlugins(keyword string) ([]Plugin, error)
	Create(plugin *Plugin) error
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

func (dao *mysqlPluginDAO) FindPlugins(keyword string) ([]Plugin, error) {
	plugins := []Plugin{}
	pattern := "%" + keyword + "%"
	stmt, err := dao.db.PrepareNamed(`SELECT * FROM plugins WHERE name LIKE :pattern OR description LIKE :pattern`)
	if err != nil {
		return plugins, err
	}

	err = stmt.Select(&plugins, map[string]interface{}{
		"pattern": pattern,
	})
	if err != nil {
		return plugins, err
	}
	return plugins, nil
}

func (dao *mysqlPluginDAO) Create(plugin *Plugin) error {
	_, err := dao.db.NamedExec(`INSERT INTO plugins SET
	user_id=:user_id,
	name=:name,
	version=:version,
	description=:description,
	url=:url`, plugin)
	if err != nil {
		return err
	}
	return nil
}
