package storage

import (
	"github.com/Yangiboev/golang-neo4j/storage/neo4j"
	"github.com/Yangiboev/golang-neo4j/storage/repo"
)

type StorageI interface {
	Responsible() repo.ResponsibleStorageI
}

type responsible struct {
	responsibleRepo repo.ResponsibleStorageI
}

func NewResponsibleStorage(db *db.Database) StorageI {
	return &responsible{
		responsibleRepo: neo4j.NewResponsibleRepo(db),
	}
}

func (pr responsible) Responsible() repo.ResponsibleStorageI {
	return pr.responsibleRepo
}
