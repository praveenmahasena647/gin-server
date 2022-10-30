package dbs

import "gorm.io/gorm"

type ToDo struct {
	gorm.Model
	Act  string `json:"act"`
	Done bool   `json:"done"`
}
