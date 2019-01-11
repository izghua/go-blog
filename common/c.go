/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2019-01-12
 * Time: 00:15
 */
package common

import "github.com/izghua/go-blog/entity"

type Category struct {
	Cates entity.ZCategories `json:"cates"`
	Html string `json:"html"`
}

