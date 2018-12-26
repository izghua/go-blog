package entity

type ZPostViews struct {
	Id     int `xorm:"not null pk autoincr INT(10)"`
	PostId int `xorm:"not null comment('文章ID') index INT(11)"`
	Num    int `xorm:"not null comment('阅读次数') INT(11)"`
}
