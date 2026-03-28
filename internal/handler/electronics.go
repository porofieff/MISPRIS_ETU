package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type createElectronicsInput struct {
	ControllerID string `json:"controller_id" binding:"required"`
	SensorID     string `json:"sensor_id" binding:"required"`
	WiringID     string `json:"wiring_id" binding:"required"`
}

func (h *Handler) CreateElectronics(c *gin.Context) {
	var input createElectronicsInput
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Electronics.Create(c.Request.Context(), input.ControllerID, input.SensorID, input.WiringID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, idResponse{ID: id})
}

func (h *Handler) GetElectronics(c *gin.Context) {
	id := c.Param("id")
	item, err := h.services.Electronics.GetByID(c.Request.Context(), id)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) ListElectronics(c *gin.Context) {
	list, err := h.services.Electronics.List(c.Request.Context())
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, list)
}

func (h *Handler) UpdateElectronics(c *gin.Context) {
	id := c.Param("id")
	var input createElectronicsInput
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.services.Electronics.Update(c.Request.Context(), id, input.ControllerID, input.SensorID, input.WiringID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "updated"})
}

func (h *Handler) DeleteElectronics(c *gin.Context) {
	id := c.Param("id")
	if err := h.services.Electronics.Delete(c.Request.Context(), id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}

type createControllerInput struct {
	Name string `json:"name" binding:"required"`
	Info string `json:"info"`
}

func (h *Handler) CreateController(c *gin.Context) {
	var input createControllerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Controller.Create(c.Request.Context(), input.Name, input.Info)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, idResponse{ID: id})
}

func (h *Handler) GetController(c *gin.Context) {
	id := c.Param("id")
	item, err := h.services.Controller.GetByID(c.Request.Context(), id)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) ListControllers(c *gin.Context) {
	list, err := h.services.Controller.List(c.Request.Context())
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, list)
}

func (h *Handler) UpdateController(c *gin.Context) {
	id := c.Param("id")
	var input createControllerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.services.Controller.Update(c.Request.Context(), id, input.Name, input.Info)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "updated"})
}

func (h *Handler) DeleteController(c *gin.Context) {
	id := c.Param("id")
	if err := h.services.Controller.Delete(c.Request.Context(), id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}

type createSensorInput struct {
	Name string `json:"name" binding:"required"`
	Info string `json:"info"`
}

func (h *Handler) CreateSensor(c *gin.Context) {
	var input createSensorInput
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Sensor.Create(c.Request.Context(), input.Name, input.Info)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, idResponse{ID: id})
}

func (h *Handler) GetSensor(c *gin.Context) {
	id := c.Param("id")
	item, err := h.services.Sensor.GetByID(c.Request.Context(), id)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) ListSensors(c *gin.Context) {
	list, err := h.services.Sensor.List(c.Request.Context())
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, list)
}

func (h *Handler) UpdateSensor(c *gin.Context) {
	id := c.Param("id")
	var input createSensorInput
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.services.Sensor.Update(c.Request.Context(), id, input.Name, input.Info)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "updated"})
}

func (h *Handler) DeleteSensor(c *gin.Context) {
	id := c.Param("id")
	if err := h.services.Sensor.Delete(c.Request.Context(), id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}

type createWiringInput struct {
	Name string `json:"name" binding:"required"`
	Info string `json:"info"`
}

func (h *Handler) CreateWiring(c *gin.Context) {
	var input createWiringInput
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Wiring.Create(c.Request.Context(), input.Name, input.Info)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, idResponse{ID: id})
}

func (h *Handler) GetWiring(c *gin.Context) {
	id := c.Param("id")
	item, err := h.services.Wiring.GetByID(c.Request.Context(), id)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) ListWirings(c *gin.Context) {
	list, err := h.services.Wiring.List(c.Request.Context())
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, list)
}

func (h *Handler) UpdateWiring(c *gin.Context) {
	id := c.Param("id")
	var input createWiringInput
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.services.Wiring.Update(c.Request.Context(), id, input.Name, input.Info)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "updated"})
}

func (h *Handler) DeleteWiring(c *gin.Context) {
	id := c.Param("id")
	if err := h.services.Wiring.Delete(c.Request.Context(), id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}
