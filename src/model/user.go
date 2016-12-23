package model

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
)

type User struct {
	Id    int    `json:"id"`
	Login string `json:"login"`
	Name  string `json:"name"`
}

type UserDAO interface {
	FindById(id int) (User, error)
	FindByLogin(name string) (User, error)
	Create(user User) (User, error)
}

type mysqlUserDAO struct {
	db *sqlx.DB
}

func NewUserMySQLDAO(db *sqlx.DB) UserDAO {
	dao := new(mysqlUserDAO)
	dao.db = db
	return dao
}

func (dao *mysqlUserDAO) FindById(id int) (User, error) {
	user := User{}
	err := dao.db.Get(&user, `SELECT * FROM users WHERE id=?`, id)
	return user, err
}

func (dao *mysqlUserDAO) FindByLogin(name string) (User, error) {
	user := User{}
	err := dao.db.Get(&user, `SELECT * FROM users WHERE name=?`, name)
	if err == sql.ErrNoRows {
		return User{}, ErrNoEntry
	}
	return user, err
}

func (dao *mysqlUserDAO) Create(user User) (User, error) {
	result, err := dao.db.NamedExec(`INSERT INTO users (login, name) VALUES (:login, :name)`, user)
	if err != nil {
		return User{}, err
	}
	id, err := result.LastInsertId()
	user.Id = int(id)
	return user, err
}
