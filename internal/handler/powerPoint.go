package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type createPowerPointInput struct {
	EngineID   string `json:"engine_id" binding:"required"`
	InverterID string `json:"inverter_id" binding:"required"`
	GearboxID  string `json:"gearbox_id" binding:"required"`
}

func (h *Handler) CreatePowerPoint(c *gin.Context) {
	var input createPowerPointInput
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.PowerPoint.Create(c.Request.Context(), input.EngineID, input.InverterID, input.GearboxID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, idResponse{ID: id})
}

func (h *Handler) GetPowerPoint(c *gin.Context) {
	id := c.Param("id")
	item, err := h.services.PowerPoint.GetByID(c.Request.Context(), id)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) ListPowerPoints(c *gin.Context) {
	list, err := h.services.PowerPoint.List(c.Request.Context())
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, list)
}

func (h *Handler) UpdatePowerPoint(c *gin.Context) {
	id := c.Param("id")
	var input createPowerPointInput
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.services.PowerPoint.Update(c.Request.Context(), id, input.EngineID, input.InverterID, input.GearboxID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "updated"})
}

func (h *Handler) DeletePowerPoint(c *gin.Context) {
	id := c.Param("id")
	if err := h.services.PowerPoint.Delete(c.Request.Context(), id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}

type createEngineInput struct {
	Name       string `json:"name" binding:"required"`
	EngineType string `json:"engine_type"`
	Info       string `json:"info"`
}

func (h *Handler) CreateEngine(c *gin.Context) {
	var input createEngineInput
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Engine.Create(c.Request.Context(), input.Name, input.EngineType, input.Info)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, idResponse{ID: id})
}

func (h *Handler) GetEngine(c *gin.Context) {
	id := c.Param("id")
	item, err := h.services.Engine.GetByID(c.Request.Context(), id)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) ListEngines(c *gin.Context) {
	list, err := h.services.Engine.List(c.Request.Context())
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, list)
}

func (h *Handler) UpdateEngine(c *gin.Context) {
	id := c.Param("id")
	var input createEngineInput
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.services.Engine.Update(c.Request.Context(), id, input.Name, input.EngineType, input.Info)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "updated"})
}

func (h *Handler) DeleteEngine(c *gin.Context) {
	id := c.Param("id")
	if err := h.services.Engine.Delete(c.Request.Context(), id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}

type createInverterInput struct {
	Name string `json:"name" binding:"required"`
	Info string `json:"info"`
}

func (h *Handler) CreateInverter(c *gin.Context) {
	var input createInverterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Inverter.Create(c.Request.Context(), input.Name, input.Info)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, idResponse{ID: id})
}

func (h *Handler) GetInverter(c *gin.Context) {
	id := c.Param("id")
	item, err := h.services.Inverter.GetByID(c.Request.Context(), id)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) ListInverters(c *gin.Context) {
	list, err := h.services.Inverter.List(c.Request.Context())
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, list)
}

func (h *Handler) UpdateInverter(c *gin.Context) {
	id := c.Param("id")
	var input createInverterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.services.Inverter.Update(c.Request.Context(), id, input.Name, input.Info)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "updated"})
}

func (h *Handler) DeleteInverter(c *gin.Context) {
	id := c.Param("id")
	if err := h.services.Inverter.Delete(c.Request.Context(), id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}

type createGearboxInput struct {
	Name string `json:"name" binding:"required"`
	Info string `json:"info"`
}

func (h *Handler) CreateGearbox(c *gin.Context) {
	var input createGearboxInput
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Gearbox.Create(c.Request.Context(), input.Name, input.Info)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, idResponse{ID: id})
}

func (h *Handler) GetGearbox(c *gin.Context) {
	id := c.Param("id")
	item, err := h.services.Gearbox.GetByID(c.Request.Context(), id)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) ListGearboxes(c *gin.Context) {
	list, err := h.services.Gearbox.List(c.Request.Context())
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, list)
}

func (h *Handler) UpdateGearbox(c *gin.Context) {
	id := c.Param("id")
	var input createGearboxInput
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.services.Gearbox.Update(c.Request.Context(), id, input.Name, input.Info)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "updated"})
}

func (h *Handler) DeleteGearbox(c *gin.Context) {
	id := c.Param("id")
	if err := h.services.Gearbox.Delete(c.Request.Context(), id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}
