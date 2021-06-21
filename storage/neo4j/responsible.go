package neo4j

import (
	"fmt"

	"github.com/Yangiboev/golang-neo4j/api/models"
	"github.com/Yangiboev/golang-neo4j/storage/repo"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type responsibleRepo struct {
	driver neo4j.Driver
}

func NewResponsibleRepo(driver neo4j.Driver) repo.ResponsibleStorageI {
	return &responsibleRepo{
		driver: driver,
	}
}

func (pr *responsibleRepo) Create(responsible *models.Responsible) (*models.CreateResponse, error) {
	session := pr.driver.NewSession(neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeWrite,
	})
	defer session.Close()
	result, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		records, err := tx.Run(`CREATE (r:Responsible {
			id:$id, 
			name_of_step: $name_of_step,
			organization:$organization,
			role:$role,
			comment: $comment,
			created_at: $created_at,
			updated_at: $updated_at
			}) RETURN r.id;`, map[string]interface{}{
			"id":           responsible.ID,
			"name_of_step": responsible.NameOfStep,
			"organization": responsible.Organization,
			"role":         responsible.Role,
			"comment":      responsible.Comment,
			"created_at":   responsible.CreatedAt,
			"updated_at":   responsible.UpdatedAt,
		})
		if err != nil {
			return nil, err
		}
		record, err := records.Single()
		if err != nil {
			return nil, err
		}
		return &models.CreateResponse{
			ID: record.Values[0].(string),
		}, nil
	})
	if err != nil {
		return nil, err
	}
	return result.(*models.CreateResponse), nil
}

func (pr *responsibleRepo) Get(id string) (*models.Responsible, error) {
	session := pr.driver.NewSession(neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeRead,
	})
	defer session.Close()
	responsible, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		_, err := tx.Run(
			`MATCH (r) WHERE id(r)=$id return r`, map[string]interface{}{
				"id": id,
			})
		fmt.Println(err)

		return &models.Responsible{}, nil
	})
	fmt.Println(err)
	fmt.Println(responsible)

	return nil, nil
}

func (pr *responsibleRepo) GetAll(page, limit int32, name string) ([]*models.Responsible, int64, error) {

	return nil, 0, nil
}
