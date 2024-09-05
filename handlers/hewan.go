package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "klinik-hewan/models"
)


// CreateHewan membuat entitas Hewan baru
func CreateHewan(c *gin.Context) {
    var hewan models.Hewan
    if err := c.ShouldBindJSON(&hewan); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := db.Create(&hewan).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, hewan)
}

// GetHewan mendapatkan semua entitas Hewan
func GetHewan(c *gin.Context) {
    var hewans []models.Hewan
    if err := db.Find(&hewans).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, hewans)
}

// UpdateHewan memperbarui entitas Hewan berdasarkan ID
func UpdateHewan(c *gin.Context) {
    id := c.Param("id")
    var hewan models.Hewan
    if err := db.First(&hewan, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Hewan not found"})
        return
    }
    if err := c.ShouldBindJSON(&hewan); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := db.Save(&hewan).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, hewan)
}

// DeleteHewan menghapus entitas Hewan berdasarkan ID
func DeleteHewan(c *gin.Context) {
    id := c.Param("id")
    if err := db.Delete(&models.Hewan{}, id).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Hewan deleted successfully"})
}
