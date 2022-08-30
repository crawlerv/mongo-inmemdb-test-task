package service

import (
	"github.com/crawlerv/mongo-inmemdb-test-task/internal/db/mongo"
	"github.com/crawlerv/mongo-inmemdb-test-task/internal/dto"
	"github.com/crawlerv/mongo-inmemdb-test-task/internal/resources"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Repository interface {
	GetAllData() ([]*mongo.Model, error)
	Insert(m mongo.Model) error
	Update(id string, m mongo.Model) error
	FindById(id string) (*mongo.Model, error)
	Delete(id string) error
}

type Api struct {
	app resources.App
	r   Repository
}

func New(app resources.App, r Repository) Api {
	return Api{app, r}
}

func (a Api) Create(dto dto.CreateInput) error {
	m := mongo.Model{
		ID:          primitive.NewObjectID(),
		FirstName:   dto.FirstName,
		LastName:    dto.LastName,
		Age:         dto.Age,
		CardNumber:  dto.CardNumber,
		PhoneNumber: dto.PhoneNumber,
		Verified:    dto.Verified,
	}
	err := a.r.Insert(m)
	if err != nil {
		return err
	}
	return nil
}

func (a Api) Update(id string, dto dto.UpdateInput) error {
	m, err := a.r.FindById(id)
	if err != nil {
		return err
	}
	if dto.FirstName != nil {
		m.FirstName = *dto.FirstName
	}
	if dto.LastName != nil {
		m.LastName = *dto.LastName
	}
	if dto.Age != nil {
		m.Age = *dto.Age
	}
	if dto.CardNumber != nil {
		m.CardNumber = *dto.CardNumber
	}
	if dto.PhoneNumber != nil {
		m.PhoneNumber = *dto.PhoneNumber
	}
	if dto.Verified != nil {
		m.Verified = *dto.Verified
	}
	return a.r.Update(id, *m)
}

func (a Api) Delete(id string) error {
	return a.r.Delete(id)
}

func (a Api) List() ([]dto.Output, error) {
	m, err := a.r.GetAllData()
	if err != nil {
		return nil, err
	}
	res := make([]dto.Output, 0, len(m))
	for _, mm := range m {
		res = append(res, dto.Output{
			FirstName:   mm.FirstName,
			LastName:    mm.LastName,
			Age:         mm.Age,
			CardNumber:  mm.CardNumber,
			PhoneNumber: mm.PhoneNumber,
			Verified:    mm.Verified,
		})
	}
	return res, nil
}
