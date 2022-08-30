package service

import (
	"context"
	"github.com/crawlerv/mongo-inmemdb-test-task/internal/db/mongo"
	"github.com/crawlerv/mongo-inmemdb-test-task/internal/resources"
	"time"
)

type Ticker struct {
	app resources.App
	r   DbRepository
	s   CacheService
}
type (
	DbRepository interface {
		GetAllData() ([]*mongo.Model, error)
	}
	CacheService interface {
		Set(key, value interface{}) error
	}
)

func NewTicker(app resources.App, r DbRepository, s CacheService) Ticker {
	return Ticker{app, r, s}
}

func (t Ticker) LoadData(ctx context.Context) {
	t.app.Log().Infof("Start ticker with %d sec interval", t.app.Config().TickerIntervalSeconds)
	tt := time.NewTicker(time.Second * time.Duration(t.app.Config().TickerIntervalSeconds))
	defer tt.Stop()
	for {
		select {
		case <-ctx.Done():
			t.app.Log().Info("context done")
			return
		case <-tt.C:
			t.loadAll()
		}
	}
}

func (t Ticker) loadAll() {
	d, err := t.r.GetAllData()
	if err != nil {
		t.app.Log().Errorf("error at getting all data from db. error - %v", err)
		return
	}
	for _, dd := range d {
		err = t.s.Set(dd.ID, dd)
		if err != nil {
			t.app.Log().Errorf("error at adding data to cache. error - %v", err)
		}
	}
}
