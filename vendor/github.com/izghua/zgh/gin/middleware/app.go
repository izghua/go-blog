package ginmiddleware

import (
	"github.com/gin-gonic/gin"
)

func App() gin.HandlerFunc {
	return func(c *gin.Context) {
		//apiG := api.Gin{C: c}
		//clientId := c.Request.Header.Get("x-client-id")
		//if len(clientId) == 0 {
		//	apiG.Response(http.StatusOK, 421000000, nil)
		//	return
		//}
		//clientType := c.Request.Header.Get("x-client-type")
		//if len(clientType) == 0 {
		//	apiG.Response(http.StatusOK, 421000001, nil)
		//	return
		//}
		//clientVersion := c.Request.Header.Get("x-client-version")
		//if len(clientVersion) == 0 {
		//	apiG.Response(http.StatusOK, 421000002, nil)
		//	return
		//}
		c.Next()
	}
}
