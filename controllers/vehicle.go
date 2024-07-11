package controllers

import (
	"crypto/sha256"
	"encoding/base64"
	"fineArea/models"
	"fineArea/service"
	"fineArea/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
)

func UploadVehicle(c *gin.Context) {
	// Получение данных файла и номера машины из запроса
	requestData := struct {
		File   string `json:"file"`
		Number string `json:"number"`
	}{}

	if err := c.BindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request data",
		})
		return
	}

	requestHash := c.GetHeader("Authorization")
	localHash := utils.GenerateMD5Hash(requestData.File + requestData.Number)
	if requestHash != localHash {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid request hash",
		})
		return
	}

	// Декодирование файла из base64
	fileData, err := base64.StdEncoding.DecodeString(requestData.File)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid file data",
		})
		return
	}

	// Вычисление хэша файла
	hash := calculateFileHash(fileData)

	// Сохранение файла в папке storage
	filePath := saveFileFromByteArray(fileData, hash)

	// Сохранение информации о файле и номере машины в базе данных
	vehicleData := &models.VehicleData{
		Hash:   hash,
		Number: requestData.Number,
		Path:   filePath,
	}

	err = service.SaveVehicle(vehicleData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to save vehicle data",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "File uploaded and data saved successfully",
	})
}

func calculateFileHash(fileData []byte) string {
	h := sha256.New()
	h.Write(fileData)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func saveFileFromByteArray(fileData []byte, hash string) string {
	dst := filepath.Join("storage", fmt.Sprintf("%s_%s", hash, ".png"))
	err := os.WriteFile(dst, fileData, 0644)
	if err != nil {
		return ""
	}
	return dst
}
