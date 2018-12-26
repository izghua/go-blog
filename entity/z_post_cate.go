package entity

type ZPostCate struct {
	Id     int `xorm:"not null pk autoincr INT(10)"`
	PostId int `xorm:"not null comment('文章ID') index INT(11)"`
	CateId int `xorm:"not null comment('分类ID') index INT(11)"`
}
