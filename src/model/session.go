package model

import (
	"github.com/jmoiron/sqlx"
)

type sessionRow struct {
	id     string
	userId int
}

type Session struct {
	Id   string
	User User
}

type SessionDAO interface {
	FindBySessionId(id string) (Session, error)
}

type mysqlSessionDAO struct {
	db      *sqlx.DB
	userDAO UserDAO
}

func NewSessionMySQLDAO(db *sqlx.DB, userDAO UserDAO) SessionDAO {
	return &mysqlSessionDAO{
		db:      db,
		userDAO: userDAO,
	}
}

func (dao *mysqlSessionDAO) FindBySessionId(id string) (Session, error) {
	sessionRow := sessionRow{}
	err := dao.db.Get(&sessionRow, `SELECT * FROM sessions WHERE id=?`, id)
	if err != nil {
		return Session{}, err
	}

	user, err := dao.userDAO.FindById(sessionRow.userId)
	if err != nil {
		return Session{}, err
	}

	return Session{
		Id:   sessionRow.id,
		User: user,
	}, nil
}
