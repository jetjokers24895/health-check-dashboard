package models

import "gorm.io/gorm"

type LogChecked struct {
	gorm.Model
	HttpStatus  string
	ResponseTXT string
	Status      int

	ServicesID int
	Services   Services
}
