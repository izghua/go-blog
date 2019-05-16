/**
 * Created by GoLand.
 * User: zhu
 * Email: ylsc633@gmail.com
 * Date: 2019-05-16
 * Time: 20:17
 */
package service

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/izghua/go-blog/common"
	"github.com/izghua/go-blog/conf"
	"github.com/izghua/zgh"
	"strconv"
)

func IndexPost(page string,limit string) (indexPostIndex common.IndexPostList,err error)  {
	postKey := conf.PostIndex
	field := ":page:" + page + ":limit:"+limit
	cacheRes,err := conf.CacheClient.HGet(postKey,field).Result()
	if err == redis.Nil {
		// cache key does not exist
		// set data to the cache what use the cache key
		indexPostIndex,err := doCacheIndexPostList(postKey,field,page,limit)
		if err != nil {
			zgh.ZLog().Error("message","service.index.IndexPost","err",err.Error())
			return indexPostIndex,err
		}
		return indexPostIndex,nil
	} else if err != nil {
		zgh.ZLog().Error("message","service.index.IndexPost","err",err.Error())
		return indexPostIndex,err
	}

	if cacheRes == "" {
		// Data is  null
		// set data to the cache what use the cache key
		indexPostIndex,err := doCacheIndexPostList(postKey,field,page,limit)
		if err != nil {
			zgh.ZLog().Error("message","service.index.IndexPost","err",err.Error())
			return indexPostIndex,err
		}
		return indexPostIndex,nil
	}

	err = json.Unmarshal([]byte(cacheRes),&indexPostIndex)
	if err != nil {
		zgh.ZLog().Error("message","service.index.IndexPost","err",err.Error())
		indexPostIndex,err := doCacheIndexPostList(postKey,field,page,limit)
		if err != nil {
			zgh.ZLog().Error("message","service.index.IndexPost","err",err.Error())
			return indexPostIndex,err
		}
		return indexPostIndex,nil
	}
	return
}

func doCacheIndexPostList(cacheKey string,field string,queryPage string,queryLimit string) (res common.IndexPostList,err error) {
	limit,offset := common.Offset(queryPage,queryLimit)
	postList,err := ConsolePostIndex(limit,offset,false)
	if err != nil {
		zgh.ZLog().Error("message","service.index.doCacheIndexPostList","err",err.Error())

		return
	}
	queryPageInt,err := strconv.Atoi(queryPage)
	if err != nil {
		zgh.ZLog().Error("message","service.index.doCacheIndexPostList","err",err.Error())

		return
	}
	postCount,err := ConsolePostCount(limit,offset,false)
	paginate := common.MyPaginate(postCount,limit,queryPageInt)

	res = common.IndexPostList{
		PostListArr: postList,
		Paginate: paginate,
	}

	jsonRes,err := json.Marshal(&res)
	if err != nil {
		zgh.ZLog().Error("message","service.index.doCacheIndexPostList","err",err.Error())
		return
	}
	err = conf.CacheClient.HSet(cacheKey,field,jsonRes).Err()
	if err != nil {
		zgh.ZLog().Error("message","service.index.doCacheIndexPostList","err",err.Error())
		return
	}
	return
}

