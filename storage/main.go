package storage

import (
	neo "github.com/Yangiboev/golang-neo4j/storage/neo4j"
	"github.com/Yangiboev/golang-neo4j/storage/repo"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type StorageI interface {
	Responsible() repo.ResponsibleStorageI
}

type responsible struct {
	responsibleRepo repo.ResponsibleStorageI
}

func NewResponsibleStorage(driver neo4j.Driver) StorageI {
	return &responsible{
		responsibleRepo: neo.NewResponsibleRepo(driver),
	}
}

func (pr responsible) Responsible() repo.ResponsibleStorageI {
	return pr.responsibleRepo
}
