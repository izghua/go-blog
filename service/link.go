/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2019-05-08
 * Time: 22:35
 */
package service

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/izghua/go-blog/common"
	"github.com/izghua/go-blog/conf"
	"github.com/izghua/go-blog/entity"
	"github.com/izghua/zgh"
	"time"
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

func AllLink() (links []entity.ZLinks,err error) {
	cacheKey := conf.Cnf.LinkIndexKey
	cacheRes,err := conf.CacheClient.Get(cacheKey).Result()
	if err == redis.Nil {
		links,err := doCacheLinkList(cacheKey)
		if err != nil {
			zgh.ZLog().Info("message","service.AllLink","err",err.Error())
			return links,err
		}
		return links,nil
	} else if err != nil {
		zgh.ZLog().Info("message","service.AllLink","err",err.Error())
		return nil,err
	}

	err = json.Unmarshal([]byte(cacheRes),&links)
	if err != nil {
		zgh.ZLog().Error("message","service.AllLink","err",err.Error())
		links,err = doCacheLinkList(cacheKey)
		if err != nil {
			zgh.ZLog().Error("message","service.AllLink","err",err.Error())
			return nil,err
		}
		return links,nil
	}
	return links,nil
}

func doCacheLinkList(cacheKey string) (links []entity.ZLinks,err error) {
	links = make([]entity.ZLinks,0)
	err = conf.SqlServer.Find(&links)
	if err != nil {
		zgh.ZLog().Info("message","service.doCacheLinkList","err",err.Error())
		return links,err
	}
	jsonRes,err := json.Marshal(&links)
	if err != nil {
		zgh.ZLog().Error("message","service.doCacheLinkList","err",err.Error())
		return nil,err
	}
	err = conf.CacheClient.Set(cacheKey,jsonRes,time.Duration(conf.Cnf.DataCacheTimeDuration) * time.Hour).Err()
	if err != nil {
		zgh.ZLog().Error("message","service.doCacheLinkList","err",err.Error())
		return nil,err
	}
	return links,nil
}