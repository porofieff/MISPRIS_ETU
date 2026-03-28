package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type createChassisInput struct {
	FrameID       string `json:"frame_id" binding:"required"`
	SuspensionID  string `json:"suspension_id" binding:"required"`
	BreakSystemID string `json:"break_system_id" binding:"required"`
}

func (h *Handler) CreateChassis(c *gin.Context) {
	var input createChassisInput
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Chassis.Create(c.Request.Context(), input.FrameID, input.SuspensionID, input.BreakSystemID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, idResponse{ID: id})
}

func (h *Handler) GetChassis(c *gin.Context) {
	id := c.Param("id")
	item, err := h.services.Chassis.GetByID(c.Request.Context(), id)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) ListChassis(c *gin.Context) {
	list, err := h.services.Chassis.List(c.Request.Context())
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, list)
}

func (h *Handler) UpdateChassis(c *gin.Context) {
	id := c.Param("id")
	var input createChassisInput
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.services.Chassis.Update(c.Request.Context(), id, input.FrameID, input.SuspensionID, input.BreakSystemID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "updated"})
}

func (h *Handler) DeleteChassis(c *gin.Context) {
	id := c.Param("id")
	if err := h.services.Chassis.Delete(c.Request.Context(), id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}

type createFrameInput struct {
	Name string `json:"name" binding:"required"`
	Info string `json:"info"`
}

func (h *Handler) CreateFrame(c *gin.Context) {
	var input createFrameInput
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Frame.Create(c.Request.Context(), input.Name, input.Info)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, idResponse{ID: id})
}

func (h *Handler) GetFrame(c *gin.Context) {
	id := c.Param("id")
	item, err := h.services.Frame.GetByID(c.Request.Context(), id)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) ListFrames(c *gin.Context) {
	list, err := h.services.Frame.List(c.Request.Context())
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, list)
}

func (h *Handler) UpdateFrame(c *gin.Context) {
	id := c.Param("id")
	var input createFrameInput
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.services.Frame.Update(c.Request.Context(), id, input.Name, input.Info)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "updated"})
}

func (h *Handler) DeleteFrame(c *gin.Context) {
	id := c.Param("id")
	if err := h.services.Frame.Delete(c.Request.Context(), id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}

type createSuspensionInput struct {
	Name string `json:"name" binding:"required"`
	Info string `json:"info"`
}

func (h *Handler) CreateSuspension(c *gin.Context) {
	var input createSuspensionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Suspension.Create(c.Request.Context(), input.Name, input.Info)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, idResponse{ID: id})
}

func (h *Handler) GetSuspension(c *gin.Context) {
	id := c.Param("id")
	item, err := h.services.Suspension.GetByID(c.Request.Context(), id)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) ListSuspensions(c *gin.Context) {
	list, err := h.services.Suspension.List(c.Request.Context())
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, list)
}

func (h *Handler) UpdateSuspension(c *gin.Context) {
	id := c.Param("id")
	var input createSuspensionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.services.Suspension.Update(c.Request.Context(), id, input.Name, input.Info)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "updated"})
}

func (h *Handler) DeleteSuspension(c *gin.Context) {
	id := c.Param("id")
	if err := h.services.Suspension.Delete(c.Request.Context(), id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}

type createBreakSystemInput struct {
	Name string `json:"name" binding:"required"`
	Info string `json:"info"`
}

func (h *Handler) CreateBreakSystem(c *gin.Context) {
	var input createBreakSystemInput
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.BreakSystem.Create(c.Request.Context(), input.Name, input.Info)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, idResponse{ID: id})
}

func (h *Handler) GetBreakSystem(c *gin.Context) {
	id := c.Param("id")
	item, err := h.services.BreakSystem.GetByID(c.Request.Context(), id)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) ListBreakSystems(c *gin.Context) {
	list, err := h.services.BreakSystem.List(c.Request.Context())
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, list)
}

func (h *Handler) UpdateBreakSystem(c *gin.Context) {
	id := c.Param("id")
	var input createBreakSystemInput
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.services.BreakSystem.Update(c.Request.Context(), id, input.Name, input.Info)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "updated"})
}

func (h *Handler) DeleteBreakSystem(c *gin.Context) {
	id := c.Param("id")
	if err := h.services.BreakSystem.Delete(c.Request.Context(), id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}
