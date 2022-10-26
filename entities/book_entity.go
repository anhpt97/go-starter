package entities

import "time"

type Book struct {
	ID          uint64     `json:"id"          gorm:"column:id;          primarykey"`
	Title       *string    `json:"title"       gorm:"column:title;       type:varchar(255)"`
	Description *string    `json:"description" gorm:"column:description; type:varchar(255)"`
	Content     *string    `json:"content"     gorm:"column:content;     type:text"`
	UserID      *uint64    `json:"userId"      gorm:"column:user_id"`
	CreatedAt   *time.Time `json:"createdAt"   gorm:"column:created_at;  type:datetime;      autoCreateTime"`
	UpdatedAt   *time.Time `json:"updatedAt"   gorm:"column:updated_at;  type:datetime;      autoUpdateTime"`
	User        *User      `json:"user"        gorm:"foreignKey:UserID;  references:ID"`
}

func (Book) TableName() string {
	return "book"
}
