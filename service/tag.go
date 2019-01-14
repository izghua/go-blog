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
	"gitlab.yixinonline.org/pkg/yrdLog"
	"time"
)


func GetPostTagsByPostId(postId int) (tagsArr []int,err error) {
	postTag := new(entity.ZPostTag)
	rows,err := conf.SqlServer.Where("post_id = ?",postId).Cols("tag_id").Rows(postTag)
	if err != nil {
		zgh.ZLog().Error("message","service.GetPostTagsByPostId","error",err.Error())
		return nil,nil
	}
	defer rows.Close()
	for rows.Next() {
		postTag := new(entity.ZPostTag)
		err = rows.Scan(postTag)
		if err != nil {
			yrdLog.GetLogger().Error("Method","service.GetPostTagsByPostId","err",err.Error())
			return nil,err
		}
		tagsArr = append(tagsArr,postTag.TagId)
	}
	return
}

func GetTagsByIds(tagIds []int) ([]*entity.ZTags, error) {
	tags := make([]*entity.ZTags,0)
	err := conf.SqlServer.In("id",tagIds).Cols("id","name","display_name","seo_desc","num").Find(&tags)
	if err != nil {
		yrdLog.GetLogger().Error("Method","service.GetTagsByIds","err",err.Error())
		return nil,err
	}
	return tags,nil
}


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

