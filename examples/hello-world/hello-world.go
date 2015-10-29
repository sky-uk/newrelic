package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sky-uk/newrelic-go-agent"
	"golang.org/x/net/context"
	"log"
	"os"
	"time"
)

func main() {

	log.SetOutput(os.Stdout)

	appName := os.Getenv("NEW_RELIC_APP_NAME")
	licenseKey := os.Getenv("NEW_RELIC_LICENSE_KEY")
	if len(appName) == 0 || len(licenseKey) == 0 {
		panic("app name/license key not set")
	}

	newrelic.Init(appName, licenseKey)
	newrelic.RecordMetrics(5 * time.Second) // Record CPU/Memory usage every 5 seconds

	r := gin.Default()
	r.Use(newRelicTraceRequest())
	r.GET("/", helloWorld)

	port := os.Getenv("PORT")
	log.Print("Starting on port " + port)
	r.Run(":" + port)

}

func helloWorld(c *gin.Context) {
	c.String(200, "hello world")
}

func newRelicTraceRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		// c.HandlerName() return's the functions name which appears in the transaction's list
		// You could use the URL if you wanted to - c.Request.URL.Path, or name it per request
		ctx, t := newrelic.TraceRequest(ctx, c.HandlerName(), c.Request)
		defer t.Done()
		c.Next()
	}
}
