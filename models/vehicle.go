package models

type VehicleData struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Hash   string `json:"hash" gorm:"unique;not null"`
	Number string `json:"number" gorm:"not null"`
	Path   string `json:"path" gorm:"not null"`
}
