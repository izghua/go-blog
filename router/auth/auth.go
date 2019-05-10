/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2019-05-09
 * Time: 21:34
 */
package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/izghua/go-blog/common"
	"github.com/izghua/go-blog/conf"
	"github.com/izghua/go-blog/service"
	"github.com/izghua/zgh"
	"github.com/izghua/zgh/gin/api"
	"github.com/izghua/zgh/jwt"
	"github.com/mojocn/base64Captcha"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	"time"
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

// customizeRdsStore An object implementing Store interface
type customizeRdsStore struct {
	redisClient *redis.Client
}

// customizeRdsStore implementing Set method of  Store interface
func (s *customizeRdsStore) Set(id string, value string) {
	err := s.redisClient.Set(id, value, time.Minute*10).Err()
	if err != nil {
		zgh.ZLog().Error("message","auth.AuthLogin","error",err.Error())
	}
}

// customizeRdsStore implementing Get method of  Store interface
func (s *customizeRdsStore) Get(id string, clear bool) (value string) {
	val, err := s.redisClient.Get(id).Result()
	if err != nil {
		zgh.ZLog().Error("message","auth.AuthLogin","error",err.Error())
		return
	}
	if clear {
		err := s.redisClient.Del(id).Err()
		if err != nil {
			zgh.ZLog().Error("message","auth.AuthLogin","error",err.Error())
			return
		}
	}
	return val
}

func (c *Auth) Register(ctx *gin.Context) {
	appG := api.Gin{C: ctx}
	cnt,err := service.GetUserCnt()
	if err != nil {
		zgh.ZLog().Error("message","auth.Register","error",err.Error())
		appG.Response(http.StatusOK,400001004,nil)
		return
	}
	if cnt >= conf.UserCnt {
		zgh.ZLog().Error("message","auth.Register","error","User cnt beyond expectation")
		appG.Response(http.StatusOK,400001004,nil)
		return
	}
	appG.Response(http.StatusOK,0,nil)
	return
}
func (c *Auth) AuthRegister(ctx *gin.Context) {
	appG := api.Gin{C: ctx}
	requestJson,exists := ctx.Get("json")
	if !exists {
		zgh.ZLog().Error("message","auth.AuthRegister","error","get request_params from context fail")
		appG.Response(http.StatusOK,401000004,nil)
		return
	}
	ar,ok := requestJson.(common.AuthRegister)
	if !ok {
		zgh.ZLog().Error("message","auth.AuthRegister","error","request_params turn to error")
		appG.Response(http.StatusOK,400001001,nil)
		return
	}
	service.UserStore(ar)
	appG.Response(http.StatusOK,0,nil)
	return
}
func (c *Auth) Login(ctx *gin.Context) {
	appG := api.Gin{C: ctx}
	customStore := customizeRdsStore{conf.CacheClient}
	base64Captcha.SetCustomStore(&customStore)
	var configD = base64Captcha.ConfigDigit{
		Height:     80,
		Width:      240,
		MaxSkew:    0.7,
		DotCount:   80,
		CaptchaLen: 5,
	}
	idKeyD, capD := base64Captcha.GenerateCaptcha("", configD)
	base64stringD := base64Captcha.CaptchaWriteToBase64Encoding(capD)
	data := make(map[string]interface{})
	data["key"] = idKeyD
	data["png"] = base64stringD
	appG.Response(http.StatusOK,0,data)
	return
}
func (c *Auth) AuthLogin(ctx *gin.Context) {
	appG := api.Gin{C: ctx}
	requestJson,exists := ctx.Get("json")
	if !exists {
		zgh.ZLog().Error("message","auth.AuthLogin","error","get request_params from context fail")
		appG.Response(http.StatusOK,401000004,nil)
		return
	}
	al,ok := requestJson.(common.AuthLogin)
	if !ok {
		zgh.ZLog().Error("message","auth.AuthLogin","error","request_params turn to error")
		appG.Response(http.StatusOK,400001001,nil)
		return
	}
	verifyResult := base64Captcha.VerifyCaptcha(al.CaptchaKey, al.Captcha)
	if !verifyResult {
		zgh.ZLog().Error("message","auth.AuthLogin","error","captcha is error")
		appG.Response(http.StatusOK,407000008,nil)
		return
	}

	user,err := service.GetUserByEmail(al.Email)
	if err != nil {
		zgh.ZLog().Error("message","auth.AuthLogin","error",err.Error())
		appG.Response(http.StatusOK,407000009,nil)
		return
	}
	if user.Id <= 0 {
		zgh.ZLog().Error("message","auth.AuthLogin","error","Can get user")
		appG.Response(http.StatusOK,407000009,nil)
		return
	}

	password := []byte(al.Password)
	hashedPassword := []byte(user.Password)
	err = bcrypt.CompareHashAndPassword(hashedPassword,password)
	if err != nil {
		zgh.ZLog().Error("message","auth.AuthLogin","error",err.Error())
		appG.Response(http.StatusOK,407000010,nil)
		return
	}

	userIdStr := strconv.Itoa(user.Id)
	token,err := jwt.CreateToken(userIdStr)
	if err != nil {
		zgh.ZLog().Error("message","auth.AuthLogin","error",err.Error())
		appG.Response(http.StatusOK,407000011,nil)
		return
	}
	appG.Response(http.StatusOK,0,token)
	return
}
