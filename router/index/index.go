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
	"sort"
	"time"
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
	queryLimit := c.DefaultQuery("limit", conf.Cnf.DefaultIndexLimit)

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
	queryLimit := c.DefaultQuery("limit", conf.Cnf.DefaultIndexLimit)
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
	h["tagName"] = name
	h["tem"] = "tagList"

	c.HTML(http.StatusOK, "master.tmpl", h)
	return
}

func (w *Web)IndexCate(c *gin.Context)  {
	w.C = c
	queryPage := c.DefaultQuery("page", "1")
	queryLimit := c.DefaultQuery("limit", conf.Cnf.DefaultIndexLimit)
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
	h["cateName"] = name
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
		GithubName: conf.Cnf.GithubName,
		GithubRepo: conf.Cnf.GithubRepo,
		GithubClientId: conf.Cnf.GithubClientId,
		GithubClientSecret: conf.Cnf.GithubClientSecret,
	}

	h["post"] = postDetail
	h["github"] = github
	h["tem"] = "detail"
	w.Response(http.StatusOK,0,h)
	return
}


func (w *Web)Archives(c *gin.Context) {
	w.C = c
	h,err := service.CommonData()
	if err != nil {
		zgh.ZLog().Error("message","Index.Archives","err",err.Error())
		w.Response(http.StatusOK,408000000,h)
		return
	}

	res,err := service.PostArchives()
	if err != nil {
		zgh.ZLog().Error("message","Index.Archives","err",err.Error())
		w.Response(http.StatusOK,408000006,h)
		return
	}
	loc, _ := time.LoadLocation("Asia/Shanghai")

	var dateIndexs []int
	for k,_ := range res{
		tt, _ := time.ParseInLocation("2006-01-02 15:04:05", k+"-01 00:00:00", loc)
		dateIndexs = append(dateIndexs,int(tt.Unix()))
	}

	sort.Sort(sort.Reverse(sort.IntSlice(dateIndexs)))

	var newData []interface{}
	for _,j := range dateIndexs {
		dds := make(map[string]interface{})
		tm := time.Unix(int64(j), 0)
		dateIndex := tm.Format("2006-01")
		dds["dates"] = dateIndex
		dds["lists"] = res[dateIndex]
		newData = append(newData,dds)
	}


	h["tem"] = "archives"
	h["archives"] = newData
	w.Response(http.StatusOK,0,h)
	return
}

func (w *Web)NoFound(c *gin.Context)  {
	w.C = c
	w.Response(http.StatusOK,404,gin.H{
		"themeJs": "/static/home/assets/js",
		"themeCss": "/static/home/assets/css",
	})
	return
}