/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2018-12-27
 * Time: 21:53
 */
package ginmiddleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Permission(routerAsName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		//apiG := api.Gin{C: c}
		fmt.Println(routerAsName)
		//if routerAsName == "" {
		//	apiG.Response(http.StatusOK,0,nil)
		//	return
		//}
		c.Next()
	}
}

