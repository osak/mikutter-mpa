package model

import (
	"crypto/rand"
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
	Create(user User) (Session, error)
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

const sessionIdRunes = "abcdefghijlkmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func (dao *mysqlSessionDAO) Create(user User) (Session, error) {
	buf := make([]byte, 32)
	_, err := rand.Read(buf)
	if err != nil {
		return Session{}, err
	}
	sessionId := ""
	for _, b := range buf {
		pos := int(b) % len(sessionIdRunes)
		sessionId += sessionIdRunes[pos : pos+1]
	}

	_, err = dao.db.NamedExec(`INSERT INTO sessions (id, user_id) VALUES (:id, :userId)`, map[string]interface{}{
		"id":     sessionId,
		"userId": user.Id,
	})
	if err != nil {
		return Session{}, err
	}
	return Session{
		Id:   sessionId,
		User: user,
	}, nil
}
