/**
 * Created by GoLand.
 * User: zhu
 * Email: ylsc633@gmail.com
 * Date: 2019-05-23
 * Time: 15:35
 */
package conf

import (
	"fmt"
	"reflect"
)

var NConf = &Conf{
	AppUrl: "22222222",
}


func SetValueToStruct(key string,value string) *Conf {
	//p := &Conf{}
	v := reflect.ValueOf(NConf).Elem()

	fmt.Println(v.FieldByName(key).IsValid())
	//v.FieldByName(key).Set(reflect.ValueOf(value))
	//v.FieldByName("AppImgUrl").Set(reflect.ValueOf(appImgUrl))
	return NConf
}



type nc func(*Conf) interface{}

func (conf *Conf)SetAppUrl(appUrl string) nc  {
	return func(conf *Conf) interface{} {
		 a := conf.AppUrl
		 conf.AppUrl = appUrl
		 return a
	}
}

func (conf *Conf)SetAppImgUrl(appImgUrl string) nc  {
	return func(conf *Conf) interface{} {
		a := conf.AppImgUrl
		conf.AppImgUrl = appImgUrl
		return a
	}
}