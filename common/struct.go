/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2019-01-12
 * Time: 19:43
 */
package common

import (
	"time"
)

type PostStore struct {
	Title string `json:"title"`
	Category int `json:"category"`
	Tags []int `json:"tags"`
	Summary string `json:"summary"`
	Content string `json:"content"`
}

type CateStore struct {
	Name string `json:"name"`
	DisplayName string `json:"displayName"`
	ParentId int `json:"parentId"`
	SeoDesc string `json:"seoDesc"`
}

type TagStore struct {
	Name string `json:"name"`
	DisplayName string `json:"displayName"`
	SeoDesc string `json:"seoDesc"`
}

type LinkStore struct {
	Name string `json:"name"`
	Link string `json:"link"`
	Order int `json:"order"`
}

type AuthLogin struct {
	Email string `json:"email"`
	Password string `json:"password"`
	Captcha string `json:"captcha"`
	CaptchaKey string `json:"captchaKey"`
}

type AuthRegister struct {
	UserName string `json:"userName"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type ConsolePostList struct {
	Post ConsolePost `json:"post,omitempty"`
	Tags []ConsoleTag `json:"tags,omitempty"`
	Category ConsoleCate `json:"category,omitempty"`
	View ConsoleView `json:"view,omitempty"`
	Author ConsoleUser `json:"author,omitempty"`
}


type ConsolePost struct {
	Id        int 	`json:"id,omitempty"`
	Uid       string `json:"uid,omitempty"`
	UserId    int `json:"userId,omitempty"`
	Title     string `json:"title,omitempty"`
	Summary   string `json:"summary,omitempty"`
	Original  string `json:"original,omitempty"`
	Content   string `json:"content,omitempty"`
	Password  string `json:"password,omitempty"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type ConsoleTag struct {
	Id          int 	`json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	SeoDesc     string `json:"seoDesc,omitempty"`
	Num         int 	`json:"num,omitempty"`
}

type ConsoleCate struct {
	Id          int 	`json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	SeoDesc     string `json:"seoDesc,omitempty"`
}

type ConsoleUser struct {
	Id              int `json:"id,omitempty"`
	Name            string `json:"name,omitempty"`
	Email           string `json:"email,omitempty"`
	Status          int `json:"status,omitempty"`
}

type ConsoleSystem struct {
	Title 			string `json:"title;omitempty"`
	Keywords 		string `json:"keywords;omitempty"`
	Theme        	int `json:"theme;omitempty"`
	Description  	string `json:"description;omitempty"`
	RecordNumber 	string `json:"recordNumber;omitempty"`
}

type ConsoleView struct {
	Num int `json:"num,omitempty"`
}


type Paginate struct {
	Limit int `json:"limit"`
	Count int `json:"count"`
	Total int `json:"total"`
	Last int `json:"last"`
	Current int `json:"current"`
	Next int `json:"next"`
}

