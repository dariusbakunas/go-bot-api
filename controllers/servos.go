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
	id := c.Param("id")

	angle, err := strconv.Atoi(c.Query("angle"))

	if err != nil || angle < 0 || angle > 180 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid angle"})
	}

	c.JSON(http.StatusOK, gin.H{"status": fmt.Sprintf("%s: %s", id, angle)})
}

func ServosControllerInit(ssc32u *models.SSC32U) *ServosController {
	return &ServosController{ssc32u}
}