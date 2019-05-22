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
	r := router.RoutersInit()
	fmt.Println("开始运行")
	_ = r.Run(":8081")
}


