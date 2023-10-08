package models

type User struct {
	Id       int64  `gorm:"primary_key;column:id;type:bigint(20);not null" json:"id"`
	Username string `gorm:"column:username;type:varchar(255);not null" json:"username"`
	Age      int    `gorm:"column:age;type:int(11);not null" json:"age"`
	Email    string `gorm:"column:email;type:varchar(255);not null" json:"email"`
	AddTime  int    `gorm:"column:add_time;type:int(11);not null" json:"add_time"`
}

// 表示配置操作数据库的名称
func (User) TableName() string {
	return "user"
}
