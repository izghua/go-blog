/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2019-01-12
 * Time: 00:15
 */
package common

import (
	"github.com/izghua/go-blog/entity"
	"html/template"
)

type Category struct {
	Cates entity.ZCategories `json:"cates"`
	Html string `json:"html"`
}

type IndexCategory struct {
	Cates entity.ZCategories `json:"cates"`
	Html template.HTML `json:"html"`
}
