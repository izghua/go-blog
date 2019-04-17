/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2019-04-18
 * Time: 00:05
 */
package validate

import (
	"github.com/gin-gonic/gin"
	"github.com/izghua/go-blog/common"
	"github.com/izghua/zgh/gin/api"
	"net/http"
)

type CateStoreV struct {
}

func (cv *CateStoreV)MyValidate() gin.HandlerFunc {
	return func(c *gin.Context) {
		appG := api.Gin{C: c}
		var json common.CateStore
		//接收各种参数
		if err := c.ShouldBindJSON(&json); err != nil {
			appG.Response(http.StatusOK, 400001000, nil)
			return
		}

		reqValidate := &CateStore{
			Name:json.Name,
			DisplayName:json.DisplayName,
			ParentId:json.ParentId,
			SeoDesc:json.SeoDesc,
		}
		if b := appG.Validate(reqValidate); !b {
			return
		}
		c.Set("json",json)
		c.Next()
	}
}

type CateStore struct {
	Name string `valid:"Required;Max:100"`
	DisplayName string `valid:"Required:Max:100"`
	ParentId int `valid:"Required"`
	SeoDesc string `valid:"Required;Max:250"`
}


func (c *CateStore) Message() map[string]int {
	return map[string]int{
		"Name.Required":402000002,
		"Name.Max":402000006,
		"DisplayName.Required":402000003,
		"DisplayName.Max":402000007,
		"ParentId.Required":402000004,
		"SeoDesc.Required":402000005,
		"SeoDesc.Max":402000008,
	}
}