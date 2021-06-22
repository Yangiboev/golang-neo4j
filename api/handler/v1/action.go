package v1

import (
	"net/http"

	"github.com/Yangiboev/golang-neo4j/api/models"
	"github.com/Yangiboev/golang-neo4j/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Router /v1/action/{action_id} [get]
// @Summary Get action
// @Description API for getting a action
// @Tags action
// @Accept  json
// @Produce  json
// @Param action_id path string true "action_id"
// @Success 200 {object} models.Action
// @Failure 400 {object} models.BadRequestError
// @Failure 500 {object} models.InternalServerError
func (h *handlerV1) GetAction(c *gin.Context) {
	actionID := c.Param("action_id")
	_, err := uuid.Parse(actionID)
	if err != nil {
		HandleInternalServerError(c, err, "action_id format is invalid format!")
		return
	}
	action, err := h.storage.Action().Get(actionID)

	if err != nil {
		HandleInternalServerError(c, err, "error while getting action")
		return
	}

	c.JSON(http.StatusOK, action)

}

//@Router /v1/action [post]
//@Summary Create action
//@Description API for creating action
//@Tags action
//@Accept json
//@Produce json
//@Param action body models.CreateActionRequest  true "action"
// @Success 200 {object} models.CreateResponse
// @Failure 400 {object} models.BadRequestError
// @Failure 500 {object} models.InternalServerError
func (h *handlerV1) CreateAction(c *gin.Context) {

	var (
		action models.Action
	)
	err := c.ShouldBindJSON(&action)
	if err != nil {
		HandleBadRequest(c, err, "error while binding action to json")
		return
	}

	id, err := uuid.NewRandom()
	if err != nil {
		HandleInternalServerError(c, err, "error while generating uuid")
		return
	}

	action.ID = id.String()
	resp, err := h.storage.Action().Create(
		&action)

	if err != nil {
		h.log.Error("error while creating action", logger.Error(err))
		HandleInternalServerError(c, err, "Error while creating  action")
		return
	}

	c.JSON(http.StatusCreated,
		gin.H{
			"success": true,
			"data":    resp,
		})
}

// @Router /v1/action [get]
// @Summary Get All actions
// @Description API for getting all actions
// @Tags action
// @Accept  json
// @Produce  json
// @Success 200 {object} models.GetAllActionsResponse
// @Failure 400 {object} models.BadRequestError
// @Failure 500 {object} models.InternalServerError
func (h *handlerV1) GetAllActions(c *gin.Context) {
	page, err := ParsePageQueryParam(c)
	if err != nil {
		HandleBadRequest(c, err, "Error while parsing page")
		return
	}
	limit, err := ParseLimitQueryParam(c)
	if err != nil {
		HandleBadRequest(c, err, "Error while parsing page")
		return
	}
	actions, count, err := h.storage.Action().GetAll(int32(page), int32(limit))

	if err != nil {
		HandleBadRequest(c, err, "Error while getting all actions")
		return
	}

	c.JSON(http.StatusOK,
		gin.H{
			"success": true,
			"count":   count,
			"data":    actions,
		})

}

//@Router /v1/action/{action_id} [put]
//@Summary Update action
//@Description API for creating action
//@Tags action
//@Accept json
//@Produce json
// @Param action_id path string true "action_id"
//@Param action body models.UpdateActionRequest  true "action"
// @Success 200 {object} models.CreateResponse
// @Failure 400 {object} models.BadRequestError
// @Failure 500 {object} models.InternalServerError
func (h *handlerV1) UpdateAction(c *gin.Context) {
	var (
		action   models.Action
		actionID string
	)
	actionID = c.Param("action_id")
	_, err := uuid.Parse(actionID)
	if err != nil {
		HandleInternalServerError(c, err, "action_id format is an invalid format!")
		return
	}

	err = c.ShouldBindJSON(&action)
	if err != nil {
		HandleBadRequest(c, err, "Error while binding action!")
		return
	}
	action.ID = actionID
	err = h.storage.Action().Update(&action)

	if err != nil {
		HandleBadRequest(c, err, "Error while updating action")
		return
	}

	c.JSON(http.StatusCreated,
		gin.H{
			"success": true,
		})
}

// @Router /v1/action/{action_id} [delete]
// @Summary Delete action
// @Description API for deleting action
// @Tags action
// @Accept json
// @Produce json
// @Param action_id path string  true "action_id"
// @Success 200 {object} models.SuccessResponse
// @Failure 400 {object} models.BadRequestError
// @Failure 500 {object} models.InternalServerError
func (h *handlerV1) DeleteAction(c *gin.Context) {
	actionId := c.Param("action_id")

	_, err := uuid.Parse(actionId)
	if err != nil {
		HandleBadRequest(c, err, "Error while parsing uuid")
		return
	}

	err = h.storage.Action().Delete(actionId)

	if err != nil {
		HandleBadRequest(c, err, "Error while deleting action")
		return
	}

	c.JSON(http.StatusOK,
		gin.H{
			"success": true,
		})
}
