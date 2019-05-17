/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2019-05-11
 * Time: 00:17
 */
package service

import (
	"fmt"
	"github.com/izghua/go-blog/common"
	"github.com/izghua/go-blog/conf"
	"github.com/izghua/go-blog/entity"
	"github.com/izghua/zgh"
	"golang.org/x/crypto/bcrypt"
)

func GetUserByEmail(email string) (user *entity.ZUsers,err error) {
	user = new(entity.ZUsers)
	_,err = conf.SqlServer.Where("email = ?",email).Get(user)
	return
}

func GetUserCnt() (cnt int64,err error) {
	user := new(entity.ZUsers)
	cnt,err = conf.SqlServer.Count(user)
	return
}

func UserStore(ar common.AuthRegister) (user *entity.ZUsers, err error) {
	password := []byte(ar.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		zgh.ZLog().Error("message","service.UserStore","error",err.Error())
		return
	}
	userInsert := entity.ZUsers{
		Name: ar.UserName,
		Email:ar.Email,
		Password: string(hashedPassword),
		Status: 1,
	}
	_,err = conf.SqlServer.Insert(&userInsert)
	if err != nil {
		zgh.ZLog().Error("message","service.UserStore","error",err.Error())
		return
	}
	fmt.Println(userInsert.Id)
	return
}

func DelAllCache() {
	conf.CacheClient.Del(conf.CateListKey,conf.TagListKey,conf.PostIndexKey,conf.TagPostIndexKey,conf.CatePostIndexKey,conf.LinkIndexKey,conf.SystemIndexKey,conf.PostDetailIndexKey)
}