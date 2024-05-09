package models

import "gorm.io/gorm"

type Tile struct {
	gorm.Model
	Landscape   string `json:"landscape" gorm:"text;not null;default:null"`
	Description string `json:"description" gorm:"text;not null;default:null"`
	Accessible  bool   `json:"accessible" gorm:"bool;not null;default:true"`
}
