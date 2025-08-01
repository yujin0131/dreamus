package models

import (
	"time"
)

type Build struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	CommitID  string     `gorm:"size:40;not null" json:"commit_id"`
	DockerTag string     `gorm:"size:225;not null" json:"docker_tag"`
	Status    string     `gorm:"size:20;not null" json:"status"`
	StartTime *time.Time `gorm:"type:datetime" json:"start_time"`
	EndTime   *time.Time `gorm:"type:datetime" json:"end_time"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
}

type BuildUpdateInput struct {
	DockerTag string     `gorm:"size:225;not null" json:"docker_tag"`
	Status    string     `gorm:"size:20;not null" json:"status"`
	StartTime *time.Time `gorm:"type:datetime" json:"start_time"`
	EndTime   *time.Time `gorm:"type:datetime" json:"end_time"`
}
