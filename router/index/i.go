/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2019-05-14
 * Time: 22:01
 */
package index

import "github.com/gin-gonic/gin"

type Home interface {
	Index(*gin.Context)
}