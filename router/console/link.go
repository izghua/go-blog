/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2019-05-06
 * Time: 23:33
 */
package console

import "github.com/gin-gonic/gin"

type Link struct {
}

func NewLink() Console {
	return &Link{}
}

func (l *Link) Index(c *gin.Context) {
}
func (l *Link) Create(c *gin.Context) {
}
func (l *Link) Store(c *gin.Context) {
}
func (l *Link) Edit(c *gin.Context) {
}
func (l *Link) Update(c *gin.Context) {
}
func (l *Link) Destroy(c *gin.Context) {
}


