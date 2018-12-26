/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2018-12-20
 * Time: 23:36
 */
package index

import (
	"github.com/gin-gonic/gin"
	"github.com/izghua/zgh/gin/api"
	"net/http"
)

func Index(c *gin.Context) {
	data := make(map[string]string)
	data["he"] = "开玩笑"
	data["ha"] = "大小"

	appG := api.Gin{C: c}
	appG.Response(http.StatusOK,0,data)
	return
}