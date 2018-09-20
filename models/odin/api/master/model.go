package master

import (
	"time"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

type Model struct {
	ID      uint `gorm:"primary_key"`
	InsDate time.Time
}
