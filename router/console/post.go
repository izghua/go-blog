/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2018-12-27
 * Time: 00:14
 */
package console

import (
	"github.com/gin-gonic/gin"
	"github.com/izghua/go-blog/service"
	"github.com/izghua/zgh"
	"github.com/izghua/zgh/gin/api"
	"net/http"
)

type Post struct {
}

func NewPost() Console {
	return &Post{}
}

func (p *Post)Index(c *gin.Context) {
	data := make(map[string]string)
	data["res"] = "11122"
	data["res2"] = "12341122"
	appG := api.Gin{C: c}
	appG.Response(http.StatusOK,0,data)
	return
}

func (p *Post)Create(c *gin.Context) {
	cates,err := service.CateListBySort()
	appG := api.Gin{C: c}
	if err != nil {
		zgh.ZLog().Error("message","console.Create",err,err.Error())
		appG.Response(http.StatusOK,500000000,nil)
		return
	}
	tags,err := service.AllTags()
	if err != nil {
		zgh.ZLog().Error("message","console.Create",err,err.Error())
		appG.Response(http.StatusOK,500000000,nil)
		return
	}
	data := make(map[string]interface{})
	data["cates"] = cates
	data["tags"] = tags
	appG.Response(http.StatusOK,0,data)
	return
}

func (p *Post)Store(c *gin.Context) {

}

func (p *Post)Edit(c *gin.Context) {

}

func (p *Post)Update(c *gin.Context) {

}

func (p *Post)Destroy(c *gin.Context) {

}