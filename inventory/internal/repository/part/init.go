package part

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"

	"inventory/internal/model"
)

func NewRepository(ctx context.Context, client *mongo.Client, dbName string) *repo {
	r := &repo{
		col: client.Database(dbName).Collection("parts"),
	}

	r.setupMockData(ctx)

	return r
}

func (r *repo) setupMockData(ctx context.Context) {
	part := &model.Part{
		UUID:          "11111111-1111-1111-1111-111111111111",
		Name:          "Двигатель Гипердрайва",
		Description:   "Ускоряет до гипера",
		Price:         999.99,
		StockQuantity: 5,
		Category:      1,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	filter := bson.M{"_id": part.UUID}

	opts := options.Replace().SetUpsert(true)

	_, err := r.col.ReplaceOne(ctx, filter, part, opts)
	if err != nil {
		fmt.Println("failed to setup mock data:", err)
	}
}
