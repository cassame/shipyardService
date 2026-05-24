package part

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *repo) UpdateStock(ctx context.Context, uuid string, delta int64) error {
	filter := bson.M{"_id": uuid}

	update := bson.M{
		"$inc": bson.M{"stock_quantity": delta},
	}

	_, err := r.col.UpdateOne(ctx, filter, update)
	return err
}
