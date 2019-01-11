/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2019-01-12
 * Time: 01:57
 */
package service

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/izghua/go-blog/conf"
	"github.com/izghua/go-blog/entity"
	"github.com/izghua/zgh"
	"time"
)

func AllTags() ([]entity.ZTags,error) {
	cacheKey := "all:tag"
	cacheRes,err := conf.CacheClient.Get(cacheKey).Result()
	if err == redis.Nil {
		tags,err := doCacheTagList(cacheKey)
		if err != nil {
			zgh.ZLog().Info("message","service.AllTags",err,err.Error())
			return tags,err
		}
		return tags,nil
	} else if err != nil {
		zgh.ZLog().Info("message","service.AllTags",err,err.Error())
		return nil,err
	}

	var cacheTag []entity.ZTags
	err = json.Unmarshal([]byte(cacheRes),&cacheTag)
	if err != nil {
		zgh.ZLog().Error("message","service.AllTags",err,err.Error())
		tags,err := doCacheTagList(cacheKey)
		if err != nil {
			zgh.ZLog().Error("message","service.AllTags",err,err.Error())
			return nil,err
		}
		return tags,nil
	}
	return cacheTag,nil
}

func doCacheTagList(cacheKey string) ([]entity.ZTags,error) {
	tags,err := tags()
	if err != nil {
		zgh.ZLog().Info("message","service.doCacheTagList",err,err.Error())
		return tags,err
	}
	jsonRes,err := json.Marshal(&tags)
	if err != nil {
		zgh.ZLog().Error("message","service.doCacheTagList",err,err.Error())
		return nil,err
	}
	err = conf.CacheClient.Set(cacheKey,jsonRes,conf.DataCacheTimeDuration * time.Hour).Err()
	if err != nil {
		zgh.ZLog().Error("message","service.doCacheTagList",err,err.Error())
		return nil,err
	}
	return tags,nil
}


func tags() ([]entity.ZTags, error) {
	tags := make([]entity.ZTags,0)
	err := conf.SqlServer.Find(&tags)
	if err != nil {
		zgh.ZLog().Info("message","service.Tags",err,err.Error())
		return tags,err
	}

	return tags,nil
}

