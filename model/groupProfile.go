package model

import (
	"time"

	"gorm.io/gorm"
)

type GroupProfile struct {
	Id        uint `json:"id"`
	GroupId   uint `json:"groupId"`
	ProfileId uint `json:"profileId"`

	Group   *Group   `json:"group" gorm:"foreignKey:GroupId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Profile *Profile `json:"profile" gorm:"ProfileId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`

	CreatedAt time.Time      `json:"createdAt" gorm:"autoCreateTime:true"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"autoUpdateTime:true"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}
