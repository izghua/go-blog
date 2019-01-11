/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2018-12-20
 * Time: 23:24
 */
package router

import (
	"github.com/gin-gonic/gin"
	"github.com/izghua/go-blog/router/console"
	"github.com/izghua/zgh"
	"github.com/izghua/zgh/gin/api"
	m "github.com/izghua/zgh/gin/middleware"
	"github.com/izghua/zgh/gin/util"
	"net/http"
)

func RoutersInit() *gin.Engine{
	gin.SetMode(gin.DebugMode)
	r := gin.New()
	r.Use(m.CORS(m.CORSOptions{Origin: ""}))
	r.Use(m.RequestID(m.RequestIDOptions{AllowSetting: true}))
	r.Use(ginutil.Recovery(recoverHandler))
	consolePost := console.NewPost()

	c := r.Group("/console")
	{
		p := c.Group("/post")
		{
			p.GET("/",m.Permission("console.post.index"),consolePost.Index)
			p.GET("/create",m.Permission("console.post.create"),consolePost.Create)
			p.POST("/",m.Permission("console.post.store"),consolePost.Store)
			p.PUT("/:id",m.Permission("console.post.update"),consolePost.Update)
			p.DELETE("/:id",m.Permission("console.post.destroy"),consolePost.Destroy)
		}
		//cate := c.Group("/cate")
		//p.Use()
		//{
		//
		//}
	}
	//r.Use(m.RouterAsName("last"))
	//index := r.Group("/index")
	//index.Use()
	//{
	//	index.GET("/index",index2.Index)
	//}
	zgh.ZLog().Info("标记","路由")
	return r
}

func recoverHandler(c *gin.Context, err interface{}) {
	apiG := api.Gin{C: c}
	apiG.Response(http.StatusOK, 400000000, []string{})
	return
}

