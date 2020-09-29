package repo

import (
	"context"
	"graph/ent"
	"graph/ent/user"
)

type User struct {
	Db *ent.Client
}

func NewUserRepo(db *ent.Client) *User {
	return &User{
		Db: db,
	}
}

func (u *User) Create(name string, age int) (err error) {
	_, err = u.Db.User.
		Create().
		SetAge(age).
		SetName(name).
		Save(context.Background())
	return err
}

func (u *User) GetByName(name string) (res *ent.User, err error) {
	res, err = u.Db.User.Query().Where(user.Name(name)).First(context.Background())
	if err != nil {
		return nil, err
	}
	return res, err
}

func (u *User) GetAll() (res []*ent.User, err error) {
	res, err = u.Db.User.Query().All(context.Background())
	if err != nil {
		return nil, err
	}
	return res, err
}