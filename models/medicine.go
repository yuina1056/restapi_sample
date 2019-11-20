package models

import(
	"github.com/jinzhu/gorm"
)

// Medicine struct
type Medicine struct {
	gorm.Model
	Name string `json:"medicineName"`
	Type string `json:"medicineType"`
}
