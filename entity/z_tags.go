package entity

import (
	"time"
)

type ZTags struct {
	Id          int       `xorm:"not null pk autoincr INT(10)"`
	Name        string    `xorm:"not null comment('标签名') VARCHAR(255)"`
	DisplayName string    `xorm:"not null comment('标签别名') index VARCHAR(255)"`
	SeoDesc     string    `xorm:"comment('seo描述') VARCHAR(255)"`
	Num         int       `xorm:"not null default 0 comment('被使用次数') INT(11)"`
	CreatedAt   time.Time `xorm:"created not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	UpdatedAt   time.Time `xorm:"updated not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
