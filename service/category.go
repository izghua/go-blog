/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2019-01-11
 * Time: 23:24
 */
package service

import (
	"encoding/json"
	"errors"
	"github.com/go-redis/redis"
	"github.com/izghua/go-blog/common"
	"github.com/izghua/go-blog/conf"
	"github.com/izghua/go-blog/entity"
	"github.com/izghua/zgh"
	"time"
)


func GetPostCateByPostId(postId int) ( cates *entity.ZCategories,err error) {
	postCate := new(entity.ZPostCate)
	has,err := conf.SqlServer.Cols("cate_id").Where("post_id = ?",postId).Get(postCate)
	if err != nil {
		zgh.ZLog().Error("message","service.GetPostCateByPostId",err,err.Error())
		return cates,err
	}
	if has {
		cates = new(entity.ZCategories)
		has,err =  conf.SqlServer.Where("id = ?",postCate.CateId).Cols("id","name","display_name","seo_desc").Get(cates)
		if err != nil {
			zgh.ZLog().Error("message","service.GetPostCateByPostId",err,err.Error())
			return cates,err
		}
		if !has {
			zgh.ZLog().Error("message","service.GetPostCateByPostId","err","there has not data")
			return cates,errors.New("can not get the post cate")
		}
	} else {
		zgh.ZLog().Error("message","service.GetPostCateByPostId","err","there has not data")
		return cates,errors.New("can not get the post cate")
	}

	return cates,nil

}


// Get the cate list what by parent sort
func CateListBySort() ([]common.Category, error) {
	cacheKey := "all:cate:sort"
	cacheRes,err := conf.CacheClient.Get(cacheKey).Result()
	if err == redis.Nil {
		// cache key does not exist
		// set data to the cache what use the cache key
		cates,err := doCacheCateList(cacheKey)
		if err != nil {
			zgh.ZLog().Error("message","service.CateListBySort",err,err.Error())
			return nil,err
		}
		return cates,nil
	} else if err != nil {
		zgh.ZLog().Error("message","service.CateListBySort","err",err.Error())
		return nil,err
	}

	if cacheRes == "" {
		// Data is  null
		// set data to the cache what use the cache key
		cates,err := doCacheCateList(cacheKey)
		if err != nil {
			zgh.ZLog().Error("message","service.CateListBySort",err,err.Error())
			return nil,err
		}
		return cates,nil
	}

	var comCate []common.Category
	err = json.Unmarshal([]byte(cacheRes),&comCate)
	if err != nil {
		zgh.ZLog().Error("message","service.CateListBySort",err,err.Error())
		cates,err := doCacheCateList(cacheKey)
		if err != nil {
			zgh.ZLog().Error("message","service.CateListBySort",err,err.Error())
			return nil,err
		}
		return cates,nil
	}
	return comCate,nil
}

// Get the all cate
// then set it to cache
func doCacheCateList(cacheKey string) ([]common.Category,error) {
	allCates,err := allCates()
	if err != nil {
		zgh.ZLog().Error("message","service.CateListBySort",err,err.Error())
		return nil,err
	}
	var cate common.Category
	var cates []common.Category
	for _,v := range allCates {
		cate.Cates = v
		cates = append(cates,cate)
	}
	res := tree(cates,0,0,0)
	jsonRes,err := json.Marshal(&res)
	if err != nil {
		zgh.ZLog().Error("message","service.CateListBySort",err,err.Error())
		return nil,err
	}
	err = conf.CacheClient.Set(cacheKey,jsonRes,conf.DataCacheTimeDuration * time.Hour).Err()
	if err != nil {
		zgh.ZLog().Error("message","service.CateListBySort",err,err.Error())
		return nil,err
	}
	return res,nil
}

// data recursion
func tree(cate []common.Category,parent int,level int,key int) []common.Category {
	html := "-"
	var data []common.Category
	for _,v := range cate {
		var ParentId = v.Cates.ParentId
		var Id = v.Cates.Id
		if ParentId == parent {
			var newHtml string
			if level != 0 {
				newHtml = common.GoRepeat("&nbsp;&nbsp;&nbsp;&nbsp;", level) + "|"
			}
			v.Html = newHtml + common.GoRepeat(html, level)
			data = append(data,v)
			data = merge(data,tree(cate, Id, level+1,key+1))
		}
	}
	return data
}

// merge two arr
func merge(arr1 []common.Category,arr2 []common.Category) []common.Category {
	for _,val := range arr2 {
		arr1 = append(arr1,val)
	}
	return arr1
}

// Get all cate
// if not exists
// create the default one
func allCates() ([]entity.ZCategories,error) {
	cates := make([]entity.ZCategories,0)
	err := conf.SqlServer.Find(&cates)

	if err != nil {
		zgh.ZLog().Info("message","service.AllCates",err,err.Error())
		return cates,err
	}

	if len(cates) == 0 {
		cateCreate := entity.ZCategories{
			Name:"default",
			DisplayName:"默认分类",
			SeoDesc: "默认的分类",
			ParentId: 0,
		}
		affected,err := conf.SqlServer.Insert(&cateCreate)
		if err != nil {
			zgh.ZLog().Info("message","service.AllCates",err,err.Error())
			return cates,err
		}
		if affected < 1 {
			zgh.ZLog().Info("message","service.AllCates",err,"未成功插入数据")
			return cates,errors.New("插入默认分类数据失败")
		}
		err = conf.SqlServer.Find(&cates)

		if err != nil {
			zgh.ZLog().Info("message","service.AllCates",err,err.Error())
			return cates,err
		}
		return cates,nil
	}

	return cates,nil
}

