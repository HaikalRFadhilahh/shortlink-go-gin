package models

import (
	"time"
)

type Link struct {
	Id        int    `gorm:"type:int;autoIncrement;primaryKey" json:"id,omitempty"`
	IsActive  bool   `gorm:"type:bool;default:true" json:"isActive,omitempty"`
	Alias     string `gorm:"type:varchar(255);not null;unique" json:"alias,omitempty"`
	Link      string `gorm:"type:varchar(255);not null" json:"link,omitempty"`
	UserId    int    `gorm:"type:int" json:"userId,omitempty"`
	CreatedAt time.Time
	UpdatedAt time.Time
	User      User `gorm:"foreignKey:UserId;refences:Id" json:"user,omitempty"`
}
