package part

import (
	"context"
	"inventory/internal/model"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *repo) ListParts(ctx context.Context) ([]*model.Part, error) {

	var parts []*model.Part

	cursor, err := r.col.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &parts); err != nil {
		return nil, err
	}

	return parts, nil
}
