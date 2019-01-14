package entity

import (
	"time"
)

type ZUsers struct {
	Id              int       `xorm:"not null pk autoincr INT(10)"`
	Name            string    `xorm:"not null comment('用户名') VARCHAR(255)"`
	Email           string    `xorm:"not null comment('邮箱') index unique VARCHAR(255)"`
	Status          int       `xorm:"not null default 0 comment('用户状态 0创建,1正常') TINYINT(4)"`
	EmailVerifiedAt time.Time `xorm:"TIMESTAMP"`
	Password        string    `xorm:"not null comment('密码') VARCHAR(255)"`
	RememberToken   string    `xorm:"VARCHAR(100)"`
	CreatedAt       time.Time `xorm:"created not null default 'CURRENT_TIMESTAMP' TIMESTAMP";json:"created_at,omitempty"`
	UpdatedAt       time.Time `xorm:"updated not null default 'CURRENT_TIMESTAMP' TIMESTAMP";json:"updated_at,omitempty"`
}
