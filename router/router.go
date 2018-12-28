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
	r.Use(m.RequestID(m.RequestIDOptions{AllowSetting: false}))
	r.Use(ginutil.Recovery(recoverHandler))
	consolePost := console.NewPost()

	c := r.Group("/console")
	c.Use()
	{
		p := c.Group("/post")
		{
			p.Use(m.Permission("console.post.index")).GET("/",consolePost.Index)
			p.Use(m.Permission("console.post.create")).GET("/create",consolePost.Create)
			p.Use(m.Permission("console.post.store")).POST("/",consolePost.Store)
			p.Use(m.Permission("console.post.edit")).GET("/edit/:id",consolePost.Edit)
			p.Use(m.Permission("console.post.update")).PUT("/:id",consolePost.Update)
			p.Use(m.Permission("console.post.destroy")).DELETE("/:id",consolePost.Destroy)
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

