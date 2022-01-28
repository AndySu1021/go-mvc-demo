package models

import (
	"time"
)

type User struct {
	ID        uint       `gorm:"comment:主鍵" json:"id"`
	Username  string     `gorm:"type:varchar(20) NOT NULL;comment:用戶名" json:"username"`
	Password  string     `gorm:"type:varchar(100) NOT NULL;comment:密碼" json:"password,omitempty"`
	Status    bool       `gorm:"type:bool;default:1;comment:狀態" json:"status"`
	CreatedAt *time.Time `gorm:"type:datetime NOT NULL;comment:創建時間" json:"created_at,omitempty"`
	UpdatedAt *time.Time `gorm:"type:datetime NOT NULL;comment:更新時間" json:"updated_at,omitempty"`
}

func (User) TableName() string {
	return "user"
}
