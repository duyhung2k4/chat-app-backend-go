package model

type ChatSection struct {
	Id          uint `json:"id"`
	ProfileId_1 uint `json:"profileId_1"`
	ProfileId_2 uint `json:"profileId_2"`

	Profile_1 *Profile `json:"profile_1" gorm:"foreignKey:ProfileId_1; constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Profile_2 *Profile `json:"profile_2" gorm:"foreignKey:ProfileId_2; constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
