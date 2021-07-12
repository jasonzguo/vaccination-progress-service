package controller

import (
	"context"

	model "github.com/jasonzguo/vaccination-progress-service/models"
	service "github.com/jasonzguo/vaccination-progress-service/services"
)

type progressionController struct{}

var prograssionControllerInstance *progressionController = nil

func GetProgressionController() *progressionController {
	if prograssionControllerInstance == nil {
		prograssionControllerInstance = new(progressionController)
	}
	return prograssionControllerInstance
}

func (vc *progressionController) GetAll(ctx context.Context, lastId string) ([]model.ProgressionModel, error) {
	documents, err := service.GetProgressionService().FindAll(ctx, lastId)
	return documents, err
}
