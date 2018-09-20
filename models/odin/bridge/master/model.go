package master

import (
	"time"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

type Model struct {
	InsDate time.Time
}
