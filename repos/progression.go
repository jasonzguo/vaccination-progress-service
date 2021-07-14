package repo

import (
	"context"
	"fmt"

	model "github.com/jasonzguo/vaccination-progress-service/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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

func (pr *progressionRepo) FindAll(ctx context.Context, lastId string) ([]model.ProgressionModel, error) {
	var progressions model.PaginatedProgressionModel

	facetStage := bson.M{
		"$facet": bson.M{
			"data": []bson.M{
				{"$match": bson.M{}},
				{"$limit": 10},
			},
			"count": []bson.M{{"$count": "count"}},
		},
	}

	addFieldsStage := bson.M{
		"$addFields": bson.M{
			"meta.count":  bson.M{"$arrayElemAt": []interface{}{"$count.count", 0}},
			"meta.lastId": bson.M{"$arrayElemAt": []interface{}{"$data._id", -1}},
		},
	}

	projectStage := bson.M{
		"$project": bson.M{
			"data": 1,
			"meta": 1,
		},
	}

	cursor, err := pr.collection.Aggregate(ctx, []bson.M{facetStage, addFieldsStage, projectStage})

	if err != nil {
		return nil, fmt.Errorf("[GetProgressionRepo.Find] error in calling collection.Find %v", err)
	}

	if cursor.TryNext(ctx) {
		cursor.Next(ctx)
		err := cursor.Decode(&progressions)
		if err != nil {
			return nil, fmt.Errorf("[GetProgressionRepo.Find] error in calling cursor.Decode %v", err)
		}
	}

	return progressions.Data, nil
}
