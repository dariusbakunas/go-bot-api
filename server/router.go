package server

import (
	"github.com/dariusbakunas/go-bot-api/controllers"
	"github.com/dariusbakunas/go-bot-api/models"
	"github.com/gin-gonic/gin"
)

func NewRouter(ssc32u *models.SSC32U) *gin.Engine {
	router := gin.New()

	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one
	router.Use(gin.Recovery())

	servos := controllers.ServosControllerInit(ssc32u)

	servosRoutes := router.Group("/servos")
	{
		servosRoutes.POST("/:id", servos.Turn)
	}

	return router
}