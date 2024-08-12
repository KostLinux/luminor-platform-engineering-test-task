package middleware

import (
	"log"
	"messager-web-app/config"
	"messager-web-app/controller"

	"github.com/gin-gonic/gin"
	nrgin "github.com/newrelic/go-agent/v3/integrations/nrgin"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func Router(cfg *config.Application) *gin.Engine {
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("messager-web-app"),
		newrelic.ConfigLicense(cfg.NRLicense),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)
	if err != nil {
		log.Printf("Failed to create New Relic application: %v", err)
	}

	router := gin.Default()
	router.Use(nrgin.Middleware(app))

	router.POST("/message", controller.SendMessage(cfg))
	router.GET("/ping", controller.Ping)

	return router
}
