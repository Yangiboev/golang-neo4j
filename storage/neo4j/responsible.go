package neo4j

import (
	"context"

	"github.com/Yangiboev/golang-neo4j/api/models"
	"github.com/Yangiboev/golang-neo4j/config"
	"github.com/Yangiboev/golang-neo4j/storage/repo"

)

type responsibleRepo struct {
	collection *neo4j.Collection
}

func NewResponsibleRepo(db *neo4j.Database) repo.ResponsibleStorageI {
	return &responsibleRepo{
		collection: db.Collection(config.CollectionName)}
}

func (pr *responsibleRepo) Create(responsible *models.Responsible) (string, error) {
	return responsible.ID, nil
}

func (pr *responsibleRepo) Get(id string) (*models.Responsible, error) {

	return &responsible, nil
}

func (pr *responsibleRepo) GetAll(page, limit int64, name string) ([]*models.Responsible, int64, error) {

	return [], 0, nil
}
