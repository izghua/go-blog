package entity

type ZPostTag struct {
	Id     int `xorm:"not null pk autoincr INT(10)"`
	PostId int `xorm:"not null comment('文章ID') index INT(11)"`
	TagId  int `xorm:"not null comment('标签ID') index INT(11)"`
}
