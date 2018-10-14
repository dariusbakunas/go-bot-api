package controllers

import (
	"fmt"
	"github.com/dariusbakunas/go-bot-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ServosController struct{
	ssc32u *models.SSC32U
}

func (s ServosController) Turn(c *gin.Context) {
	angle, err := strconv.Atoi(c.Query("angle"))

	if err != nil || angle < 0 || angle > 180 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid angle"})
		c.Abort()
		return
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil || id < 0 || id > 31 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Servo ID must be between 0 and 31"})
	}

	// TODO: make this configurable
	pulseRange := 2100.0 - 900
	pulse :=  int(900 + pulseRange/180.0 * float64(angle))

	command := fmt.Sprintf("#%d P%d \r", id, pulse)
	_, err = s.ssc32u.Write(command)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func ServosControllerInit(ssc32u *models.SSC32U) *ServosController {
	return &ServosController{ssc32u}
}