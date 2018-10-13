package controllers

import (
	"fmt"
	"github.com/dariusbakunas/go-bot-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ServosController struct{
	ssc32u *models.SSC32U
}

func (s ServosController) Turn(c *gin.Context) {
	id := c.Param("id")
	angle := c.Query("angle")
	c.JSON(http.StatusOK, gin.H{"status": fmt.Sprintf("%s: %s", id, angle)})
}

func ServosControllerInit(ssc32u *models.SSC32U) *ServosController {
	return &ServosController{ssc32u}
}