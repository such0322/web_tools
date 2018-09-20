package tool

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"time"

	"github.com/go-macaron/session"
)

type User struct {
	ID        int
	Account   string
	Password  string
	Name      string
	Status    int
	CreatedAt int64
}

func NewUser() *User {
	return &User{}
}

func GetUserById(id int) *User {
	var user User
	DB.First(&user, id)
	return &user
}

func GetTestUser() *User {
	t := time.Now()
	return &User{
		ID:        1,
		Account:   "moz1",
		Password:  "",
		Name:      "moz_name",
		Status:    1,
		CreatedAt: t.Unix(),
	}
}

func UserSignin(sess session.Store) *User {
	uid, ok := sess.Get("uid").(int)
	if !ok {
		return nil
	}
	user := GetUserById(uid)
	return user
}

func UserLogin(account, passwd string) (*User, error) {
	var user User
	err := DB.First(&user, "account = ?", account).Error
	if err != nil {
		return nil, err
	}
	md5ctx := md5.New()
	md5ctx.Write([]byte(passwd))
	md5pwd := hex.EncodeToString(md5ctx.Sum(nil))
	if md5pwd != user.Password {
		return nil, errors.New("密码错误，请重试")
	}
	return &user, nil
}

func getUserByAccount(account string) (*User, error) {
	return nil, nil
}
