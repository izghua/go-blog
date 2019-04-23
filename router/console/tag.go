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
	"github.com/izghua/go-blog/conf"
	"github.com/izghua/go-blog/service"
	"github.com/izghua/zgh"
	"github.com/izghua/zgh/gin/api"
	"net/http"
	"strconv"
)

type Tag struct {
}

func NewTag() Console {
	return &Tag{}
}

func (t *Tag)Index(c *gin.Context) {
	appG := api.Gin{C: c}

	queryPage := c.DefaultQuery("page", "1")
	queryLimit := c.DefaultQuery("limit", conf.DefaultLimit)

	limit,offset := common.Offset(queryPage,queryLimit)
	count,tags,err := service.TagsIndex(limit,offset)
	if err != nil {
		zgh.ZLog().Error("message","console.Tag.Index","err",err.Error())
		appG.Response(http.StatusOK,402000001,nil)
		return
	}
	queryPageInt,err := strconv.Atoi(queryPage)
	if err != nil {
		zgh.ZLog().Error("message","console.Tag.Index","err",err.Error())
		appG.Response(http.StatusOK,500000000,nil)
		return
	}
	data := make(map[string]interface{})
	data["list"] = tags
	data["page"] = common.MyPaginate(count,limit,queryPageInt)

	appG.Response(http.StatusOK,0,data)
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
	tagIdStr := c.Param("id")
	tagIdInt,err := strconv.Atoi(tagIdStr)
	appG := api.Gin{C: c}

	if err != nil {
		zgh.ZLog().Error("message","console.Tag.Edit","err",err.Error())
		appG.Response(http.StatusOK,400001002,nil)
		return
	}
	tagData,err := service.GetTagById(tagIdInt)
	if err != nil {
		zgh.ZLog().Error("message","console.Tag.Edit","err",err.Error())
		appG.Response(http.StatusOK,403000008,nil)
		return
	}
	appG.Response(http.StatusOK,0,tagData)
	return
}

func (t *Tag)Update(c *gin.Context) {
	tagIdStr := c.Param("id")
	tagIdInt,err := strconv.Atoi(tagIdStr)
	appG := api.Gin{C: c}

	if err != nil {
		zgh.ZLog().Error("message","console.Tag.Update","err",err.Error())
		appG.Response(http.StatusOK,400001002,nil)
		return
	}
	requestJson,exists := c.Get("json")
	if !exists {
		zgh.ZLog().Error("message","Tag.Update","error","get request_params from context fail")
		appG.Response(http.StatusOK,400001003,nil)
		return
	}
	ts,ok := requestJson.(common.TagStore)
	if !ok {
		zgh.ZLog().Error("message","Tag.Update","error","request_params turn to error")
		appG.Response(http.StatusOK,400001001,nil)
		return
	}
	err = service.TagUpdate(tagIdInt,ts)
	if err != nil {
		zgh.ZLog().Error("message","Tag.Update","error",err.Error())
		appG.Response(http.StatusOK,403000007,nil)
		return
	}
	appG.Response(http.StatusOK,0,nil)
	return
}

func (t *Tag)Destroy(c *gin.Context){
	tagIdStr := c.Param("id")
	tagIdInt,err := strconv.Atoi(tagIdStr)
	appG := api.Gin{C: c}

	if err != nil {
		zgh.ZLog().Error("message","console.Tag.Destroy","err",err.Error())
		appG.Response(http.StatusOK,400001002,nil)
		return
	}

	_,err = service.GetTagById(tagIdInt)
	if err != nil {
		zgh.ZLog().Error("message","console.Tag.Destroy","err",err.Error())
		appG.Response(http.StatusOK,403000008,nil)
		return
	}
	service.DelTagRel(tagIdInt)
	appG.Response(http.StatusOK,0,nil)
	return
}