package mongo

import "go.mongodb.org/mongo-driver/bson/primitive"

type Model struct {
	ID          primitive.ObjectID `bson:"_id" faker:"-"`
	FirstName   string             `bson:"first_name" faker:"first_name"`
	LastName    string             `bson:"last_name" faker:"last_name"`
	Age         uint8              `bson:"age" faker:"-"`
	CardNumber  uint32             `bson:"card_number"`
	PhoneNumber string             `bson:"phone_number" faker:"phone_number"`
	Verified    bool               `bson:"verified" faker:"-"`
}
