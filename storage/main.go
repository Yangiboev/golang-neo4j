package storage

import (
	neo "github.com/Yangiboev/golang-neo4j/storage/neo4j"
	"github.com/Yangiboev/golang-neo4j/storage/repo"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type StorageI interface {
	Responsible() repo.ResponsibleStorageI
	Action() repo.ActionI
}

type responsible struct {
	responsibleRepo repo.ResponsibleStorageI
	actionRepo      repo.ActionI
}

func NewResponsibleStorage(driver neo4j.Driver) StorageI {
	return &responsible{
		responsibleRepo: neo.NewResponsibleRepo(driver),
		actionRepo:      neo.NewActionRepo(driver),
	}
}

func (pr responsible) Responsible() repo.ResponsibleStorageI {
	return pr.responsibleRepo
}

func (pr responsible) Action() repo.ActionI {
	return pr.actionRepo
}
