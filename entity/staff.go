package entity

import "time"

type Staff struct {
	ID             uint       `gorm:"primary_key" json:"id"`
	CreatedAt      time.Time  `json:"-"`
	UpdatedAt      time.Time  `json:"-"`
	DeletedAt      *time.Time `sql:"index" json:"-"`
	Name           string     `json:"name"`
	Birthday       time.Time  `json:"birthday"`
	IsActor        bool       `json:"isActor"`
	IsProducer     bool       `json:"isProducer"`
	IsDirector     bool       `json:"isDirector"`
	IsScreenWriter bool       `json:"isScreenWriter"`
}
