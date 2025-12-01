package models

import "time"

type Streak struct {
	ID uint `gorm:"primaryKey"`

	UserID       uint      `gorm:"not null;index"`
	User         User      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Current      int       `gorm:"not null;default:0"`
	Longest      int       `gorm:"not null;default:0"`
	ActivityDate time.Time `gorm:"type:date;default:CURRENT_DATE"`
}
