/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2018-12-02
 * Time: 01:34
 */
package main

import (
	"github.com/izghua/zghua/conf"
	"github.com/izghua/zghua/router"
)

func main() {
	conf.DefaultInit()
	//csrf
	//建表

	r := router.RoutersInit()
	_ = r.Run(":8081")
}


