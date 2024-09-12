package models

import (
	"time"
)

type User struct {
	Id        int     `gorm:"type:int;autoIncrement;primaryKey" json:"id,omitempty"`
	Name      string  `gorm:"type:varchar(255);not null" json:"name,omitempty"`
	Email     string  `gorm:"type:varchar(255);not null;unique" json:"email,omitempty"`
	Username  string  `gorm:"type:varchar(255);not null;unique" json:"username,omitempty"`
	Password  *string `gorm:"type:varchar(255);not null" json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Links     []Link `gorm:"foreignKey:UserId;refences:Id" json:"link,omitempty"`
}
