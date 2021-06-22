package neo4j

import (
	"fmt"

	"github.com/Yangiboev/golang-neo4j/api/models"
	"github.com/Yangiboev/golang-neo4j/storage/repo"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type actionRepo struct {
	driver neo4j.Driver
}

func NewActionRepo(driver neo4j.Driver) repo.ActionI {
	return &actionRepo{
		driver: driver,
	}
}

func (ar *actionRepo) Create(action *models.Action) (*models.CreateResponse, error) {
	query := `
		CREATE (a:Action {
			id: $id,
			type: $type,
			status: $status,
			comment: $comment,
			role: $role,
			created_at: $created_at
		}) RETURN a.id;`

	session := ar.driver.NewSession(neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeWrite,
	})
	defer session.Close()
	result, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		records, err := tx.Run(
			query, map[string]interface{}{
				"id":         action.ID,
				"type":       action.Type,
				"status":     action.Status,
				"comment":    action.Comment,
				"role":       action.Role,
				"created_at": action.CreatedAt,
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

func (ar *actionRepo) Get(id string) (*models.Action, error) {
	query := `
		MATCH (a:Action {id: $id}) RETURN
		a.id,
		a.type,
		a.status,
		a.comment,
		a.role,
		a.created_at`

	session := ar.driver.NewSession(neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeRead,
	})
	defer session.Close()
	action, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		result, err := tx.Run(
			query, map[string]interface{}{
				"id": id,
			})
		if err != nil {
			return nil, err
		}

		record, err := result.Single()
		if err != nil {
			return nil, err
		}

		return scanAction(record), nil
	})
	if err != nil {
		return nil, err
	}

	return action.(*models.Action), nil
}

func (ar *actionRepo) GetAll(page, limit int32) ([]*models.Action, int64, error) {
	var (
		count int64
		query = `
		MATCH (a:Action) RETURN
		a.id,
		a.type,
		a.status,
		a.comment,
		a.role,
		a.created_at
		ORDER BY a.created_at
		SKIP $skip
		LIMIT $limit`
		countQuery = `MATCH (a:Action) RETURN count(a) as count`
		session    = ar.driver.NewSession(neo4j.SessionConfig{
			AccessMode: neo4j.AccessModeRead,
		})
	)

	res, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		var actions []*models.Action
		result, err := tx.Run(
			countQuery, map[string]interface{}{})
		if err != nil {
			return nil, err
		}
		record, err := result.Single()
		if err != nil {
			return nil, err
		}

		fmt.Println(record.Keys)
		countRes, _ := record.Get("count")
		count = countRes.(int64)

		result, err = tx.Run(
			query, map[string]interface{}{
				"skip":  page * limit,
				"limit": limit,
			})
		if err != nil {
			return nil, err
		}
		for result.Next() {
			var record = result.Record()

			actions = append(actions, scanAction(record))
		}

		return actions, nil
	})

	if err != nil {
		return nil, 0, err
	}
	fmt.Println(count)
	fmt.Println(count)
	fmt.Println(count)
	return res.([]*models.Action), count, nil
}

func (ar *actionRepo) Update(action *models.Action) error {
	query := `
		MATCH(a:Action {id: $id}) SET
		a.type=$type,
		a.status=$status,
		a.comment=$comment,
		a.role=$role,
		a.created_at=$created_at`

	session := ar.driver.NewSession(neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeWrite,
	})
	defer session.Close()
	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		_, err := tx.Run(
			query, map[string]interface{}{
				"id":         action.ID,
				"type":       action.Type,
				"status":     action.Status,
				"comment":    action.Comment,
				"role":       action.Role,
				"created_at": action.CreatedAt,
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

func (ar *actionRepo) Delete(id string) error {
	query := `
		MATCH (a:Action {id: $id}) DELETE a`

	session := ar.driver.NewSession(neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeWrite,
	})
	defer session.Close()
	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		_, err := tx.Run(
			query, map[string]interface{}{
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

func scanAction(record *neo4j.Record) *models.Action {
	var (
		ID, _         = record.Get("a.id")
		actionType, _ = record.Get("a.type")
		status, _     = record.Get("a.status")
		comment, _    = record.Get("a.comment")
		role, _       = record.Get("a.role")
		createdAt, _  = record.Get("a.created_at")
	)
	return &models.Action{
		ID:        ID.(string),
		Type:      actionType.(string),
		Status:    status.(string),
		Comment:   comment.(string),
		Role:      role.(string),
		CreatedAt: createdAt.(int64),
	}
}
