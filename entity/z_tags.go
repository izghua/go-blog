package entity

import (
	"time"
)

type ZTags struct {
	Id          int       `json:"id,omitempty" xorm:"not null pk autoincr INT(10)"`
	Name        string    `json:"name,omitempty" xorm:"not null comment('标签名') VARCHAR(255)"`
	DisplayName string    `json:"displayName,omitempty" xorm:"not null comment('标签别名') index VARCHAR(255)"`
	SeoDesc     string    `json:"seoDesc,omitempty" xorm:"comment('seo描述') VARCHAR(255)"`
	Num         int       `json:"num,omitempty" xorm:"not null default 0 comment('被使用次数') INT(11)"`
	CreatedAt   time.Time `json:"createdAt,omitempty" xorm:"created not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	UpdatedAt   time.Time `json:"updatedAt,omitempty" xorm:"updated not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
