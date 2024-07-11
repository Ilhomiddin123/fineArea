package service

import (
	"fineArea/models"
	"fineArea/repository"
)

func SaveVehicle(v *models.VehicleData) error {
	return repository.SaveVehicle(v)
}
