/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2019-01-12
 * Time: 23:06
 */
package console

import (
	"github.com/gin-gonic/gin"
	"github.com/izghua/go-blog/common"
	"github.com/izghua/go-blog/service"
	"github.com/izghua/zgh"
	"github.com/izghua/zgh/gin/api"
	"net/http"
	"strconv"
)

type Category struct {
}

func NewCategory() Console {
	return &Category{}
}

func (cate *Category)Index(c *gin.Context) {
	appG := api.Gin{C: c}
	cates,err := service.CateListBySort()
	if err != nil {
		zgh.ZLog().Error("message","console.Cate.Index","err",err.Error())
		appG.Response(http.StatusOK,402000001,nil)
		return
	}
	appG.Response(http.StatusOK,0,cates)
	return
}

func (cate *Category)Create(c *gin.Context) {

}

func (cate *Category)Store(c *gin.Context) {
	appG := api.Gin{C: c}
	requestJson,exists := c.Get("json")
	if !exists {
		zgh.ZLog().Error("message","Cate.Store","error","get request_params from context fail")
		appG.Response(http.StatusOK,400001003,nil)
		return
	}
	var cs common.CateStore
	cs,ok := requestJson.(common.CateStore)
	if !ok {
		zgh.ZLog().Error("message","Cate.Store","error","request_params turn to error")
		appG.Response(http.StatusOK,400001001,nil)
		return
	}

	_,err := service.CateStore(cs)
	if err != nil {
		zgh.ZLog().Error("message","console.Cate.Store","err",err.Error())
		appG.Response(http.StatusOK,402000010,nil)
		return
	}
	appG.Response(http.StatusOK,0,nil)
	return
}

func (cate *Category)Edit(c *gin.Context) {
	cateIdStr := c.Param("id")
	cateIdInt,err := strconv.Atoi(cateIdStr)
	appG := api.Gin{C: c}

	if err != nil {
		zgh.ZLog().Error("message","console.Cate.Edit","err",err.Error())
		appG.Response(http.StatusOK,400001002,nil)
		return
	}
	cateData,err := service.GetCateById(cateIdInt)
	if err != nil {
		zgh.ZLog().Error("message","console.Cate.Edit","err",err.Error())
		appG.Response(http.StatusOK,402000000,nil)
		return
	}
	appG.Response(http.StatusOK,0,cateData)
	return
}

func (cate *Category)Update(c *gin.Context) {
	cateIdStr := c.Param("id")
	cateIdInt,err := strconv.Atoi(cateIdStr)
	appG := api.Gin{C: c}

	if err != nil {
		zgh.ZLog().Error("message","console.Cate.Update","err",err.Error())
		appG.Response(http.StatusOK,400001002,nil)
		return
	}
	requestJson,exists := c.Get("json")
	if !exists {
		zgh.ZLog().Error("message","Cate.Update","error","get request_params from context fail")
		appG.Response(http.StatusOK,400001003,nil)
		return
	}
	var cs common.CateStore
	cs,ok := requestJson.(common.CateStore)
	if !ok {
		zgh.ZLog().Error("message","cate.Update","error","request_params turn to error")
		appG.Response(http.StatusOK,400001001,nil)
		return
	}
	_,err = service.CateUpdate(cateIdInt,cs)
	if err != nil {
		zgh.ZLog().Error("message","cate.Update","error",err.Error())
		appG.Response(http.StatusOK,402000009,nil)
		return
	}
	appG.Response(http.StatusOK,0,nil)
	return
}

func (cate *Category)Destroy(c *gin.Context) {
	cateIdStr := c.Param("id")
	cateIdInt,err := strconv.Atoi(cateIdStr)
	appG := api.Gin{C: c}

	if err != nil {
		zgh.ZLog().Error("message","console.Cate.Destroy","err",err.Error())
		appG.Response(http.StatusOK,400001002,nil)
		return
	}

	_,err = service.GetCateById(cateIdInt)
	if err != nil {
		zgh.ZLog().Error("message","console.Cate.Destroy","err",err.Error())
		appG.Response(http.StatusOK,402000000,nil)
		return
	}

	pd,err := service.GetCateByParentId(cateIdInt)
	if err != nil {
		zgh.ZLog().Error("message","console.Cate.Destroy","err",err.Error())
		appG.Response(http.StatusOK,402000000,nil)
		return
	}
	if pd.Id > 0 {
		zgh.ZLog().Error("message","console.Cate.Destroy",err,"It has children node")
		appG.Response(http.StatusOK,402000011,nil)
		return
	}

	service.DelCateRel(cateIdInt)
	appG.Response(http.StatusOK,0,nil)
	return
}