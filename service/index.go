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
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/izghua/go-blog/common"
	"github.com/izghua/go-blog/conf"
	"github.com/izghua/go-blog/entity"
	"github.com/izghua/zgh"
	"html/template"
	"strconv"
)

type IndexType string

const  (
	IndexTypeOne IndexType = "tag"
	IndexTypeTwo IndexType = "cate"
	IndexTypeThree IndexType = "default"
)


func CommonData() (h gin.H,system *entity.ZSystems,catess []common.IndexCategory,tags []entity.ZTags,links []entity.ZLinks,err error) {
	h = gin.H{
		"themeJs": "/static/home/assets/js",
		"themeCss": "/static/home/assets/css",
		"themeImg": "/static/home/assets/img",
		"themeHLight": "/static/home/assets/highlightjs",
		"themeFancyboxCss": "/static/home/assets/fancybox",
		"themeFancyboxJs": "/static/home/assets/fancybox",
	}
	cates,err := CateListBySort()
	if err != nil {
		zgh.ZLog().Error("message","service.Index.CommonData","err",err.Error())
		return
	}

	for _,v := range cates {
		c := common.IndexCategory{
			Cates: v.Cates,
			Html: template.HTML(v.Html),
		}
		catess = append(catess,c)
	}

	tags,err = AllTags()
	if err != nil {
		zgh.ZLog().Error("message","service.Index.CommonData","err",err.Error())
		return
	}

	links,err = AllLink()
	if err != nil {
		zgh.ZLog().Error("message","service.Index.CommonData","err",err.Error())
		return
	}

	system,err = IndexSystem()
	if err != nil {
		zgh.ZLog().Error("message","service.Index.CommonData","err",err.Error())
		return
	}
	return
}


func IndexPost(page string,limit string,indexType IndexType,name string) (indexPostIndex common.IndexPostList,err error)  {
	var postKey string
	switch indexType {
	case IndexTypeOne:
		postKey = conf.TagPostIndexKey
	case IndexTypeTwo:
		postKey = conf.CatePostIndexKey
	case IndexTypeThree:
		postKey = conf.PostIndexKey
	default:
		postKey = conf.PostIndexKey
	}

	field := ":page:" + page + ":limit:"+limit
	cacheRes,err := conf.CacheClient.HGet(postKey,field).Result()
	if err == redis.Nil {
		// cache key does not exist
		// set data to the cache what use the cache key
		indexPostIndex,err := doCacheIndexPostList(postKey,field,indexType,name,page,limit)
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
		indexPostIndex,err := doCacheIndexPostList(postKey,field,indexType,name,page,limit)
		if err != nil {
			zgh.ZLog().Error("message","service.index.IndexPost","err",err.Error())
			return indexPostIndex,err
		}
		return indexPostIndex,nil
	}

	err = json.Unmarshal([]byte(cacheRes),&indexPostIndex)
	if err != nil {
		zgh.ZLog().Error("message","service.index.IndexPost","err",err.Error())
		indexPostIndex,err := doCacheIndexPostList(postKey,field,indexType,name,page,limit)
		if err != nil {
			zgh.ZLog().Error("message","service.index.IndexPost","err",err.Error())
			return indexPostIndex,err
		}
		return indexPostIndex,nil
	}
	return
}


func doCacheIndexPostList(cacheKey string,field string,indexType IndexType,name string,queryPage string,queryLimit string) (res common.IndexPostList,err error) {
	limit,offset := common.Offset(queryPage,queryLimit)
	queryPageInt,err := strconv.Atoi(queryPage)
	if err != nil {
		zgh.ZLog().Error("message","service.index.doCacheIndexPostList","err",err.Error())
		return
	}
	var postList []*common.ConsolePostList
	var postCount int64
	switch indexType {
	case IndexTypeOne:
		tag := new(entity.ZTags)
		_,err = conf.SqlServer.Where("Name = ?",name).Get(tag)
		if err != nil {
			zgh.ZLog().Error("message","service.index.doCacheIndexPostList","err",err.Error())
			return
		}
		postList,err = PostTagList(tag.Id,limit,offset)
		if err != nil {
			zgh.ZLog().Error("message","service.index.doCacheIndexPostList","err",err.Error())
			return
		}
		postCount,err = PostTagListCount(tag.Id,limit,offset)
		if err != nil {
			zgh.ZLog().Error("message","service.index.doCacheIndexPostList","err",err.Error())
			return
		}
	case IndexTypeTwo:
		cate := new(entity.ZCategories)
		_,err = conf.SqlServer.Where("Name = ?",name).Get(cate)
		if err != nil {
			zgh.ZLog().Error("message","service.index.doCacheIndexPostList","err",err.Error())
			return
		}
		postList,err = PostCateList(cate.Id,limit,offset)
		if err != nil {
			zgh.ZLog().Error("message","service.index.doCacheIndexPostList","err",err.Error())
			return
		}
		postCount,err = PostCateListCount(cate.Id,limit,offset)
		if err != nil {
			zgh.ZLog().Error("message","service.index.doCacheIndexPostList","err",err.Error())
			return
		}
	case IndexTypeThree:
		postList,err = ConsolePostIndex(limit,offset,false)
		if err != nil {
			zgh.ZLog().Error("message","service.index.doCacheIndexPostList","err",err.Error())
			return
		}
		postCount,err = ConsolePostCount(limit,offset,false)
		if err != nil {
			zgh.ZLog().Error("message","service.index.doCacheIndexPostList","err",err.Error())
			return
		}
	default:
		postList,err = ConsolePostIndex(limit,offset,false)
		if err != nil {
			zgh.ZLog().Error("message","service.index.doCacheIndexPostList","err",err.Error())
			return
		}

		postCount,err = ConsolePostCount(limit,offset,false)
		if err != nil {
			zgh.ZLog().Error("message","service.index.doCacheIndexPostList","err",err.Error())
			return
		}
	}

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

