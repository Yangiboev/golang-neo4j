package service

import (
	"context"

	"github.com/Yangiboev/golang-neo4j/api/models"
	"github.com/Yangiboev/golang-neo4j/pkg/helper"
	"github.com/Yangiboev/golang-neo4j/pkg/logger"
	"github.com/Yangiboev/golang-neo4j/storage"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type actionService struct {
	logger  logger.Logger
	storage storage.StorageI
}

func NewActionService(driver neo4j.Driver, log logger.Logger) *actionService {
	return &actionService{
		storage: storage.NewResponsibleStorage(driver),
		logger:  log,
	}
}

func (as *actionService) Create(ctx context.Context, req *models.Action) (*models.CreateResponse, error) {
	res, err := as.storage.Action().Create(req)

	if err != nil {
		as.logger.Error("error while creating Action", logger.Error(err))
		return nil, helper.HandleError(as.logger, err, "error while creating Action", req)
	}
	return res, nil
}

func (as *actionService) Get(ctx context.Context, req *models.GetRequest) (*models.Action, error) {

	response, err := as.storage.Action().Get(req.ID)

	if err != nil {
		return nil, helper.HandleError(as.logger, err, "error while getting Action", req)
	}
	return response, nil
}
func (as *actionService) GetAll(ctx context.Context, req *models.GetAllActionsRequest) (*models.GetAllActionsResponse, error) {

	response, _, err := as.storage.Action().GetAll(req.Page, req.Limit)
	if err != nil {
		return nil, helper.HandleError(as.logger, err, "error while getting all Actions", req)
	}
	return &models.GetAllActionsResponse{
		Count:   int64(len(response)),
		Actions: response,
	}, nil
}

func (as *actionService) Update(ctx context.Context, req *models.Action) error {
	err := as.storage.Action().Update(req)
	if err != nil {
		return helper.HandleError(as.logger, err, "error while updating Action", nil)
	}

	return nil
}

func (as *actionService) Delete(ctx context.Context, req *models.DeleteRequest) error {
	err := as.storage.Action().Delete(req.ID)

	if err != nil {
		return helper.HandleError(as.logger, err, "error while deleting Action", nil)
	}

	return nil
}
