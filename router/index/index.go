/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2018-12-20
 * Time: 23:36
 */
package index

import (
	"github.com/gin-gonic/gin"
	"github.com/izghua/go-blog/conf"
	"github.com/izghua/go-blog/service"
	"github.com/izghua/zgh"
	"net/http"
)

type Web struct {
}

func NewIndex() Home {
	return &Web{}
}

func (w *Web)Index(c *gin.Context) {
	// post with paginate
	// cate
	// tag
	// link
	// system

	queryPage := c.DefaultQuery("page", "1")
	queryLimit := c.DefaultQuery("limit", conf.DefaultIndexLimit)

	h,system,catess,tags,links,err := service.CommonData()
	if err != nil {
		zgh.ZLog().Error("message","Index.Index","err",err.Error())
		c.HTML(http.StatusOK, "5xx.tmpl", h)
		return
	}

	postData,err := service.IndexPost(queryPage,queryLimit,"default","")
	if err != nil {
		zgh.ZLog().Error("message","Index.Index","err",err.Error())
		c.HTML(http.StatusOK, "5xx.tmpl", h)
		return
	}

	h["cates"] = catess
	h["system"] = system
	h["links"] = links
	h["tags"] = tags
	h["post"] = postData.PostListArr
	h["paginate"] = postData.Paginate
	c.HTML(http.StatusOK, "master.tmpl", h)
	return
}

func (w *Web)IndexTag(c *gin.Context) {
	queryPage := c.DefaultQuery("page", "1")
	queryLimit := c.DefaultQuery("limit", conf.DefaultIndexLimit)
	name := c.Param("name")
	h,system,cates,tags,links,err := service.CommonData()
	if err != nil {
		zgh.ZLog().Error("message","Index.Index","err",err.Error())
		c.HTML(http.StatusOK, "5xx.tmpl", h)
		return
	}

	postData,err := service.IndexPost(queryPage,queryLimit,"tag",name)
	if err != nil {
		zgh.ZLog().Error("message","Index.Index","err",err.Error())
		c.HTML(http.StatusOK, "5xx.tmpl", h)
		return
	}

	h["cates"] = cates
	h["system"] = system
	h["links"] = links
	h["tags"] = tags
	h["post"] = postData.PostListArr
	h["paginate"] = postData.Paginate
	//funcMap := template.FuncMap{"rem": common.Rem}
	//t := template.New("tags").Funcs(funcMap)
	//t = template.Must(t.ParseFiles("template/home/tags.tmpl"))
	//t.ExecuteTemplate(w, "layout", time.Now())

	c.HTML(http.StatusOK, "tags.tmpl", h)
	return
}

func (w *Web)IndexCate(c *gin.Context)  {
	queryPage := c.DefaultQuery("page", "1")
	queryLimit := c.DefaultQuery("limit", conf.DefaultIndexLimit)
	name := c.Param("name")

	h,system,cates,tags,links,err := service.CommonData()
	if err != nil {
		zgh.ZLog().Error("message","Index.IndexCate","err",err.Error())
		c.HTML(http.StatusOK, "5xx.tmpl", h)
		return
	}

	postData,err := service.IndexPost(queryPage,queryLimit,"cate",name)
	if err != nil {
		zgh.ZLog().Error("message","Index.IndexCate","err",err.Error())
		c.HTML(http.StatusOK, "5xx.tmpl", h)
		return
	}

	h["cates"] = cates
	h["system"] = system
	h["links"] = links
	h["tags"] = tags
	h["post"] = postData.PostListArr
	h["paginate"] = postData.Paginate
	c.HTML(http.StatusOK, "categories.tmpl", h)
	return

}

func (w *Web)Archives(c *gin.Context) {

}

