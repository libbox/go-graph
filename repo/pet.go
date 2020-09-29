package repo

import (
	"context"
	"graph/ent"
)

type Pet struct {
	Db *ent.Client
}

func NewPetRepo(client *ent.Client) *Pet {
	return &Pet{
		Db: client,
	}
}

func (p *Pet) Create(name string) (pet *ent.Pet, err error){
	pet, err = p.Db.Pet.Create().
		SetName(name).
		SetAge(99).
		Save(context.Background())
	return pet, err
}

func (p *Pet) GetAll() ([]*ent.Pet, error){
	res, err := p.Db.Debug().Pet.Query().All(context.Background())
	if err != nil {
		return nil, err
	}
	return res, err
}