package main

// https://github.com/go-swagger/go-swagger/issues/962

import (
	"log"

	"github.com/gusibi/nCov/api"

	"github.com/aws/aws-lambda-go/events"
	scf "github.com/tencentyun/scf-go-lib/cloudevents/scf"
	"github.com/tencentyun/scf-go-lib/cloudfunction"

	"github.com/gusibi/nCov/cmd/api/scf/httpadapter"
)

var httpAdapter *httpadapter.HandlerAdapter

func init() {
	log.Println("start server...")
	router := api.GetRouters()

	httpAdapter = httpadapter.New(router)
	log.Println("adapter: ", httpAdapter)
}

// Handler go swagger aws lambda handler
func Handler(req events.APIGatewayProxyRequest) (scf.APIGatewayProxyResponse, error) {

	return httpAdapter.Proxy(req)
}

func main() {
	cloudfunction.Start(Handler)
}
