package repo

import (
	"github.com/Yangiboev/golang-neo4j/api/models"
)

type ResponsibleStorageI interface {
	Create(responsible *models.Responsible) (string, error)
	Get(id string) (*models.Responsible, error)
	GetAll(page, limit int64, name string) ([]*models.Responsible, int64, error)
}
