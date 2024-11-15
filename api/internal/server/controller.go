package server

import (
	brokerApiService "dpo-document-api/internal/business/service"
	brokerApiTransport "dpo-document-api/internal/business/transport/http"
	"dpo-document-api/internal/server/deps"
	httpTransport "github.com/mpu-project-dpo/documentsGenerantorAll/pkg/app-kit/http-transport"
	appKit "github.com/mpu-project-dpo/documentsGenerantorAll/pkg/app-kit/server"
	"net/http"
)

func GetController(deps *deps.Providers) appKit.IController {
	service := brokerApiTransport.New(
		brokerApiService.New(
			deps.NatsClient,
		),
	)
	mux := http.NewServeMux()

	mux.Handle("POST /api/v1/document/process", http.HandlerFunc(service.ProcessDocument))

	return httpTransport.New(
		deps.ServerHTTP,
		mux,
		false,
	)
}
