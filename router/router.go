/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2018-12-20
 * Time: 23:24
 */
package router

import (
	"github.com/gin-gonic/gin"
	"github.com/izghua/go-blog/common"
	"github.com/izghua/go-blog/conf"
	m2 "github.com/izghua/go-blog/middleware"
	"github.com/izghua/go-blog/router/auth"
	"github.com/izghua/go-blog/router/console"
	"github.com/izghua/go-blog/router/index"
	"github.com/izghua/go-blog/validate"
	"github.com/izghua/zgh"
	"github.com/izghua/zgh/gin/api"
	m "github.com/izghua/zgh/gin/middleware"
	"github.com/izghua/zgh/gin/util"
	"html/template"
	"net/http"
)

func RoutersInit() *gin.Engine{
	if conf.Env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	r := gin.New()
	r.Use(m.CORS(m.CORSOptions{Origin: ""}))
	r.Use(m.RequestID(m.RequestIDOptions{AllowSetting: true}))
	r.Use(ginutil.Recovery(recoverHandler))
	r.Use(m2.CheckExist())
	r.Static("/static/uploads/images/","./static/uploads/images/")
	consolePost := console.NewPost()
	consoleCate := console.NewCategory()
	consoleTag := console.NewTag()
	postImg := console.NewPostImg()
	trash := console.NewTrash()
	consoleSystem := console.NewHome()
	consoleLink := console.NewLink()
	consoleAuth := auth.NewAuth()
	consoleHome := console.NewStatistics()
	c := r.Group("/console")
	{
		r.Static("/static/console","./static/console")
		r.StaticFile("/backend/","static/console/index.html")
		r.StaticFile("/backend/register","static/console/index.html")
		p := c.Group("/post")
		{
			postV := validate.NewValidate().NewPostV.MyValidate()
			p.GET("/",m2.Permission("console.post.index"),consolePost.Index)
			p.GET("/create",m2.Permission("console.post.create"),consolePost.Create)
			p.POST("/",m2.Permission("console.post.store"),postV,consolePost.Store)
			p.GET("/edit/:id",m2.Permission("console.post.edit"),consolePost.Edit)
			p.PUT("/:id",m2.Permission("console.post.update"),postV,consolePost.Update)
			p.DELETE("/:id",m2.Permission("console.post.destroy"),consolePost.Destroy)
			p.GET("/trash",m2.Permission("console.post.trash"),trash.TrashIndex)
			p.PUT("/:id/trash",m2.Permission("console.post.unTrash"),trash.UnTrash)

			p.POST("/imgUpload",m2.Permission("console.post.imgUpload"),postImg.ImgUpload)
		}
		cate := c.Group("/cate")
		{
			cateV := validate.NewValidate().NewCateV.MyValidate()
			cate.GET("/",m2.Permission("console.cate.index"),consoleCate.Index)
			cate.GET("/edit/:id",m2.Permission("console.cate.edit"),consoleCate.Edit)
			cate.PUT("/:id",m2.Permission("console.cate.update"),cateV,consoleCate.Update)
			cate.POST("/",m2.Permission("console.cate.store"),cateV,consoleCate.Store)
			cate.DELETE("/:id",m2.Permission("console.cate.destroy"),consoleCate.Destroy)
		}
		tag := c.Group("/tag")
		{
			tagV := validate.NewValidate().NewTagV.MyValidate()
			tag.GET("/",m2.Permission("console.tag.index"),consoleTag.Index)
			tag.POST("/",m2.Permission("console.tag.store"),tagV,consoleTag.Store)
			tag.GET("/edit/:id",m2.Permission("console.tag.edit"),consoleTag.Edit)
			tag.PUT("/:id",m2.Permission("console.tag.update"),tagV,consoleTag.Update)
			tag.DELETE("/:id",m2.Permission("console.tag.destroy"),consoleTag.Destroy)
		}
		system := c.Group("/system")
		{
			systemV := validate.NewValidate().NewSystemV.MyValidate()
			system.GET("/",m2.Permission("console.system.index"),consoleSystem.Index)
			system.PUT("/:id",m2.Permission("console.system.update"),systemV,consoleSystem.Update)
		}
		link := c.Group("/link")
		{
			linkV := validate.NewValidate().NewLinkV.MyValidate()
			link.GET("/",m2.Permission("console.link.index"),consoleLink.Index)
			link.POST("/",m2.Permission("console.link.store"),linkV,consoleLink.Store)
			link.GET("/edit/:id",m2.Permission("console.link.edit"),consoleLink.Edit)
			link.PUT("/:id",m2.Permission("console.link.update"),linkV,consoleLink.Update)
			link.DELETE("/:id",m2.Permission("console.link.destroy"),consoleLink.Destroy)
		}
		c.DELETE("/logout",m2.Permission("console.auth.logout"),consoleAuth.Logout)
		c.DELETE("/cache",m2.Permission("console.auth.cache"),consoleAuth.DelCache)
		h := c.Group("/home")
		{
			h.GET("/",m2.Permission("console.home.index"),consoleHome.Index)
		}

		// 不需要登录状态权限

		al := c.Group("/login")
		{
			authLoginV := validate.NewValidate().NewAuthLoginV.MyValidate()
			al.GET("/",m.Permission("console.login.index"),consoleAuth.Login)
			al.POST("/",m.Permission("console.login.store"),authLoginV,consoleAuth.AuthLogin)
		}
		ar := c.Group("/register")
		{
			authRegisterV := validate.NewValidate().NewAuthRegister.MyValidate()
			ar.GET("/",m.Permission("console.register.index"),consoleAuth.Register)
			ar.POST("/",m.Permission("console.register.store"),authRegisterV,consoleAuth.AuthRegister)
		}
	}

	web := index.NewIndex()
	h := r.Group("")
	{
		r.SetFuncMap(template.FuncMap{
			"rem": common.Rem,
			"MDate": common.MDate,
			"MDate2": common.MDate2,
		})
		r.LoadHTMLGlob("template/home/*.tmpl")

		r.Static("/static/home","./static/home")
		h.GET("/",web.Index)
		h.GET("/categories/:name",web.IndexCate)
		h.GET("/tags/:name",web.IndexTag)
		h.GET("/detail/:id",web.Detail)
		h.GET("/archives",web.Archives)
		h.GET("/404",web.NoFound)
	}

	zgh.ZLog().Info("note","router")
	return r
}

func recoverHandler(c *gin.Context, err interface{}) {
	apiG := api.Gin{C: c}
	apiG.Response(http.StatusOK, 400000000, []string{})
	return
}

