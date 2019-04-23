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


func GetCateById(cateId int) (cate *entity.ZCategories,err error) {
	cate = new(entity.ZCategories)
	_,err = conf.SqlServer.Id(cateId).Get(cate)
	return
}

func GetCateByParentId(parentId int) (cate *entity.ZCategories,err error) {
	cate = new(entity.ZCategories)
	_,err = conf.SqlServer.Where("parent_id = ?",parentId).Get(cate)
	return
}

func DelCateRel(cateId int) {
	session := conf.SqlServer.NewSession()
	defer session.Close()
	postCate := new(entity.ZPostCate)
	_,err := session.Where("cate_id = ?",cateId).Delete(postCate)
	if err != nil {
		_ = session.Rollback()
		zgh.ZLog().Error("message","service.DelCateRel","err",err.Error())
		return
	}
	cate := new(entity.ZCategories)
	_,err = session.ID(cateId).Delete(cate)
	if err != nil {
		_ = session.Rollback()
		zgh.ZLog().Error("message","service.DelCateRel","err",err.Error())
		return
	}
	err = session.Commit()
	if err != nil {
		_ = session.Rollback()
		zgh.ZLog().Error("message","service.DelCateRel","err",err.Error())
		return
	}
	conf.CacheClient.Del(conf.CateListKey)
	return
}

func CateStore(cs common.CateStore) (bool,error) {

	if cs.ParentId > 0 {
		cate := new(entity.ZCategories)
		_,err := conf.SqlServer.Id(cs.ParentId).Get(cate)
		if err != nil {
			zgh.ZLog().Error("message","service.CateStore","err",err.Error())
			return false,err
		}
		if cate.Id <= 0 {
			zgh.ZLog().Error("message","service.CateStore",err,"The parent id has not data ")
			return false,errors.New("The parent id has not data ")
		}
	}

	cate := entity.ZCategories{
		Name: cs.Name,
		DisplayName: cs.DisplayName,
		SeoDesc: cs.SeoDesc,
		ParentId: cs.ParentId,
	}
	_,err := conf.SqlServer.Insert(cate)
	if err != nil {
		zgh.ZLog().Error("message","service.CateStore","err",err.Error())
		return false,err
	}
	conf.CacheClient.Del(conf.CateListKey)
	return true,nil
}


func CateUpdate(cateId int,cs common.CateStore) (bool,error) {
	cate := new(entity.ZCategories)
	if cs.ParentId != 0 {
		res,err := conf.SqlServer.Id(cs.ParentId).Get(cate)
		if err != nil {
			zgh.ZLog().Error("message","service.CateUpdate","err",err.Error())
			return false,err
		}
		if !res || cate.Id < 1 {
			zgh.ZLog().Error("message","service.CateUpdate",err,"the parent id is not exists ")
			return false,errors.New("the parent id is not exists ")
		}
		ids := []int{cateId}
		resIds := []int{0}
		_,res2,_ := GetSimilar(ids,resIds,0)
		for _,v := range res2 {
			if v == cs.ParentId {
				return false,errors.New("Can not be you child node ")
			}
		}
	}
	cateUpdate := &entity.ZCategories{
		Name: cs.Name,
		DisplayName: cs.DisplayName,
		SeoDesc: cs.SeoDesc,
		ParentId: cs.ParentId,
	}
	_,err := conf.SqlServer.Id(cateId).Cols("name","display_name","seo_desc","parent_id").Update(cateUpdate)
	if err != nil {
		zgh.ZLog().Error("message","service.CateUpdate","err",err.Error())
		return false,err
	}
	conf.CacheClient.Del(conf.CateListKey)
	return true,nil
}

func GetSimilar(beginId []int,resIds []int,level int) (beginId2 []int,resIds2 []int,level2 int) {
	if len(beginId) != 0 {
		cates := make([]*entity.ZCategories,0)
		err := conf.SqlServer.In("parent_id",beginId).Find(&cates)
		if err != nil {
			zgh.ZLog().Error("message","service.GetSimilar",err,"the parent id data is not exists ")
			return []int{},[]int{},0
		}
		if len(cates) != 0 {
			if level == 0 {
				resIds2 = beginId
			} else {
				resIds2 = resIds
			}
			for _,v := range cates {
				id := v.Id
				beginId2 = append(beginId2,id)
				resIds2 = append(resIds2,id)
			}
			level2 = level + 1
			return GetSimilar(beginId2,resIds2,level2)
		}
		if level == 0 && len(cates) == 0 {
			return beginId,beginId,level
		} else {
			return beginId,resIds,level
		}
	}
	return beginId,resIds,level
}




func GetPostCateByPostId(postId int) ( cates *entity.ZCategories,err error) {
	postCate := new(entity.ZPostCate)
	has,err := conf.SqlServer.Cols("cate_id").Where("post_id = ?",postId).Get(postCate)
	if err != nil {
		zgh.ZLog().Error("message","service.GetPostCateByPostId","err",err.Error())
		return cates,err
	}
	if has {
		cates = new(entity.ZCategories)
		has,err =  conf.SqlServer.Where("id = ?",postCate.CateId).Cols("id","name","display_name","seo_desc").Get(cates)
		if err != nil {
			zgh.ZLog().Error("message","service.GetPostCateByPostId","err",err.Error())
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
	cacheKey := conf.CateListKey
	cacheRes,err := conf.CacheClient.Get(cacheKey).Result()
	if err == redis.Nil {
		// cache key does not exist
		// set data to the cache what use the cache key
		cates,err := doCacheCateList(cacheKey)
		if err != nil {
			zgh.ZLog().Error("message","service.CateListBySort","err",err.Error())
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
			zgh.ZLog().Error("message","service.CateListBySort","err",err.Error())
			return nil,err
		}
		return cates,nil
	}

	var comCate []common.Category
	err = json.Unmarshal([]byte(cacheRes),&comCate)
	if err != nil {
		zgh.ZLog().Error("message","service.CateListBySort","err",err.Error())
		cates,err := doCacheCateList(cacheKey)
		if err != nil {
			zgh.ZLog().Error("message","service.CateListBySort","err",err.Error())
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
		zgh.ZLog().Error("message","service.CateListBySort","err",err.Error())
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
		zgh.ZLog().Error("message","service.CateListBySort","err",err.Error())
		return nil,err
	}
	err = conf.CacheClient.Set(cacheKey,jsonRes,conf.DataCacheTimeDuration * time.Hour).Err()
	if err != nil {
		zgh.ZLog().Error("message","service.CateListBySort","err",err.Error())
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
		zgh.ZLog().Info("message","service.AllCates","err",err.Error())
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
			zgh.ZLog().Info("message","service.AllCates","err",err.Error())
			return cates,err
		}
		if affected < 1 {
			zgh.ZLog().Info("message","service.AllCates",err,"未成功插入数据")
			return cates,errors.New("插入默认分类数据失败")
		}
		err = conf.SqlServer.Find(&cates)

		if err != nil {
			zgh.ZLog().Info("message","service.AllCates","err",err.Error())
			return cates,err
		}

		return cates,nil
	}

	return cates,nil
}

