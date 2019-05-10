/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2018-12-20
 * Time: 23:24
 */
package router

import (
	"github.com/gin-gonic/gin"
	"github.com/izghua/go-blog/router/auth"
	"github.com/izghua/go-blog/router/console"
	"github.com/izghua/go-blog/validate"
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
	consoleCate := console.NewCategory()
	consoleTag := console.NewTag()
	postImg := console.NewPostImg()
	trash := console.NewTrash()
	consoleSystem := console.NewHome()
	consoleLink := console.NewLink()
	c := r.Group("/console")
	{
		p := c.Group("/post")
		{
			postV := validate.NewValidate().NewPostV.MyValidate()
			p.GET("/",m.Permission("console.post.index"),consolePost.Index)
			p.GET("/create",m.Permission("console.post.create"),consolePost.Create)
			p.POST("/",m.Permission("console.post.store"),postV,consolePost.Store)
			p.GET("/edit/:id",m.Permission("console.post.edit"),consolePost.Edit)
			p.PUT("/:id",m.Permission("console.post.update"),postV,consolePost.Update)
			p.DELETE("/:id",m.Permission("console.post.destroy"),consolePost.Destroy)
			p.GET("/trash",m.Permission("console.post.trash"),trash.TrashIndex)
			p.PUT("/:id/trash",m.Permission("console.post.unTrash"),trash.UnTrash)

			p.POST("/imgUpload",m.Permission("console.post.imgUpload"),postImg.ImgUpload)
		}
		cate := c.Group("/cate")
		{
			cateV := validate.NewValidate().NewCateV.MyValidate()
			cate.GET("/",m.Permission("console.cate.index"),consoleCate.Index)
			cate.GET("/edit/:id",m.Permission("console.cate.edit"),consoleCate.Edit)
			cate.PUT("/:id",m.Permission("console.cate.update"),cateV,consoleCate.Update)
			cate.POST("/",m.Permission("console.cate.store"),cateV,consoleCate.Store)
			cate.DELETE("/:id",m.Permission("console.cate.destroy"),consoleCate.Destroy)
		}
		tag := c.Group("/tag")
		{
			tagV := validate.NewValidate().NewTagV.MyValidate()
			tag.GET("/",m.Permission("console.tag.index"),consoleTag.Index)
			tag.POST("/",m.Permission("console.tag.store"),tagV,consoleTag.Store)
			tag.GET("/edit/:id",m.Permission("console.tag.edit"),consoleTag.Edit)
			tag.PUT("/:id",m.Permission("console.tag.update"),tagV,consoleTag.Update)
			tag.DELETE("/:id",m.Permission("console.tag.destroy"),consoleTag.Destroy)
		}
		system := c.Group("/system")
		{
			systemV := validate.NewValidate().NewSystemV.MyValidate()
			system.GET("/",m.Permission("console.system.index"),consoleSystem.Index)
			system.PUT("/:id",m.Permission("console.system.update"),systemV,consoleSystem.Update)
		}
		link := c.Group("/link")
		{
			linkV := validate.NewValidate().NewLinkV.MyValidate()
			link.GET("/",m.Permission("console.link.index"),consoleLink.Index)
			link.POST("/",m.Permission("console.link.store"),linkV,consoleLink.Store)
			link.GET("/edit/:id",m.Permission("console.link.edit"),consoleLink.Edit)
			link.PUT("/:id",m.Permission("console.link.update"),linkV,consoleLink.Update)
			link.DELETE("/:id",m.Permission("console.link.destroy"),consoleLink.Destroy)
		}
		consoleAuth := auth.NewAuth()
		au := c.Group("/register")
		{
			authRegisterV := validate.NewValidate().NewAuthLoginV.MyValidate()
			au.GET("/",m.Permission("console.auth.index"),consoleAuth.Login)
			au.POST("/",m.Permission("console.auth.index"),authRegisterV,consoleAuth.AuthLogin)
		}
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

