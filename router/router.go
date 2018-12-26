/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2018-12-20
 * Time: 23:24
 */
package router

import (
	"github.com/gin-gonic/gin"
	"github.com/izghua/zgh/gin/api"
	"github.com/izghua/zgh/gin/middleware"
	"github.com/izghua/zgh/gin/util"
	index2 "github.com/izghua/zghua/router/index"
	"net/http"
)

func RoutersInit() *gin.Engine{
	gin.SetMode(gin.DebugMode)
	r := gin.New()
	r.Use(ginmiddleware.CORS(ginmiddleware.CORSOptions{Origin: ""}))
	r.Use(ginmiddleware.RequestID(ginmiddleware.RequestIDOptions{AllowSetting: false}))
	r.Use(ginutil.Recovery(recoverHandler))

	console := r.Group("console/")
	console.Use()
	{
		console.GET("/")
	}
	index := r.Group("/index")
	index.Use()
	{
		index.GET("/index",index2.Index)
	}

	return r
}

func recoverHandler(c *gin.Context, err interface{}) {
	apiG := api.Gin{C: c}
	apiG.Response(http.StatusOK, 400000000, []string{})
	return
}

