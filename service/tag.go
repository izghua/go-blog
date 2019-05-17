/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2019-01-12
 * Time: 01:57
 */
package service

import (
	"encoding/json"
	"github.com/go-errors/errors"
	"github.com/go-redis/redis"
	"github.com/izghua/go-blog/common"
	"github.com/izghua/go-blog/conf"
	"github.com/izghua/go-blog/entity"
	"github.com/izghua/zgh"
	"time"
)

func TagStore(ts common.TagStore) (err error)  {
	tag := new(entity.ZTags)
	_,err = conf.SqlServer.Where("name = ?",ts.Name).Get(tag)
	if err != nil {
		zgh.ZLog().Error("message","service.TagStore","error",err.Error())
		return err
	}

	if tag.Id > 0 {
		zgh.ZLog().Error("message","service.TagStore","error","Tag has exists")
		return errors.New("Tag has exists")
	}

	tagInsert := &entity.ZTags{
		Name: ts.Name,
		DisplayName: ts.DisplayName,
		SeoDesc: ts.SeoDesc,
		Num: 0,
	}
	_,err = conf.SqlServer.Insert(tagInsert)
	conf.CacheClient.Del(conf.TagListKey)
	return
}

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
			return nil,err
		}
		tagsArr = append(tagsArr,postTag.TagId)
	}
	return
}

func GetTagById(tagId int) (tag *entity.ZTags,err error) {
	tag = new(entity.ZTags)
	_,err = conf.SqlServer.ID(tagId).Get(tag)
	return
}

func TagUpdate(tagId int,ts common.TagStore) error {
	tagUpdate := &entity.ZTags{
		Name: ts.Name,
		DisplayName: ts.DisplayName,
		SeoDesc: ts.SeoDesc,
	}
	_,err := conf.SqlServer.ID(tagId).Update(tagUpdate)
	return err
}

func GetTagsByIds(tagIds []int) ([]*entity.ZTags, error) {
	tags := make([]*entity.ZTags,0)
	err := conf.SqlServer.In("id",tagIds).Cols("id","name","display_name","seo_desc","num").Find(&tags)
	if err != nil {
		return nil,err
	}
	return tags,nil
}

func TagsIndex(limit int,offset int) (num int64,tags []*entity.ZTags,err error) {
	tags = make([]*entity.ZTags,0)
	num,err = conf.SqlServer.Desc("num").Limit(limit,offset).FindAndCount(&tags)
	return
}

func DelTagRel(tagId int) {
	session := conf.SqlServer.NewSession()
	defer session.Close()
	postTag := new(entity.ZPostTag)
	_,err := session.Where("tag_id = ?",tagId).Delete(postTag)
	if err != nil {
		_ = session.Rollback()
		zgh.ZLog().Error("message","service.DelTagRel","err",err.Error())
		return
	}
	tag := new(entity.ZTags)
	_,err = session.ID(tagId).Delete(tag)
	if err != nil {
		_ = session.Rollback()
		zgh.ZLog().Error("message","service.DelTagRel","err",err.Error())
		return
	}
	err = session.Commit()
	if err != nil {
		_ = session.Rollback()
		zgh.ZLog().Error("message","service.DelTagRel","err",err.Error())
		return
	}
	conf.CacheClient.Del(conf.TagListKey)
	return
}

func AllTags() ([]entity.ZTags,error) {
	cacheKey := conf.TagListKey
	cacheRes,err := conf.CacheClient.Get(cacheKey).Result()
	if err == redis.Nil {
		tags,err := doCacheTagList(cacheKey)
		if err != nil {
			zgh.ZLog().Info("message","service.AllTags","err",err.Error())
			return tags,err
		}
		return tags,nil
	} else if err != nil {
		zgh.ZLog().Info("message","service.AllTags","err",err.Error())
		return nil,err
	}

	var cacheTag []entity.ZTags
	err = json.Unmarshal([]byte(cacheRes),&cacheTag)
	if err != nil {
		zgh.ZLog().Error("message","service.AllTags","err",err.Error())
		tags,err := doCacheTagList(cacheKey)
		if err != nil {
			zgh.ZLog().Error("message","service.AllTags","err",err.Error())
			return nil,err
		}
		return tags,nil
	}
	return cacheTag,nil
}

func doCacheTagList(cacheKey string) ([]entity.ZTags,error) {
	tags,err := tags()
	if err != nil {
		zgh.ZLog().Info("message","service.doCacheTagList","err",err.Error())
		return tags,err
	}
	jsonRes,err := json.Marshal(&tags)
	if err != nil {
		zgh.ZLog().Error("message","service.doCacheTagList","err",err.Error())
		return nil,err
	}
	err = conf.CacheClient.Set(cacheKey,jsonRes,conf.DataCacheTimeDuration * time.Hour).Err()
	if err != nil {
		zgh.ZLog().Error("message","service.doCacheTagList","err",err.Error())
		return nil,err
	}
	return tags,nil
}


func tags() ([]entity.ZTags, error) {
	tags := make([]entity.ZTags,0)
	err := conf.SqlServer.Find(&tags)
	if err != nil {
		zgh.ZLog().Info("message","service.Tags","err",err.Error())
		return tags,err
	}

	return tags,nil
}

func TagCnt() (cnt int64,err error) {
	tag := new(entity.ZTags)
	cnt,err = conf.SqlServer.Count(tag)
	return
}