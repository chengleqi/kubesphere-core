package v1alpha

import (
	"github.com/chengleqi/kubesphere-core/pkg/apis/v1alpha/ping"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	test := router.Group("/test")
	{
		test.GET("/ping", ping.PingHandler)
	}
	return router
}
