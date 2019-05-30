/**
 * Created by GoLand.
 * User: zhu
 * Email: ylsc633@gmail.com
 * Date: 2019-05-17
 * Time: 11:18
 */
package index

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/izghua/zgh"
	"github.com/izghua/zgh/conf"
	"net/http"
	"strconv"
	"time"
)

type ApiController struct {
	C *gin.Context
}



func (a *ApiController) Response(httpCode, errCode int, data gin.H) {
	if data == nil {
		panic("常规信息应该设置")
	}
	msg := conf.GetMsg(errCode)
	beginTime, _ := strconv.ParseInt(a.C.Writer.Header().Get("X-Begin-Time"), 10, 64)

	duration := time.Now().Sub(time.Unix(0, beginTime))
	milliseconds := float64(duration) / float64(time.Millisecond)
	rounded := float64(int(milliseconds*100+.5)) / 100
	roundedStr := fmt.Sprintf("%.3fms", rounded)
	a.C.Writer.Header().Set("X-Response-time", roundedStr)
	requestUrl := a.C.Request.URL.String()
	requestMethod := a.C.Request.Method
	zgh.ZLog().Info("message", "Index Response","Request Url",requestUrl,"Request method",requestMethod, "code", errCode, "errMsg", msg, "took", roundedStr)
	if errCode == 500 {
		a.C.HTML(http.StatusOK,"5xx.tmpl",data)
	} else if errCode == 404 {
		a.C.HTML(http.StatusOK,"4xx.tmpl",data)
	} else if errCode == 0  {
		a.C.HTML(http.StatusOK,"master.tmpl",data)
	} else {
		a.C.HTML(http.StatusOK,"5xx.tmpl",nil)
	}

	a.C.Abort()
	return
}
