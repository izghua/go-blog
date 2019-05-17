/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2019-05-14
 * Time: 22:24
 */
package console

import (
	"github.com/gin-gonic/gin"
	"github.com/izghua/go-blog/service"
	"github.com/izghua/zgh"
	"github.com/izghua/zgh/gin/api"
	"net/http"
)

type HomeStatistics struct {
}

func NewStatistics() Statistics {
	return &HomeStatistics{}
}

func (h *HomeStatistics)Index(c *gin.Context)  {
	appG := api.Gin{C: c}
	postCnt,err := service.PostCnt()
	if err != nil {
		zgh.ZLog().Info("message","console.Home.Index","err",err.Error())
		appG.Response(http.StatusOK,400001004,nil)
		return
	}

	cateCnt,err := service.CateCnt()
	if err != nil {
		zgh.ZLog().Info("message","console.Home.Index","err",err.Error())
		appG.Response(http.StatusOK,400001004,nil)
		return
	}

	tagCnt,err := service.TagCnt()
	if err != nil {
		zgh.ZLog().Info("message","console.Home.Index","err",err.Error())
		appG.Response(http.StatusOK,400001004,nil)
		return
	}

	linkCnt,err := service.LinkCnt()
	if err != nil {
		zgh.ZLog().Info("message","console.Home.Index","err",err.Error())
		appG.Response(http.StatusOK,400001004,nil)
		return
	}

	userCnt,err := service.UserCnt()
	if err != nil {
		zgh.ZLog().Info("message","console.Home.Index","err",err.Error())
		appG.Response(http.StatusOK,400001004,nil)
		return
	}

	//{ title: '新增用户', icon: 'md-person-add', count: 803, color: '#2d8cf0' },
	//{ title: '累计点击', icon: 'md-locate', count: 232, color: '#19be6b' },
	//{ title: '新增问答', icon: 'md-help-circle', count: 142, color: '#ff9900' },
	//{ title: '分享统计', icon: 'md-share', count: 657, color: '#ed3f14' },
	//{ title: '新增互动', icon: 'md-chatbubbles', count: 12, color: '#E46CBB' },
	//{ title: '新增页面', icon: 'md-map', count: 14, color: '#9A66E4' }

	var data  []interface{}
	pcnt := Res{
		Title: "文章总数",
		Icon: "ios-book-outline",
		Count: postCnt,
		Color: "#ff9900",
	}
	data = append(data,pcnt)
	ucnt := Res{
		Title: "用户总数",
		Icon: "md-person-add",
		Count: userCnt,
		Color: "#2d8cf0",
	}
	data = append(data,ucnt)
	lcnt := Res{
		Title: "外链总数",
		Icon: "ios-link",
		Count: linkCnt,
		Color: "#E46CBB",
	}
	data = append(data,lcnt)
	ccnt := Res{
		Title: "分类总数",
		Icon: "md-locate",
		Count: cateCnt,
		Color: "#19be6b",
	}
	data = append(data,ccnt)
	tcnt := Res{
		Title: "标签总数",
		Icon: "md-share",
		Count: tagCnt,
		Color: "#39ed14",
	}
	data = append(data,tcnt)
	qcnt := Res{
		Title: "未知BUG",
		Icon: "ios-bug",
		Count: 998,
		Color: "#ed3f14",
	}
	data = append(data,qcnt)
	appG.Response(http.StatusOK,0,data)
	return
}

type Res struct {
	Title string `json:"title"`
	Icon string `json:"icon"`
	Count int64 `json:"count"`
	Color string `json:"color"`
}