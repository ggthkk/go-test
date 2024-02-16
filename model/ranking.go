package model

import "time"

type Ranking struct {
	DateTime    time.Time `json:"datetime"`
	Name        string    `json:"name"`
	Lavel       int       `json:"lavel"`
	Descirption string    `json:"descirption"`
}
