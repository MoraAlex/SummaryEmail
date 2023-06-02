package handler

import (
	se "github.com/MoraAlex/SummaryEmail/summaryEmailService/proto"
	"github.com/gin-gonic/gin"
)

func SummaryAccount(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "call go-micro v3 http server success"})
	client := se.NewSummaryAccountService("contex", c.Accepted)
	// TODO create grpc client here?
}
