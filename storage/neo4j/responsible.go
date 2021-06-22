package neo4j

import (
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
	var (
		responsible models.Responsible
	)
	session := pr.driver.NewSession(neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeRead,
	})
	query := `MATCH (r:Responsible {id: $id}) return r`
	defer session.Close()
	_, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		get, err := tx.Run(
			query, map[string]interface{}{
				"id": id,
			})

		if err != nil {
			return nil, err
		}

		singleRecord, err := get.Single()
		if err != nil {
			return nil, err
		}
		return &models.Responsible{
			ID:           singleRecord.Values[0].(string),
			NameOfStep:   singleRecord.Values[1].(string),
			Organization: singleRecord.Values[2].(string),
			Role:         singleRecord.Values[3].(string),
			Comment:      singleRecord.Values[4].(string),
			CreatedAt:    singleRecord.Values[5].(int32),
			UpdatedAt:    singleRecord.Values[6].(int32),
		}, nil
	})

	if err != nil {
		return nil, err
	}

	return &responsible, nil
}

func (pr *responsibleRepo) GetAll(page, limit int32, name string) ([]*models.Responsible, int64, error) {

	return nil, 0, nil
}

func (pr *responsibleRepo) Update(responsible *models.Responsible) error {
	query := `MATCH (r:Responsible {id:$id}) SET
		r.name_of_step=$name_of_step,
		r.organization=$organization,
		r.role=$role,
		r.comment=$comment,
		r.updated_at=$updated_at`
	session := pr.driver.NewSession(neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeWrite,
	})
	defer session.Close()
	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		_, err := tx.Run(query, map[string]interface{}{
			"id":           responsible.ID,
			"name_of_step": responsible.NameOfStep,
			"organization": responsible.Organization,
			"role":         responsible.Role,
			"comment":      responsible.Comment,
			"updated_at":   responsible.UpdatedAt,
		})

		if err != nil {
			return nil, err
		}

		return nil, nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (pr *responsibleRepo) Delete(id string) error {
	session := pr.driver.NewSession(neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeWrite,
	})
	defer session.Close()
	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		_, err := tx.Run(`MATCH (r:Responsible {
			id: $id
			}) DELETE r`, map[string]interface{}{
			"id": id,
		})

		if err != nil {
			return nil, err
		}

		return nil, nil
	})

	if err != nil {
		return err
	}

	return nil
}
