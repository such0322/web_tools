package user

import "time"

const UserMuidPrefix = "enish.jp:"

type User struct {
	Uid             int
	Cid             int
	Muid            string
	Status          string `gorm:"type:enum('tmp','open','suspend')"`
	IsGm            int
	IsGhost         int
	Birthday        time.Time
	TotalLoginNum   int
	LastPaymentDate time.Time
	TotalPayment    int
	LastAccess_date time.Time
	invite_code     string
	session_id      string
	ins_date        time.Time
	resume_date     time.Time
	is_payed        int
}

func LoadUserByGuid(guid string) *User {
	muid := UserMuidPrefix + guid
	var user User
	DB.Where("muid = ?", muid).First(&user)
	return &user
}

func LoadUserByCid(cid int) *User {
	var user User
	DB.Where("cid = ?", cid).First(&user)
	return &user
}

func GetUsersByGuids(guids []string) []User {
	var muids []string
	for _, vo := range guids {
		muids = append(muids, UserMuidPrefix+vo)
	}
	var users []User
	DB.Where("muid in (?)", muids).Find(&users)
	return users
}
