/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2019-05-06
 * Time: 23:17
 */
package console

import "github.com/gin-gonic/gin"

type Home struct {
}

func NewHome() System {
	return &Home{}
}

func (s *Home) Index(c *gin.Context){

}
func (s *Home) Update(c *gin.Context){

}




