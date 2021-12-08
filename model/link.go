package model

import "time"

// Link Model
type Link struct {
	ID          string `gorm:"primary_key"`                       // ID
	Target      string `gorm:"varchar(4096);not null;default ''"` // Target link
	Description string `gorm:"varchar(4096);not null;default ''"` // Description
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
