/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2018-12-27
 * Time: 00:14
 */
package console

import (
	"github.com/gin-gonic/gin"
	"github.com/izghua/go-blog/common"
	"github.com/izghua/go-blog/conf"
	"github.com/izghua/go-blog/service"
	"github.com/izghua/zgh"
	"github.com/izghua/zgh/gin/api"
	"net/http"
	"path/filepath"
	"strconv"
)

type Post struct {
}

func NewPost() Console {
	return &Post{}
}

func NewPostImg() Img {
	return &Post{}
}

func (p *Post)Index(c *gin.Context) {
	appG := api.Gin{C: c}

	queryPage := c.DefaultQuery("page", "1")
	queryLimit := c.DefaultQuery("limit", conf.DefaultLimit)

	limit,offset := common.Offset(queryPage,queryLimit)
	postList,err := service.ConsolePostIndex(limit,offset)
	if err != nil {
		zgh.ZLog().Error("message","console.Index",err,err.Error())
		appG.Response(http.StatusOK,500000000,nil)
		return
	}
	queryPageInt,err := strconv.Atoi(queryPage)
	if err != nil {
		zgh.ZLog().Error("message","console.Index",err,err.Error())
		appG.Response(http.StatusOK,500000000,nil)
		return
	}
	postCount,err := service.ConsolePostCount(limit,offset)

	data := make(map[string]interface{})
	data["list"] = postList
	data["page"] = common.MyPaginate(postCount,limit,queryPageInt)

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
	data["imgUploadUrl"] = conf.ImgUploadUrl
	appG.Response(http.StatusOK,0,data)
	return
}

func (p *Post)Store(c *gin.Context) {
	appG := api.Gin{C: c}
	requestJson,exists := c.Get("json")
	if !exists {
		zgh.ZLog().Error("message","post.Store","error","get request_params from context fail")
		appG.Response(http.StatusOK,401000004,nil)
		return
	}
	var ps common.PostStore
	ps,ok := requestJson.(common.PostStore)
	if !ok {
		zgh.ZLog().Error("message","post.Store","error","request_params turn to error")
		appG.Response(http.StatusOK,400001001,nil)
		return
	}

	service.PostStore(ps)
	appG.Response(http.StatusOK,0,nil)
	return
}

func (p *Post)Edit(c *gin.Context) {
	postIdStr := c.Param("id")
	postIdInt,err := strconv.Atoi(postIdStr)
	appG := api.Gin{C: c}

	if err != nil {
		zgh.ZLog().Error("message","console.Edit",err,err.Error())
		appG.Response(http.StatusOK,500000000,nil)
		return
	}
	post,err := service.PostDetail(postIdInt)
	if err != nil {
		zgh.ZLog().Error("message","console.Edit",err,err.Error())
		appG.Response(http.StatusOK,500000000,nil)
		return
	}
	postTags,err := service.PostIdTag(postIdInt)
	if err != nil {
		zgh.ZLog().Error("message","console.Edit",err,err.Error())
		appG.Response(http.StatusOK,500000000,nil)
		return
	}
	postCate,err := service.PostCate(postIdInt)
	if err != nil {
		zgh.ZLog().Error("message","console.Edit",err,err.Error())
		appG.Response(http.StatusOK,500000000,nil)
		return
	}
	data := make(map[string]interface{})
	posts := make(map[string]interface{})
	posts["post"] = post
	posts["postCate"] = postCate
	posts["postTag"] = postTags
	data["post"] = posts
	cates,err := service.CateListBySort()
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
	data["cates"] = cates
	data["tags"] = tags
	data["imgUploadUrl"] = conf.ImgUploadUrl
	appG.Response(http.StatusOK,0,data)
	return
}

func (p *Post)Update(c *gin.Context) {
	postIdStr := c.Param("id")
	postIdInt,err := strconv.Atoi(postIdStr)
	appG := api.Gin{C: c}

	if err != nil {
		zgh.ZLog().Error("message","console.Update",err,err.Error())
		appG.Response(http.StatusOK,500000000,nil)
		return
	}

	requestJson,exists := c.Get("json")
	if !exists {
		zgh.ZLog().Error("message","post.Store","error","get request_params from context fail")
		appG.Response(http.StatusOK,401000004,nil)
		return
	}
	var ps common.PostStore
	ps,ok := requestJson.(common.PostStore)
	if !ok {
		zgh.ZLog().Error("message","post.Store","error","request_params turn to error")
		appG.Response(http.StatusOK,400001001,nil)
		return
	}
	service.PostUpdate(postIdInt,ps)
	appG.Response(http.StatusOK,0,nil)
	return
}

func (p *Post)Destroy(c *gin.Context) {

}

func (p *Post)ImgUpload(c *gin.Context) {
	appG := api.Gin{C: c}

	file, err := c.FormFile("file")
	if err != nil {
		zgh.ZLog().Info("message","post.ImgUpload","err",err.Error())
		appG.Response(http.StatusOK,401000004,nil)
		return
	}

	filename := filepath.Base(file.Filename)
	dst := conf.ImgUploadDst + filename
	if err := c.SaveUploadedFile(file, dst); err != nil {
		zgh.ZLog().Info("message","post.ImgUpload","error",err.Error())
		appG.Response(http.StatusOK,401000005,nil)
		return
	}
	data := make(map[string]interface{})
	if conf.QiNiuUploadImg {
		go service.Qiniu(dst,filename)
		data["path"] = conf.QiNiuHostName + filename
	} else {
		data["path"] = conf.AppUrl + filename
	}

	appG.Response(http.StatusOK,0,data)
	return
}