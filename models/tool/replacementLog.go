package tool

import (
	"time"
)

const (
	DefaultAppID    = "12odin"
	DefaultOperator = "default"
)

type ReplacementLog struct {
	Id        int
	AppId     string
	DepositId string
	Operator  string
	CreatedAt time.Time
}

func (m *ReplacementLog) Create() error {
	m.AppId = DefaultAppID
	m.CreatedAt = time.Now()
	return DB.Create(m).Error
}

func NewReplacementLog(depositId, operator string) error {
	replog := ReplacementLog{}
	replog.DepositId = depositId
	replog.Operator = operator
	return replog.Create()
}
