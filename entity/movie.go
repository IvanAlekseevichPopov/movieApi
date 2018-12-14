package entity

import (
	"time"
)

type Movie struct {
	ID          uint       `json:"id" gorm:"primary_key"`
	CreatedAt   time.Time  `json:"-"`
	UpdatedAt   time.Time  `json:"-"`
	DeletedAt   *time.Time `json:"-" sql:"index" `
	Title       string     `json:"name"`
	Description string     `json:"description"`
	Actors      []Staff    `json:"actors" gorm:"many2many:movie_actors"`
}
