package mongo

import (
	"context"
	"github.com/crawlerv/mongo-inmemdb-test-task/internal/resources"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repo struct {
	ctx context.Context
	app resources.App
}

func NewRepo(ctx context.Context, app resources.App) Repo {
	return Repo{ctx, app}
}

func (r Repo) Insert(model Model) error {
	_, err := r.app.Db().Collection("users").InsertOne(r.ctx, model)
	return err
}

func (r Repo) InsertMany(models []interface{}) error {
	_, err := r.app.Db().Collection("users").InsertMany(r.ctx, models)
	return err
}

func (r Repo) Update(id string, model Model) error {
	filter := bson.M{"_id": getObjectIdFromString(id)}
	_, err := r.app.Db().Collection("users").ReplaceOne(r.ctx, filter, model)
	return err
}

func (r Repo) Delete(id string) error {
	filter := bson.M{"_id": getObjectIdFromString(id)}
	_, err := r.app.Db().Collection("users").DeleteOne(r.ctx, filter)
	return err
}

func (r Repo) GetAllData() ([]*Model, error) {
	return r.getInternal(bson.D{{}}, options.Find().SetSort(bson.D{{"_id", -1}}))
}

func (r Repo) FindById(id string) (*Model, error) {
	m := &Model{}
	filter := bson.M{"_id": getObjectIdFromString(id)}
	err := r.app.Db().Collection("users").FindOne(r.ctx, filter).Decode(m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (r Repo) getInternal(filter interface{}, opts *options.FindOptions) ([]*Model, error) {
	cursor, err := r.app.Db().Collection("users").Find(r.ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	var results []*Model
	if err = cursor.All(r.ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}

func getObjectIdFromString(id string) primitive.ObjectID {
	oId, _ := primitive.ObjectIDFromHex(id)
	return oId
}
