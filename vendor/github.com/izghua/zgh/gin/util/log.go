package ginutil

import (
	"github.com/gin-gonic/gin"
)

func GetCommonRequestData(c *gin.Context) string {
	referer := c.Request.Referer()
	if len(referer) == 0 {
		referer = "nil"
	}
	return "{Worker-Id:" + c.Writer.Header().Get("X-Request-Id") +
		", Method:" + c.Request.Method +
		", Path:" + c.Request.RequestURI +
		", Host:" + c.Request.Host +
		", Referer:" + referer +
		", User-Agent:" + c.Request.Header.Get("User-Agent") +
		", Client-IP:" + c.ClientIP() + "}"
}
