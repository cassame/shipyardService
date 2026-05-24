package part

import (
	"context"
	"inventory/internal/model"
	"log"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (r *repo) GetPart(ctx context.Context, uuid string) (*model.Part, error) {
	log.Printf("Запрос детали с UUID: %s", uuid)

	part := &model.Part{}

	filter := bson.M{"uuid": uuid}

	err := r.col.FindOne(ctx, filter).Decode(part)
	if err == mongo.ErrNoDocuments {
		log.Printf("Деталь с UUID %s не найдена", uuid)
		return nil, nil
	}

	return part, err
}
