/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2019-01-12
 * Time: 21:03
 */
package service

import (
	"github.com/izghua/go-blog/common"
	"github.com/izghua/go-blog/conf"
	"github.com/izghua/go-blog/entity"
	"github.com/izghua/zgh"
	"github.com/microcosm-cc/bluemonday"
	"gopkg.in/russross/blackfriday.v2"
)

type ConsolePostList struct {
	Post *entity.ZPosts `json:"post,omitempty"`
	Tags []*entity.ZTags `json:"tags,omitempty"`
	Category *entity.ZCategories `json:"category,omitempty"`
	View *entity.ZPostViews `json:"view,omitempty"`
	Author *entity.ZUsers `json:"author,omitempty"`
}


func ConsolePostIndex(limit int,offset int) (postListArr []*ConsolePostList,err error) {
	post := new(entity.ZPosts)
	rows,err := conf.SqlServer.Limit(limit,offset).Rows(post)

	if err != nil {
		zgh.ZLog().Error("message","service.PostIndex",err,err.Error())
		return nil,err
	}

	defer rows.Close()
	for rows.Next() {
		//post
		post := new(entity.ZPosts)
		err = rows.Scan(post)
		if err != nil {
			zgh.ZLog().Error("message","service.PostIndex",err,err.Error())
			return nil,err
		}
		//category
		cates,err :=GetPostCateByPostId(post.Id)
		if err != nil {
			zgh.ZLog().Error("message","service.PostIndex",err,err.Error())
			return nil,err
		}

		//tag
		tagIds,err := GetPostTagsByPostId(post.Id)
		if err != nil {
			zgh.ZLog().Error("message","service.PostIndex",err,err.Error())
			return nil,err
		}
		tags,err := GetTagsByIds(tagIds)
		if err != nil {
			zgh.ZLog().Error("message","service.PostIndex",err,err.Error())
			return nil,err
		}

		//view
		view,err := PostView(post.Id)
		if err != nil {
			zgh.ZLog().Error("message","service.PostIndex",err,err.Error())
			return nil,err
		}

		//user
		user,err := GetUserById(post.UserId)
		if err != nil {
			zgh.ZLog().Error("message","service.PostIndex",err,err.Error())
			return nil,err
		}

		postList := ConsolePostList{
			Post: post,
			Category: cates,
			Tags: tags,
			View: view,
			Author: user,
		}
		postListArr = append(postListArr,&postList)
	}

	return postListArr,nil
}

func PostView(postId int) (*entity.ZPostViews,error) {
	postV := new(entity.ZPostViews)
	_,err := conf.SqlServer.Cols("Num").Cols("num").Get(postV)
	if err != nil {
		zgh.ZLog().Error("message","service.PostView",err,err.Error())
	}
	return postV,nil
}


func PostStore(ps common.PostStore) {
	postCreate := &entity.ZPosts{
		Title: ps.Title,
		UserId: 1,
		Summary: ps.Summary,
		Original: ps.Content,
	}

	unsafe := blackfriday.Run([]byte(ps.Content))
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
	postCreate.Content = string(html)


	session := conf.SqlServer.NewSession()
	defer session.Close()
	affected,err := session.Insert(postCreate)
	if err != nil {
		zgh.ZLog().Error("message","service.PostStore",err,err.Error())
		_ = session.Rollback()
		return
	}
	if affected < 1 {
		zgh.ZLog().Error("message","service.PostStore","err","post store no succeed")
		_ = session.Rollback()
		return
	}



	if ps.Category > 0 {
		postCateCreate := entity.ZPostCate{
			PostId: postCreate.Id,
			CateId: ps.Category,
		}
		affected,err := session.Insert(postCateCreate)
		if err != nil {
			zgh.ZLog().Error("message","service.PostStore",err,err.Error())
			_ = session.Rollback()
			return
		}

		if affected < 1 {
			zgh.ZLog().Error("message","service.PostStore","err","post cate store no succeed")
			_ = session.Rollback()
			return
		}
	}

	if len(ps.Tags) > 0 {
		for _,v := range ps.Tags {
			postTagCreate := entity.ZPostTag{
				PostId: postCreate.Id,
				TagId: v,
			}
			affected,err := session.Insert(postTagCreate)
			if err != nil {
				zgh.ZLog().Error("message","service.PostStore post tag insert err",err,err.Error())
				_ = session.Rollback()
				return
			}
			if affected < 1 {
				zgh.ZLog().Error("message","service.PostStore","err","post cate store no succeed")
				_ = session.Rollback()
				return
			}
		}
	}

	postView := entity.ZPostViews{
		PostId: postCreate.Id,
		Num: 1,
	}

	affected,err = session.Insert(postView)
	if err != nil {
		zgh.ZLog().Error("message","service.PostStore",err,err.Error())
		_ = session.Rollback()
		return
	}

	if affected < 1 {
		zgh.ZLog().Error("message","service.PostStore","err","post view store no succeed")
		_ = session.Rollback()
		return
	}

	_ = session.Commit()

	uid,err := conf.ZHashId.Encode([]int{postCreate.Id})
	if err != nil {
		zgh.ZLog().Error("message","service.PostStore create uid error",err,err.Error())
		return
	}

	newPostCreate := entity.ZPosts{
		Uid:uid,
	}
	affected,err = session.Where("id = ?",postCreate.Id).Update(newPostCreate)
	if err != nil {
		zgh.ZLog().Error("message","service.PostStore",err,err.Error())
		return
	}

	if affected < 1 {
		zgh.ZLog().Error("message","service.PostStore","err","post view store no succeed")
		return
	}

	return
}

