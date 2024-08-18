package models

import "gorm.io/gorm"

type Services struct {
	gorm.Model
	Name          string
	URL           string
	LastCheckTime int64
	Status        int
}
