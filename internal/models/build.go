package models

import (
	"time"
)

type Build struct {
	ID        uint      `gorm:"primaryKey"`
	CommitID  string    `gorm:"size:40;not null`
	DockerTag string    `gorm:"size:225;not null`
	Status    string    `gorm:"size:20;not null`
	StartTime time.Time `gorm:"type:datetime"`
	EndTime   time.Time `gorm:"type:datetime"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
