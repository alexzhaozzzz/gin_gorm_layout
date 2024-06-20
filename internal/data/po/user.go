// Package po
// Generated by sql2struct-0.0.5 at 2024-06-19 07:59:50
package po

// User Generated by sql2struct-0.0.5 at 2024-06-19 07:59:50
type User struct {
    Id int64 `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
    Guid string `gorm:"column:guid" json:"guid"`
    Username string `gorm:"column:username" json:"username"`
    Password string `gorm:"column:password" json:"password"`
    NickName string `gorm:"column:nick_name" json:"nick_name"`
    MerchantId int64 `gorm:"column:merchant_id" json:"merchant_id"`
    Role string `gorm:"column:role" json:"role"`
    Avatar string `gorm:"column:avatar" json:"avatar"`
    Email string `gorm:"column:email" json:"email"`
    Phone string `gorm:"column:phone" json:"phone"`
    IsMaster int32 `gorm:"column:is_master" json:"is_master"`
    CreateTime int32 `gorm:"column:create_time" json:"create_time"`
    UpdateTime int32 `gorm:"column:update_time" json:"update_time"`
    LastLoginTime int32 `gorm:"column:last_login_time" json:"last_login_time"`
}

func (u *User) TableName() string {
    return "user"
}