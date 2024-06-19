package models

import "gorm.io/gorm"

type Services struct {
	gorm.Model
	Name          string
	Command       string
	LastCheckTime int64
	Status        int
}
