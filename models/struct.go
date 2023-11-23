package models

import (
	"gorm.io/gorm"
)

type Environtment struct {
	gorm.Model
	Wind  int `json:"wind" form:"wind" query:"wind" gorm:"not null"`
	Water int `json:"water" form:"water" query:"water" gorm:"not null"`
}

type Response struct {
	Messages string
	Success  bool
	Data     interface{}
}
