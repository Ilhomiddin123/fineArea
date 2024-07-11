package repository

import (
	"fineArea/db"
	"fineArea/models"
)

func SaveVehicle(vehicle *models.VehicleData) error {
	return db.GetConn().Save(vehicle).Error
}
