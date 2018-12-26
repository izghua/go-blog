/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2018-12-02
 * Time: 01:34
 */
package main

import (
	"github.com/izghua/go-blog/conf"
	"github.com/izghua/go-blog/router"
)

func main() {
	conf.DefaultInit()
	//csrf

	r := router.RoutersInit()
	_ = r.Run(":8081")
}


