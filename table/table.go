package table

type UserInfo struct { // 数据库的表名为 user_infos
	ID    int
	Name  string `gorm:"default:'max'"`
	Hobby string
}
