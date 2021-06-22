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
		result, err := tx.Run(
			`MATCH (r:Responsible {id:$id}) return 
			r.id,
			r.name_of_step,
			r.organization,
			r.role,
			r.comment,
			r.created_at,
			r.updated_at
			`, map[string]interface{}{
				"id": id,
			})
		if err != nil {
			return nil, err
		}
		record, err := result.Single()
		// var (
		// 	ID, _           = record.Get("r.id")
		// 	nameOfStep, _   = record.Get("r.name_of_step")
		// 	organization, _ = record.Get("r.organization")
		// 	role, _         = record.Get("r.role")
		// 	comment, _      = record.Get("r.comment")
		// 	createdAt, _    = record.Get("r.created_at")
		// 	updatedAt, _    = record.Get("r.updated_at")
		// )
		if err != nil {
			return nil, err
		}
		return scan(record), nil
	})
	if err != nil {
		return nil, err
	}

	return responsible.(*models.Responsible), nil
}

func (pr *responsibleRepo) GetAll(page, limit int32, name string) ([]*models.Responsible, int64, error) {

	session := pr.driver.NewSession(neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeRead,
	})
	fmt.Println("ssssssssssss")

	defer session.Close()
	res, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		var results []*models.Responsible
		fmt.Println(page)
		fmt.Println(limit)
		responsibles, err := tx.Run(
			`MATCH (r:Responsible) RETURN 
			r.id,
			r.name_of_step,
			r.organization,
			r.role,
			r.comment,
			r.created_at,
			r.updated_at
			ORDER BY r.created_at
			SKIP $skip
			LIMIT $limit
			`, map[string]interface{}{
				"skip":  page * limit,
				"limit": limit,
			})
		fmt.Println((page - 1) * limit)
		if err != nil {
			return nil, err
		}
		for responsibles.Next() {
			var (
				record = responsibles.Record()
				// ID, _           = record.Get("r.id")
				// nameOfStep, _   = record.Get("r.name_of_step")
				// organization, _ = record.Get("r.organization")
				// role, _         = record.Get("r.role")
				// comment, _      = record.Get("r.comment")
				// createdAt, _    = record.Get("r.created_at")
				// updatedAt, _    = record.Get("r.updated_at")
			)

			results = append(results, scan(record))
		}
		return results, nil
	})
	if err != nil {
		return nil, 0, err
	}
	return res.([]*models.Responsible), 0, nil
}
func scan(record *neo4j.Record) *models.Responsible {
	var (
		ID, _           = record.Get("r.id")
		nameOfStep, _   = record.Get("r.name_of_step")
		organization, _ = record.Get("r.organization")
		role, _         = record.Get("r.role")
		comment, _      = record.Get("r.comment")
		createdAt, _    = record.Get("r.created_at")
		updatedAt, _    = record.Get("r.updated_at")
	)
	return &models.Responsible{
		ID:           ID.(string),
		NameOfStep:   nameOfStep.(string),
		Organization: organization.(string),
		Role:         role.(string),
		Comment:      comment.(string),
		CreatedAt:    createdAt.(int64),
		UpdatedAt:    updatedAt.(int64),
	}
}
