package repo

import (
	"github.com/Yangiboev/golang-neo4j/api/models"
)

type ActionI interface {
	Create(action *models.Action) (*models.CreateResponse, error)
	Get(id string) (*models.Action, error)
	GetAll(page, limit int32) ([]*models.Action, int64, error)
	Update(action *models.Action) error
	Delete(id string) error
}
