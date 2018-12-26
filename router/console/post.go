/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2018-12-27
 * Time: 00:14
 */
package console

import (
	"github.com/gin-gonic/gin"
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

}

func (p *Post)Store(c *gin.Context) {

}

func (p *Post)Edit(c *gin.Context) {

}

func (p *Post)Update(c *gin.Context) {

}

func (p *Post)Destroy(c *gin.Context) {

}