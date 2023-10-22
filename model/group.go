package model

import (
	"time"

	"gorm.io/gorm"
)

type Group struct {
	Id        uint   `json:"id"`
	CreatorId uint   `json:"creatorId"`
	Name      string `json:"name"`
	Members   uint   `json:"members"`

	Creator *Profile `json:"creator" gorm:"foreignKey:CreatorId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`

	CreatedAt time.Time      `json:"createdAt" gorm:"autoCreateTime:true"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"autoUpdateTime:true"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}
