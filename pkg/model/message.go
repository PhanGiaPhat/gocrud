package model

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Type    string
	Message string
}
