package service

import (
	"context"

	"github.com/Yangiboev/golang-neo4j/api/models"
	"github.com/Yangiboev/golang-neo4j/pkg/helper"
	"github.com/Yangiboev/golang-neo4j/pkg/logger"
	"github.com/Yangiboev/golang-neo4j/storage"
)

type responsibleService struct {
	logger  logger.Logger
	storage storage.StorageI
}

func NewResponsibleService(db *db.Database, log logger.Logger) *responsibleService {
	return &responsibleService{
		storage: storage.NewResponsibleStorage(db),
		logger:  log,
	}
}

func (ps *responsibleService) Create(ctx context.Context, req *models.Responsible) (*models.CreateResponse, error) {
	ID, err := ps.storage.Responsible().Create(req)

	if err != nil {
		ps.logger.Error("error while creating responsible", logger.Error(err))
		return nil, helper.HandleError(ps.logger, err, "error while creating responsible", req)
	}
	return &models.CreateResponse{
		ID: ID,
	}, nil
}

func (ps *responsibleService) Get(ctx context.Context, req *models.GetRequest) (*models.GetResponsibleResponse, error) {

	response, err := ps.storage.Responsible().Get(req.ID)

	if err != nil {
		return nil, helper.HandleError(ps.logger, err, "error while getting responsible", req)
	}
	return &models.GetResponsibleResponse{
		Responsible: response,
	}, nil
}
func (ps *responsibleService) GetAll(ctx context.Context, req *models.GetAllResponsiblesRequest) (*models.GetAllResponsiblesResponse, error) {

	response, count, err := ps.storage.Responsible().GetAll(req.Page, req.Limit, req.Name)

	if err != nil {
		return nil, helper.HandleError(ps.logger, err, "error while getting responsible", req)
	}
	return &models.GetAllResponsiblesResponse{
		Count:        count,
		Responsibles: response,
	}, nil
}
