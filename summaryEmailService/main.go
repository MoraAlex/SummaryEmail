package main

import (
	"fmt"

	"context"

	proto "github.com/MoraAlex/SummaryEmail/summaryEmailService/proto"

	"go-micro.dev/v4"
)

/*

Example usage of top level service initialisation

*/

type SummaryAccount struct{}

func (g *SummaryAccount) GetSummaryAccount(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	rsp.Msg = "Hello " + req.Name
	return nil
}

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("greeter"),
		micro.Version("latest"),
	)

	service.Init()

	// By default we'll run the server unless the flags catch us

	// Setup the server

	// Register handler
	proto.RegisterSummaryAccountHandler(service.Server(), new(SummaryAccount))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
