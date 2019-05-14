/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2019-05-08
 * Time: 22:35
 */
package service

import (
	"github.com/izghua/go-blog/common"
	"github.com/izghua/go-blog/conf"
	"github.com/izghua/go-blog/entity"
)

func LinkList(offset int, limit int) (links []entity.ZLinks,cnt int64, err error) {
	links = make([]entity.ZLinks,0)
	cnt,err = conf.SqlServer.Asc("order").Limit(limit,offset).FindAndCount(&links)
	return
}

func LinkSore(ls common.LinkStore) (err error) {
	LinkInsert := entity.ZLinks{
		Name: ls.Name,
		Link: ls.Link,
		Order: ls.Order,
	}
	_,err = conf.SqlServer.Insert(&LinkInsert)
	return
}

func LinkDetail(linkId int) (link *entity.ZLinks,err error) {
	link = new(entity.ZLinks)
	_,err = conf.SqlServer.Id(linkId).Get(link)
	return
}

func LinkUpdate(ls common.LinkStore,linkId int) (err error) {
	linkUpdate := entity.ZLinks{
		Link:ls.Link,
		Name:ls.Name,
		Order:ls.Order,
	}
	_,err = conf.SqlServer.Id(linkId).Update(&linkUpdate)
	return
}

func LinkDestroy(linkId int) (err error) {
	link := new(entity.ZLinks)
	_,err = conf.SqlServer.Id(linkId).Delete(link)
	return
}

func LinkCnt() (cnt int64,err error) {
	link := new(entity.ZLinks)
	cnt,err = conf.SqlServer.Count(link)
	return
}