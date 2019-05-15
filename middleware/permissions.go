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
	"github.com/izghua/zgh"
	"github.com/izghua/zgh/gin/api"
	"github.com/izghua/zgh/jwt"
	"net/http"
	"strconv"
)

func Permission(routerAsName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		apiG := api.Gin{C: c}
		fmt.Println(routerAsName,c.Request.Method)
		res :=  common.CheckPermissions(routerAsName,c.Request.Method)
		if !res {
			zgh.ZLog().Error("method","middleware.Permission","info","router permission")
			apiG.Response(http.StatusOK,400001005,nil)
			return
		}

		token := c.GetHeader("x-auth-token")
		if routerAsName == "console.post.imgUpload" {
			token = c.PostForm("upload-token")
		}

		if token == "" {
			zgh.ZLog().Error("method","middleware.Permission","info","token null")
			apiG.Response(http.StatusOK,400001005,nil)
			return
		}

		userId,err := jwt.ParseToken(token)
		if err != nil {
			zgh.ZLog().Error("method","middleware.Permission","info","parse token error")
			apiG.Response(http.StatusOK,400001005,nil)
			return
		}

		userIdInt,err := strconv.Atoi(userId)
		if err != nil {
			zgh.ZLog().Error("method","middleware.Permission","info","strconv token error")
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

