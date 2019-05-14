/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2018-12-27
 * Time: 00:07
 */
package console

import "github.com/gin-gonic/gin"

type Console interface {
	Index(*gin.Context)
	Create(*gin.Context)
	Store(*gin.Context)
	Edit(*gin.Context)
	Update(*gin.Context)
	Destroy(*gin.Context)
}

type Trash interface {
	TrashIndex(*gin.Context)
	UnTrash(*gin.Context)
}

type Img interface {
	ImgUpload(*gin.Context)
}

type System interface {
	Index(*gin.Context)
	Update(*gin.Context)
}

type Statistics interface {
	Index(*gin.Context)
}



