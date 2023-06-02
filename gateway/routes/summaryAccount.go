package routes

import (
	"github.com/MoraAlex/SummaryEmail/gateway/handler"

	"github.com/gin-gonic/gin"
)

func addSummaryAccountRoutes(rg *gin.RouterGroup) {
	rg.GET("/", handler.SummaryAccount)
}
