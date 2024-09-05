package handlers

import (
	"klinik-hewan/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateProduk membuat entitas Produk baru
func CreateProduk(c *gin.Context) {
	var produk models.Produk
	if err := c.ShouldBindJSON(&produk); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.Create(&produk).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, produk)
}

// GetProduk mendapatkan semua entitas Produk
func GetProduk(c *gin.Context) {
	var produks []models.Produk
	if err := db.Find(&produks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, produks)
}

// UpdateProduk memperbarui entitas Produk berdasarkan ID
func UpdateProduk(c *gin.Context) {
	id := c.Param("id")
	var produk models.Produk
	if err := db.First(&produk, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produk not found"})
		return
	}
	if err := c.ShouldBindJSON(&produk); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.Save(&produk).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, produk)
}

// DeleteProduk menghapus entitas Produk berdasarkan ID
func DeleteProduk(c *gin.Context) {
	id := c.Param("id")
	if err := db.Delete(&models.Produk{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Produk deleted successfully"})
}
