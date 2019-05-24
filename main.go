/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2018-12-02
 * Time: 01:34
 */
package main

import (
	"fmt"
	"github.com/izghua/go-blog/conf"
	"github.com/izghua/go-blog/router"
	"github.com/izghua/zgh/utils/cron"
	"time"
)

func main() {
	conf.DefaultInit()
	//csrf
	f := func() {
		fmt.Println(time.Now().Format(time.RFC3339))
	}
	spec := "0 * * * * *"
	cron.ZgCron(spec,f)


	//d, err := yaml.Marshal(&c)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}

	//conf.GetAppUrls()
	//fmt.Println(conf.NConf.AppUrl)
	//conf.SetValueToStruct("AppUrl","./imgages222")
	//conf.SetValueToStruct("AppImgUrl","./imgages3333333")
	//conf.SetValueToStruct("aaa","./imgages3333333")
	//fmt.Println(conf.NConf.AppImgUrl,conf.NConf.AppUrl)
	r := router.RoutersInit()
	fmt.Println("开始运行")
	_ = r.Run(":8081")

}


