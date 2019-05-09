/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2019-05-09
 * Time: 21:34
 */
package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/izghua/zgh/gin/api"
	"github.com/mojocn/base64Captcha"
	"net/http"
)

type ConsoleAuth interface {
	Register(*gin.Context)
	AuthRegister(*gin.Context)
	Login(*gin.Context)
	AuthLogin(*gin.Context)
}


type Auth struct {
}

func NewAuth() ConsoleAuth {
	return &Auth{}
}

func (c *Auth) Register(ctx *gin.Context) {

}
func (c *Auth) AuthRegister(ctx *gin.Context) {

}
func (c *Auth) Login(ctx *gin.Context) {
	appG := api.Gin{C: ctx}

	var configD = base64Captcha.ConfigDigit{
		Height:     80,
		Width:      240,
		MaxSkew:    0.7,
		DotCount:   80,
		CaptchaLen: 5,
	}
	idKeyD, capD := base64Captcha.GenerateCaptcha("", configD)
	//write to base64 string.
	base64stringD := base64Captcha.CaptchaWriteToBase64Encoding(capD)
	fmt.Println(idKeyD,base64stringD)
	data := make(map[string]interface{})
	data["key"] = idKeyD
	data["png"] = base64stringD
	appG.Response(http.StatusOK,0,data)
	return
}
func (c *Auth) AuthLogin(ctx *gin.Context) {

}
