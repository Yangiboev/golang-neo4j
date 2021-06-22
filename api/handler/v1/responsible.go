package v1

import (
	"net/http"

	"github.com/Yangiboev/golang-neo4j/api/models"
	"github.com/Yangiboev/golang-neo4j/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Router /v1/responsible/{responsible_id} [get]
// @Summary Get responsible
// @Description API for getting a responsible
// @Tags responsible
// @Accept  json
// @Produce  json
// @Param responsible_id path string true "responsible_id"
// @Success 200 {object} models.GetResponsibleResponse
// @Failure 400 {object} models.BadRequestError
// @Failure 500 {object} models.InternalServerError
func (h *handlerV1) GetResponsible(c *gin.Context) {
	responsibleID := c.Param("responsible_id")
	_, err := uuid.Parse(responsibleID)
	if err != nil {
		HandleInternalServerError(c, err, "responsible_id format is invalid format!")
		return
	}
	responsible, err := h.storage.Responsible().Get(responsibleID)

	if err != nil {
		HandleInternalServerError(c, err, "error while getting responsible")
		return
	}

	c.JSON(http.StatusOK,
		gin.H{
			"success": true,
			"data":    responsible,
		})

}

//@Router /v1/responsible [post]
//@Summary Create responsible
//@Description API for creating responsible
//@Tags responsible
//@Accept json
//@Produce json
//@Param responsible body models.CreateUpdateResponsibleRequest  true "responsible"
// @Success 200 {object} models.CreateResponse
// @Failure 400 {object} models.BadRequestError
// @Failure 500 {object} models.InternalServerError
func (h *handlerV1) CreateResponsible(c *gin.Context) {

	var (
		responsible models.Responsible
	)
	err := c.ShouldBindJSON(&responsible)
	if err != nil {
		HandleBadRequest(c, err, "error while binding responsible to json")
		return
	}

	id, err := uuid.NewRandom()
	if err != nil {
		HandleInternalServerError(c, err, "error while generating uuid")
		return
	}

	responsible.ID = id.String()
	resp, err := h.storage.Responsible().Create(
		&responsible)

	if err != nil {
		h.log.Error("error while creating responsible", logger.Error(err))
		HandleInternalServerError(c, err, "Error while creating  responsible")
		return
	}

	c.JSON(http.StatusCreated,
		gin.H{
			"success": true,
			"data":    resp,
		})
}

// @Router /v1/responsible [get]
// @Summary Get All responsibles
// @Description API for getting all responsibles
// @Tags responsible
// @Accept  json
// @Produce  json
// @Param name path string false "name"
// @Success 200 {object} models.GetAllResponsiblesResponse
// @Failure 400 {object} models.BadRequestError
// @Failure 500 {object} models.InternalServerError
func (h *handlerV1) GetAllResponsibles(c *gin.Context) {
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
	responsibles, count, err := h.storage.Responsible().GetAll(int32(page), int32(limit), c.Query("name"))

	if err != nil {
		HandleBadRequest(c, err, "Error while getting all responsibles")
		return
	}

	c.JSON(http.StatusOK,
		gin.H{
			"success": true,
			"count":   count,
			"data":    responsibles,
		})

}

//@Router /v1/responsible/{responsible_id} [put]
//@Summary Update responsible
//@Description API for creating responsible
//@Tags responsible
//@Accept json
//@Produce json
// @Param responsible_id path string true "responsible_id"
//@Param responsible body models.UpdateResponsibleRequest  true "responsible"
// @Success 200 {object} models.CreateResponse
// @Failure 400 {object} models.BadRequestError
// @Failure 500 {object} models.InternalServerError
func (h *handlerV1) UpdateResponsible(c *gin.Context) {
	var (
		responsible   models.Responsible
		responsibleID string
	)
	responsibleID = c.Param("responsible_id")
	_, err := uuid.Parse(responsibleID)
	if err != nil {
		HandleInternalServerError(c, err, "responsible_id format is an invalid format!")
		return
	}

	err = c.ShouldBindJSON(&responsible)
	if err != nil {
		HandleBadRequest(c, err, "Error while binding responsible!")
		return
	}
	responsible.ID = responsibleID
	err = h.storage.Responsible().Update(&responsible)

	if err != nil {
		HandleBadRequest(c, err, "Error while updating responsible")
		return
	}

	c.JSON(http.StatusCreated,
		gin.H{
			"success": true,
		})
}

// @Router /v1/responsible/{responsible_id} [delete]
// @Summary Delete Responsible
// @Description API for deleting responsible
// @Tags responsible
// @Accept json
// @Produce json
// @Param responsible_id path string  true "responsible_id"
// @Success 200 {object} models.SuccessResponse
// @Failure 400 {object} models.BadRequestError
// @Failure 500 {object} models.InternalServerError
func (h *handlerV1) DeleteResponsible(c *gin.Context) {
	responsibleId := c.Param("responsible_id")

	_, err := uuid.Parse(responsibleId)
	if err != nil {
		HandleBadRequest(c, err, "Error while parsing uuid")
		return
	}

	err = h.storage.Responsible().Delete(responsibleId)

	if err != nil {
		HandleBadRequest(c, err, "Error while deleting responsible")
		return
	}

	c.JSON(http.StatusOK,
		gin.H{
			"success": true,
		})
}
