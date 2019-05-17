/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2018-12-20
 * Time: 23:36
 */
package index

import (
	"github.com/gin-gonic/gin"
	"github.com/izghua/go-blog/common"
	"github.com/izghua/go-blog/conf"
	"github.com/izghua/go-blog/service"
	"github.com/izghua/zgh"
	"net/http"
)

type Web struct {
	ApiController
}

func NewIndex() Home {
	return &Web{}
}

func (w *Web)Index(c *gin.Context) {
	w.C = c
	queryPage := c.DefaultQuery("page", "1")
	queryLimit := c.DefaultQuery("limit", conf.DefaultIndexLimit)

	h,err := service.CommonData()
	if err != nil {
		zgh.ZLog().Error("message","Index.Index","err",err.Error())
		w.Response(http.StatusOK,408000000,h)
		return
	}

	postData,err := service.IndexPost(queryPage,queryLimit,"default","")
	if err != nil {
		zgh.ZLog().Error("message","Index.Index","err",err.Error())
		w.Response(http.StatusOK,408000001,h)
		return
	}

	h["post"] = postData.PostListArr
	h["paginate"] = postData.Paginate
	w.Response(http.StatusOK,0,h)
	return
}

func (w *Web)IndexTag(c *gin.Context) {
	w.C = c
	queryPage := c.DefaultQuery("page", "1")
	queryLimit := c.DefaultQuery("limit", conf.DefaultIndexLimit)
	name := c.Param("name")
	h,err := service.CommonData()
	if err != nil {
		zgh.ZLog().Error("message","Index.Index","err",err.Error())
		w.Response(http.StatusOK,408000000,h)
		return
	}

	postData,err := service.IndexPost(queryPage,queryLimit,"tag",name)
	if err != nil {
		zgh.ZLog().Error("message","Index.Index","err",err.Error())
		w.Response(http.StatusOK,408000001,h)
		return
	}

	h["post"] = postData.PostListArr
	h["paginate"] = postData.Paginate
	h["tem"] = "tagList"

	c.HTML(http.StatusOK, "master.tmpl", h)
	return
}

func (w *Web)IndexCate(c *gin.Context)  {
	w.C = c
	queryPage := c.DefaultQuery("page", "1")
	queryLimit := c.DefaultQuery("limit", conf.DefaultIndexLimit)
	name := c.Param("name")

	h,err := service.CommonData()
	if err != nil {
		zgh.ZLog().Error("message","Index.IndexCate","err",err.Error())
		w.Response(http.StatusOK,408000000,h)
		return
	}

	postData,err := service.IndexPost(queryPage,queryLimit,"cate",name)
	if err != nil {
		zgh.ZLog().Error("message","Index.IndexCate","err",err.Error())
		w.Response(http.StatusOK,408000001,h)
		return
	}

	h["post"] = postData.PostListArr
	h["paginate"] = postData.Paginate
	h["tem"] = "cateList"
	w.Response(http.StatusOK,0,h)
	return

}

func (w *Web)Detail(c *gin.Context) {
	w.C = c
	postIdStr := c.Param("id")

	h,err := service.CommonData()
	if err != nil {
		zgh.ZLog().Error("message","Index.Detail","err",err.Error())
		w.Response(http.StatusOK,408000000,h)
		return
	}

	postDetail,err :=  service.IndexPostDetail(postIdStr)
	if err != nil {
		zgh.ZLog().Error("message","Index.Detail","err",err.Error())
		w.Response(http.StatusOK,408000002,h)
		return
	}

	go service.PostViewAdd(postIdStr)

	github := common.IndexGithubParam{
		GithubName: conf.GithubName,
		GithubRepo: conf.GithubRepo,
		GithubClientId: conf.GithubClientId,
		GithubClientSecret: conf.GithubClientSecret,
	}

	h["post"] = postDetail
	h["github"] = github
	h["tem"] = "detail"
	w.Response(http.StatusOK,0,h)
	return
}


func (w *Web)Archives(c *gin.Context) {

}

