package neo4j_test

import (
	"testing"

	"github.com/Yangiboev/golang-neo4j/api/models"
	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func createAction(t *testing.T) *models.Action {
	var action *models.Action
	err := faker.FakeData(&action)
	assert.NoError(t, err)
	id, _ := uuid.NewRandom()
	action.ID = id.String()
	res, err := strg.Action().Create(action)
	assert.NoError(t, err)
	assert.NotEmpty(t, res)

	return action
}

func deleteAction(t *testing.T, id string) {
	err := strg.Action().Delete(id)
	assert.NoError(t, err)
}

func TestCreateAction(t *testing.T) {
	res := createAction(t)

	assert.NotEmpty(t, res)
	deleteAction(t, res.ID)
}

func TestGetAction(t *testing.T) {
	action := createAction(t)

	res, err := strg.Action().Get(action.ID)

	assert.NoError(t, err)
	assert.Equal(t, res, action)
	deleteAction(t, action.ID)
}

func TestGetAllActions(t *testing.T) {
	res, _, err := strg.Action().GetAll(1, 10)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}
