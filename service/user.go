/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2019-01-14
 * Time: 22:25
 */
package service

import (
	"github.com/izghua/go-blog/conf"
	"github.com/izghua/go-blog/entity"
	"github.com/izghua/zgh"
)

func GetUserById(userId int) (*entity.ZUsers, error) {
	user := new(entity.ZUsers)
	_,err := conf.SqlServer.Id(userId).Cols("name","email").Get(user)
	if err != nil {
		zgh.ZLog().Error("message","service.GetUserById","error",err.Error())
		return user,err
	}
	return user,nil
}

func UserCnt() (cnt int64,err error) {
	user := new(entity.ZUsers)
	cnt,err = conf.SqlServer.Count(user)
	return
}
