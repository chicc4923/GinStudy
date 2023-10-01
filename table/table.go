package table

import (
	"errors"
	"gorm.io/gorm"
)

type UserInfo struct { // 数据库的表名为 user_infos
	ID    int
	Name  string `gorm:"default:'max'"`
	Hobby string
}

func (u *UserInfo) BeforeUpdate(tx *gorm.DB) (err error) {
	if u.ID == 1 {
		return errors.New("ID = 1  的字段不能被修改")
	}
	return
}
