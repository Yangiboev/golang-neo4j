package repo

import (
	"github.com/Yangiboev/golang-neo4j/api/models"
)

type ResponsibleStorageI interface {
	Create(responsible *models.Responsible) (*models.CreateResponse, error)
	Get(id string) (*models.Responsible, error)
	GetAll(page, limit int32, name string) ([]*models.Responsible, int64, error)
	Update(responsible *models.Responsible) error
	Delete(id string) error
}
