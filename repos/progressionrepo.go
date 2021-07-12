package repo

import (
	"context"
	"fmt"

	model "github.com/jasonzguo/vaccination-progress-service/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type progressionRepo struct {
	collection *mongo.Collection
}

var instance *progressionRepo = nil

func GetProgressionRepo() *progressionRepo {
	if instance == nil {
		instance = new(progressionRepo)
	}
	return instance
}

func (pr *progressionRepo) SetCollection(c *mongo.Collection) {
	pr.collection = c
}

func (pr *progressionRepo) Find(ctx context.Context, lastId string) ([]model.ProgressionModel, error) {
	var progressions []model.ProgressionModel

	filters := bson.M{}
	if lastId != "" {
		lastObjectId, err := primitive.ObjectIDFromHex(lastId)
		if err != nil {
			return nil, fmt.Errorf("[GetProgressionRepo.Find] error in calling primitive.ObjectIDFromHex %v", err)
		}
		filters = bson.M{"_id": bson.M{"$gt": lastObjectId}}
	}
	cursor, err := pr.collection.Find(ctx, filters, options.Find().SetLimit(10))

	if err != nil {
		return nil, fmt.Errorf("[GetProgressionRepo.Find] error in calling collection.Find %v", err)
	}

	if err = cursor.All(ctx, &progressions); err != nil {
		return nil, fmt.Errorf("[GetProgressionRepo.Find] error in calling cursor.All %v", err)
	}

	return progressions, nil
}
