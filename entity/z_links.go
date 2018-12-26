package entity

import (
	"time"
)

type ZLinks struct {
	Id        int       `xorm:"not null pk autoincr INT(10)"`
	Name      string    `xorm:"not null comment('友链名') VARCHAR(255)"`
	Link      string    `xorm:"not null comment('友链链接') VARCHAR(255)"`
	Order     int       `xorm:"not null comment('排序') INT(11)"`
	CreatedAt time.Time `xorm:"created not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	UpdatedAt time.Time `xorm:"updated not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
