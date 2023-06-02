package service

import (
	se "github.com/MoraAlex/SummaryEmail/summaryEmailService/proto"
	"github.com/gin-gonic/gin"
)

func SummaryAccount(c *gin.Context) {
	params := map[string][]string(c.Request.URL.Query())
	req := se.AccountData{
		Id: params["id"][0],
	}
	// TODO add controller for api and other one for grpc server.
	// Or use nginx to call grpc server
	se.NewGreeterService()
	se.GreeterService(c)
}
