/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2018-12-16
 * Time: 00:28
 */
package ginmiddleware

import (
	"github.com/gin-gonic/gin"
	"github.com/izghua/zgh"
	"github.com/izghua/zgh/gin/api"
	"github.com/izghua/zgh/jwt"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiG := api.Gin{C: c}

		token := c.Request.Header.Get("x-auth-token")
		if token == "" {
			zgh.ZLog().Error("method","zgh.ginmiddleware.auth","error","token is null")
			apiG.Response(http.StatusOK,400000001,nil)
			return
		}

		userId,err := jwt.ParseToken(token)
		if err != nil {
			zgh.ZLog().Error("method","zgh.ginmiddleware.auth","error",err.Error())
			apiG.Response(http.StatusOK,400000001,nil)
			return
		}

		c.Set("userId",userId)
		c.Next()
	}
}