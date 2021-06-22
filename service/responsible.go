package service

import (
	"context"
	"fmt"

	"github.com/Yangiboev/golang-neo4j/api/models"
	"github.com/Yangiboev/golang-neo4j/pkg/helper"
	"github.com/Yangiboev/golang-neo4j/pkg/logger"
	"github.com/Yangiboev/golang-neo4j/storage"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type responsibleService struct {
	logger  logger.Logger
	storage storage.StorageI
}

func NewResponsibleService(driver neo4j.Driver, log logger.Logger) *responsibleService {
	return &responsibleService{
		storage: storage.NewResponsibleStorage(driver),
		logger:  log,
	}
}

func (ps *responsibleService) Create(ctx context.Context, req *models.Responsible) (*models.CreateResponse, error) {
	res, err := ps.storage.Responsible().Create(req)

	if err != nil {
		ps.logger.Error("error while creating responsible", logger.Error(err))
		return nil, helper.HandleError(ps.logger, err, "error while creating responsible", req)
	}
	return res, nil
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

func (ps *responsibleService) Update(ctx context.Context, req *models.Responsible) error {
	err := ps.storage.Responsible().Update(req)
	if err != nil {
		return helper.HandleError(ps.logger, err, "error while updating responsible", nil)
	}

	return nil
}

func (ps *responsibleService) Delete(ctx context.Context, req *models.DeleteRequest) error {
	fmt.Println("Erererereerrers")
	err := ps.storage.Responsible().Delete(req.ID)

	if err != nil {
		return helper.HandleError(ps.logger, err, "error while deleting responsible", nil)
	}

	return nil
}
