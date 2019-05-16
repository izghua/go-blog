/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2018-12-20
 * Time: 23:36
 */
package index

import (
	"fmt"
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
	//data := make(map[string]string)
	//data["he"] = "开玩笑"
	//data["ha"] = "大小"

	// post with paginate
	// cate
	// tag
	// link
	// system

	queryPage := c.DefaultQuery("page", "1")
	queryLimit := c.DefaultQuery("limit", conf.DefaultIndexLimit)

	h := gin.H{
		"themeJs": "/static/home/assets/js",
		"themeCss": "/static/home/assets/css",
		"themeImg": "/static/home/assets/img",
		"themeHLight": "/static/home/assets/highlightjs",
		"themeFancyboxCss": "/static/home/assets/fancybox",
		"themeFancyboxJs": "/static/home/assets/fancybox",
	}
	postData,err := service.IndexPost(queryPage,queryLimit)
	if err != nil {
		zgh.ZLog().Error("message","Index.Index","err",err.Error())
		c.HTML(http.StatusOK, "5xx.tmpl", h)
		return
	}

	for _,v := range postData.PostListArr {
		fmt.Println(v)
	}


	cate,err := service.CateListBySort()
	if err != nil {
		zgh.ZLog().Error("message","Index.Index","err",err.Error())
		c.HTML(http.StatusOK, "5xx.tmpl", h)
		return
	}



	h["cate"] = cate
	h["post"] = postData.PostListArr
	h["paginate"] = postData.Paginate
	c.HTML(http.StatusOK, "master.tmpl", h)
	return
}

func TTest(a int,b int) (bs int,res bool) {
	if a > b {
		b++
		return b,true
	} else {
		return b,false
	}

}