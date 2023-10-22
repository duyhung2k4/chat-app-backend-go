package model

import (
	"time"

	"gorm.io/gorm"
)

type Message struct {
	Id            uint  `json:"id"`
	ChatSectionId *uint `json:"chatSectionId"`
	ProfileIdSend uint  `json:"profileIdSend"`
	GroupId       *uint `json:"groupId"`

	ChatSection *ChatSection `json:"chatSection" gorm:"foreignKey:ChatSectionId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Group       *Group       `json:"group" gorm:"foreignKey:GroupId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`

	CreatedAt time.Time      `json:"createdAt" gorm:"autoCreateTime:true"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"autoUpdateTime:true"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}
