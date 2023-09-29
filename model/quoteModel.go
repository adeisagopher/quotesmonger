package model

import "gorm.io/gorm"

type Quotes struct {
	gorm.Model
	Quote    string `json:"quote"`
	Author   string `json:"author"`
	Category string `json:"category"`
}
