package controllers

import (
	"fmt"
	"github.com/dariusbakunas/go-bot-api/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type ServosController struct{
	ssc32u *models.SSC32U
}

func (s ServosController) Turn(c *gin.Context) {
	var err error

	angle, err := strconv.Atoi(c.Query("angle"))

	if err != nil || angle < 0 || angle > 180 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid angle"})
		c.Abort()
		return
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil || id < 0 || id > 31 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Servo ID must be between 0 and 31"})
		c.Abort()
		return
	}

	// TODO: make this configurable
	pulseRange := 2100.0 - 900
	pulse :=  int(900 + pulseRange/180.0 * float64(angle))

	timeString := c.Query("time")

	var command string
	var time int

	if len(timeString) > 0 {
		time, err = strconv.Atoi(timeString)

		if err != nil || time < 0 || time > 65535 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Time must be between 0 and 65535"})
			c.Abort()
			return
		}

		command = fmt.Sprintf("#%d P%d T%d\r", id, pulse, time)
	} else {
		command = fmt.Sprintf("#%d P%d \r", id, pulse)
	}
	_, err = s.ssc32u.Write(command)

	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		c.Abort()
		return

	}

	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func ServosControllerInit(ssc32u *models.SSC32U) *ServosController {
	return &ServosController{ssc32u}
}