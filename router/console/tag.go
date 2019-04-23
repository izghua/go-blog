/**
 * Created by GoLand.
 * User: zhu
 * Email: ylsc633@gmail.com
 * Date: 2019-04-23
 * Time: 19:25
 */
package console

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/izghua/go-blog/common"
	"github.com/izghua/go-blog/service"
	"github.com/izghua/zgh"
	"github.com/izghua/zgh/gin/api"
	"net/http"
)

type Tag struct {
}

func NewTag() Console {
	return &Tag{}
}

func (t *Tag)Index(c *gin.Context) {
	appG := api.Gin{C: c}
	tags,err := service.AllTags()
	if err != nil {
		zgh.ZLog().Error("message","console.Tag.Index","err",err.Error())
		appG.Response(http.StatusOK,402000001,nil)
		return
	}
	appG.Response(http.StatusOK,0,tags)
	return
}

func (t *Tag)Create(c *gin.Context) {

}

func (t *Tag)Store(c *gin.Context) {
	appG := api.Gin{C: c}
	requestJson,exists := c.Get("json")
	if !exists {
		zgh.ZLog().Error("message","Tag.Store","error","get request_params from context fail")
		appG.Response(http.StatusOK,400001003,nil)
		return
	}
	var ts common.TagStore
	fmt.Println(requestJson,"000000")
	ts,ok := requestJson.(common.TagStore)
	if !ok {
		zgh.ZLog().Error("message","Tag.Store","error","request_params turn to error")
		appG.Response(http.StatusOK,400001001,nil)
		return
	}
	err := service.TagStore(ts)
	if err != nil {
		zgh.ZLog().Error("message","console.Cate.Store","err",err.Error())
		appG.Response(http.StatusOK,403000006,nil)
		return
	}
	appG.Response(http.StatusOK,0,nil)
	return
}

func (t *Tag)Edit(c *gin.Context) {

}

func (t *Tag)Update(c *gin.Context) {

}

func (t *Tag)Destroy(c *gin.Context){

}