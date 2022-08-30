package controllers

import (
	"encoding/json"
	"errors"
	"github.com/crawlerv/mongo-inmemdb-test-task/internal/dto"
	"github.com/crawlerv/mongo-inmemdb-test-task/internal/resources"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type Service interface {
	Create(dto dto.CreateInput) error
	Update(id string, dto dto.UpdateInput) error
	Delete(id string) error
	List() ([]dto.Output, error)
}

type Controller struct {
	app resources.App
	s   Service
}

func NewController(app resources.App, s Service) Controller {
	return Controller{app, s}
}

func (c Controller) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var cDto dto.CreateInput
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&cDto)
	if err != nil {
		c.app.Log().Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = c.s.Create(cDto)
	if err != nil {
		c.app.Log().Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(200)
}

func (c Controller) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var cDto dto.UpdateInput
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&cDto)
	if err != nil {
		c.app.Log().Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.s.Update(p.ByName("id"), cDto)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		c.app.Log().Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(200)
}

func (c Controller) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	err := c.s.Delete(p.ByName("id"))
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		c.app.Log().Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(200)
}

func (c Controller) List(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	m, err := c.s.List()
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		c.app.Log().Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	jData, err := json.Marshal(m)
	if err != nil {
		c.app.Log().Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	_, err = w.Write(jData)
	if err != nil {
		c.app.Log().Error(err)
	}
}
