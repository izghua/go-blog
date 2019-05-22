/**
 * Created by GoLand.
 * User: zhu
 * Email: ylsc633@gmail.com
 * Date: 2019-05-22
 * Time: 11:32
 */
package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func CheckExist() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		s := strings.Contains(path,"/backend/")
		c.Next()
		status := c.Writer.Status()
		if status == 404 {
			if s {
				c.Redirect(http.StatusMovedPermanently,"/backend/")
			} else {
				c.Redirect(http.StatusMovedPermanently,"/404")
			}
		}
	}
}
