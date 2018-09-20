package master

import (
	"time"
)

type World struct {
	Model
	Name                        string
	Chapter                     string
	UnlockCondition             string
	ContentsLockUnlockCondition string
	BgImg                       string
	SlotDiff                    int
	StartDate                   time.Time
	EndDate                     time.Time
}

func GetAllWorlds() []World {
	var worlds []World
	DB.Find(&worlds)
	return worlds
}
