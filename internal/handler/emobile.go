package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type createEmobileInput struct {
	Name            string `json:"name" binding:"required"`
	PowerPointID    string `json:"power_point_id" binding:"required"`
	BatteryID       string `json:"battery_id" binding:"required"`
	ChargerSystemID string `json:"charger_system_id" binding:"required"`
	ChassisID       string `json:"chassis_id" binding:"required"`
	BodyID          string `json:"body_id" binding:"required"`
	ElectronicsID   string `json:"electronics_id" binding:"required"`
}

func (h *Handler) CreateEmobile(c *gin.Context) {
	var input createEmobileInput
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Emobile.Create(c.Request.Context(), input.Name,
		input.PowerPointID, input.BatteryID, input.ChargerSystemID,
		input.ChassisID, input.BodyID, input.ElectronicsID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, idResponse{ID: id})
}

func (h *Handler) GetEmobile(c *gin.Context) {
	id := c.Param("id")
	emobile, err := h.services.Emobile.GetByID(c.Request.Context(), id)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, emobile)
}

func (h *Handler) ListEmobiles(c *gin.Context) {
	list, err := h.services.Emobile.List(c.Request.Context())
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, list)
}

type updateEmobileInput struct {
	Name            string `json:"name"`
	PowerPointID    string `json:"power_point_id"`
	BatteryID       string `json:"battery_id"`
	ChargerSystemID string `json:"charger_system_id"`
	ChassisID       string `json:"chassis_id"`
	BodyID          string `json:"body_id"`
	ElectronicsID   string `json:"electronics_id"`
}

func (h *Handler) UpdateEmobile(c *gin.Context) {
	id := c.Param("id")
	var input updateEmobileInput
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.services.Emobile.Update(c.Request.Context(), id, input.Name,
		input.PowerPointID, input.BatteryID, input.ChargerSystemID,
		input.ChassisID, input.BodyID, input.ElectronicsID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "updated"})
}

func (h *Handler) DeleteEmobile(c *gin.Context) {
	id := c.Param("id")
	if err := h.services.Emobile.Delete(c.Request.Context(), id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}
