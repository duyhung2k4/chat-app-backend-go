package model

import (
	"time"

	"gorm.io/gorm"
)

type Text struct {
	Id        uint   `json:"id"`
	MessageId uint   `json:"messageId"`
	Content   string `json:"content"`

	Message *Message `json:"message" gorm:"foreignKey:MessageId; contraint:OnUpdate:CASCADE,OnDelete:SET NULL"`

	CreatedAt time.Time      `json:"createdAt" gorm:"autoCreateTime:true"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"autoUpdateTime:true"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}
