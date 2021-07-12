package service

import (
	"context"

	model "github.com/jasonzguo/vaccination-progress-service/models"
	repo "github.com/jasonzguo/vaccination-progress-service/repos"
)

type progressionService struct{}

var progressionServiceInstance *progressionService = nil

func GetProgressionService() *progressionService {
	if progressionServiceInstance == nil {
		progressionServiceInstance = new(progressionService)
	}
	return progressionServiceInstance
}

func (ps *progressionService) FindAll(ctx context.Context, lastId string) ([]model.ProgressionModel, error) {
	documents, err := repo.GetProgressionRepo().Find(ctx, lastId)
	return documents, err
}
