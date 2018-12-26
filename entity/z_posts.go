package entity

import (
	"time"
)

type ZPosts struct {
	Id        int       `xorm:"not null pk autoincr INT(10)"`
	Uid       string    `xorm:"comment('uid') index VARCHAR(255)"`
	UserId    int       `xorm:"not null comment('用户ID') index INT(11)"`
	Title     string    `xorm:"not null comment('标题') index VARCHAR(255)"`
	Summary   string    `xorm:"not null comment('摘要') CHAR(255)"`
	Original  string    `xorm:"not null comment('原文章内容') TEXT"`
	Content   string    `xorm:"not null comment('文章内容') TEXT"`
	Password  string    `xorm:"not null comment('文章密码') VARCHAR(255)"`
	DeletedAt time.Time `xorm:"TIMESTAMP"`
	CreatedAt time.Time `xorm:"created not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	UpdatedAt time.Time `xorm:"updated not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
