/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2019-05-13
 * Time: 22:36
 */
package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/izghua/go-blog/common"
	"github.com/izghua/zgh/gin/api"
	"github.com/izghua/zgh/jwt"
	"net/http"
	"strconv"
)

func Permission(routerAsName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("卧槽...")
		apiG := api.Gin{C: c}
		fmt.Println(routerAsName,c.Request.Method)
		res :=  common.CheckPermissions(routerAsName,c.Request.Method)
		if !res {
			apiG.Response(http.StatusOK,400001005,nil)
			return
		}

		token := c.GetHeader("x-auth-token")
		if token == "" {
			apiG.Response(http.StatusOK,400001005,nil)
			return
		}
		userId,err := jwt.ParseToken(token)
		if err != nil {
			apiG.Response(http.StatusOK,400001005,nil)
			return
		}

		userIdInt,err := strconv.Atoi(userId)
		if err != nil {
			apiG.Response(http.StatusOK,400001005,nil)
			return
		}
		c.Set("userId",userIdInt)
		//if routerAsName == "" {
		//	apiG.Response(http.StatusOK,0,nil)
		//	return
		//}
		c.Next()
	}
}

