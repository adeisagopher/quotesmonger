package model

import "gorm.io/gorm"

type Quotes struct {
	gorm.Model
	Quote    string
	Author   string
	Category string
}
