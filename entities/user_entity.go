package entities

import (
	"go-starter/enums"
	"time"
)

type User struct {
	ID             uint64           `json:"id"        gorm:"column:id;              primarykey"`
	Username       string           `json:"username"  gorm:"column:username;        type:varchar(32);                                   uniqueIndex"`
	Email          *string          `json:"email"     gorm:"column:email;           type:varchar(255);                                  uniqueIndex"`
	HashedPassword string           `json:"-"         gorm:"column:hashed_password; type:varchar(64)"`
	Role           enums.UserRole   `json:"role"      gorm:"column:role;            type:enum('ADMIN','USER')"`
	Status         enums.UserStatus `json:"status"    gorm:"column:status;          type:enum('NOT_ACTIVATED','ACTIVE','IS_DISABLED')"`
	CreatedAt      *time.Time       `json:"createdAt" gorm:"column:created_at;      type:datetime;                                      autoCreateTime"`
	UpdatedAt      *time.Time       `json:"updatedAt" gorm:"column:updated_at;      type:datetime;                                      autoUpdateTime"`
	Books          []Book           `json:"books"     gorm:"foreignKey:UserID;      references:ID;                                      constraint:OnDelete:CASCADE"`
}

func (User) TableName() string {
	return "user"
}
