package entity

type ZMigrations struct {
	Id        int    `xorm:"not null pk autoincr INT(10)"`
	Migration string `xorm:"not null VARCHAR(255)"`
	Batch     int    `xorm:"not null INT(11)"`
}
