package neo4j_test

import (
	"testing"

	"github.com/Yangiboev/golang-neo4j/api/models"
	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func createResponsible(t *testing.T) *models.Responsible {
	var req *models.Responsible
	err := faker.FakeData(&req)
	assert.NoError(t, err)
	ID, _ := uuid.NewRandom()
	req.ID = ID.String()
	response, err := strg.Responsible().Create(req)
	assert.NoError(t, err)
	assert.NotEmpty(t, response)
	return &models.Responsible{
		ID:           req.ID,
		Organization: req.Organization,
		NameOfStep:   req.NameOfStep,
		Role:         req.Role,
		Comment:      req.Comment,
		CreatedAt:    req.CreatedAt,
		UpdatedAt:    req.UpdatedAt,
	}
}
func TestCreateResponsible(t *testing.T) {
	for i := 0; i < 1000; i++ {
		_ = createResponsible(t)
	}
}
func TestGetResponsible(t *testing.T) {
	created := createResponsible(t)
	response, err := strg.Responsible().Get(created.ID)
	assert.NoError(t, err)
	assert.Equal(t, created, response)
}
func TestGetAllResponsible(t *testing.T) {
	_ = createResponsible(t)
	response, count, err := strg.Responsible().GetAll(1, 20, "")
	assert.NoError(t, err)
	assert.NotEmpty(t, response)
	assert.True(t, count > 0)
}
