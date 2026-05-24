package part

import (
	"context"
	"inventory/internal/model"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *repo) ListParts(ctx context.Context, filterModel *model.PartsFilter) ([]*model.Part, error) {

	var parts []*model.Part

	filter := bson.M{}
	if filterModel != nil {
		if len(filterModel.UUIDs) > 0 {
			filter["_id"] = bson.M{"$in": filterModel.UUIDs}
		}
		if len(filterModel.Names) > 0 {
			filter["name"] = bson.M{"$in": filterModel.Names}
		}
		if len(filterModel.Categories) > 0 {
			filter["category"] = bson.M{"$in": filterModel.Categories}
		}
		if len(filterModel.Tags) > 0 {
			filter["tags"] = bson.M{"$in": filterModel.Tags}
		}
		if len(filterModel.ManufacturerCountries) > 0 {
			filter["manufacturer.country"] = bson.M{"$in": filterModel.ManufacturerCountries}
		}
	}

	cursor, err := r.col.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = cursor.Close(ctx)
	}()

	if err := cursor.All(ctx, &parts); err != nil {
		return nil, err
	}

	return parts, nil
}
