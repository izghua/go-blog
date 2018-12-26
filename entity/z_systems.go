package entity

import (
	"time"
)

type ZSystems struct {
	Id           int       `xorm:"not null pk autoincr INT(10)"`
	Theme        int       `xorm:"not null default 0 comment('主题') TINYINT(4)"`
	Title        string    `xorm:"not null comment('网站title') VARCHAR(255)"`
	Keywords     string    `xorm:"not null comment('网站关键字') VARCHAR(255)"`
	Description  string    `xorm:"not null comment('网站描述') VARCHAR(255)"`
	RecordNumber string    `xorm:"not null comment('备案号') VARCHAR(255)"`
	CreatedAt    time.Time `xorm:"created not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	UpdatedAt    time.Time `xorm:"updated not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
