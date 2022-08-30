package resources

import (
	"context"
	"fmt"
	"github.com/crawlerv/mongo-inmemdb-test-task/config"
	"github.com/crawlerv/mongo-inmemdb-test-task/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var _ App = (*TestTask)(nil)

type App interface {
	Log() *logger.Logger
	Config() *config.Config
	Db() *mongo.Database
	Release(ctx context.Context) error
}

type TestTask struct {
	log    *logger.Logger
	config *config.Config
	db     *mongo.Database
}

func (t TestTask) Log() *logger.Logger {
	return t.log
}

func (t TestTask) Config() *config.Config {
	return t.config
}

func (t TestTask) Db() *mongo.Database {
	return t.db
}

func New(ctx context.Context) *TestTask {
	conf := GetConfig()
	log := logger.Init(conf.Log.Dir, conf.Log.File, conf.Log.Encoding)

	ctx, _ = context.WithTimeout(ctx, 5*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%d", conf.Mongo.User, conf.Mongo.Password, conf.Mongo.Host, conf.Mongo.Port)))
	if err != nil {
		log.Fatalf("error at open mongodb connection. Err - %v", err)
	}
	db := client.Database(conf.Mongo.Name)
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	return &TestTask{
		log:    log,
		config: conf,
		db:     db,
	}
}

func (t TestTask) Release(ctx context.Context) error {
	t.Log().Info("Releasing all resources")
	if err := t.db.Client().Disconnect(ctx); err != nil {
		t.Log().Error(err)
	}
	return nil
}
