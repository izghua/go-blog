package entity

import (
	"time"
)

type ZCategories struct {
	Id          int       `xorm:"not null pk autoincr INT(10)"`
	Name        string    `xorm:"not null comment('分类名') VARCHAR(255)"`
	DisplayName string    `xorm:"not null comment('分类别名') index VARCHAR(255)"`
	SeoDesc     string    `xorm:"comment('seo描述') VARCHAR(255)"`
	ParentId    int       `xorm:"not null default 0 comment('父类ID') index INT(11)"`
	CreatedAt   time.Time `xorm:"created not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	UpdatedAt   time.Time `xorm:"updated not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
