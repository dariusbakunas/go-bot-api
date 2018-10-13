package server

import (
	"fmt"
	"github.com/dariusbakunas/go-bot-api/models"
)

func Init(port int, ssc32u *models.SSC32U) {
	r := NewRouter(ssc32u)
	r.Run(fmt.Sprintf(":%d", port))
}
