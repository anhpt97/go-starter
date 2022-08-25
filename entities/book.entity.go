package entities

import "time"

type Book struct {
	ID          uint64     `json:"id"          gorm:"column:id;          primarykey"                        example:"1"`
	Title       *string    `json:"title"       gorm:"column:title;       type:varchar(255)"                 example:"abc"`
	Description *string    `json:"description" gorm:"column:description; type:varchar(255)"                 example:"abc"`
	Content     *string    `json:"content"     gorm:"column:content;     type:text"                         example:"abc"`
	UserID      *uint64    `json:"userId"      gorm:"column:user_id"                                        example:"1"`
	CreatedAt   *time.Time `json:"createdAt"   gorm:"column:created_at;  type:datetime;     autoCreateTime" example:"1970-01-01T00:00:00Z"`
	UpdatedAt   *time.Time `json:"updatedAt"   gorm:"column:updated_at;  type:datetime;     autoUpdateTime" example:"1970-01-01T00:00:00Z"`
	User        *User      `json:"user"        gorm:"foreignKey:UserID;  references:ID"`
}

func (Book) TableName() string {
	return "book"
}
