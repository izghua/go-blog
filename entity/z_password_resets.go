package entity

import (
	"time"
)

type ZPasswordResets struct {
	Email     string    `xorm:"not null index VARCHAR(255)"`
	Token     string    `xorm:"not null VARCHAR(255)"`
	CreatedAt time.Time `xorm:"TIMESTAMP"`
}
