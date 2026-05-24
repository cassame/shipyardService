package part

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type repo struct {
	col *mongo.Collection
}

