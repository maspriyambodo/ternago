package controllers

import (
	"net/http"
	"strconv"
	"time"

	"ternago/backend-api/models"

	"github.com/gin-gonic/gin"
)

// GetAuthMstAdmins retrieves all AuthMstAdmin records
func GetAuthMstAdmins(c *gin.Context) {
	var admins []models.AuthMstAdmin
	models.DB.Find(&admins)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    admins,
	})
}

// GetAuthMstAdmin retrieves a specific AuthMstAdmin by ID
func GetAuthMstAdmin(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid ID format",
		})
		return
	}

	var admin models.AuthMstAdmin
	result := models.DB.First(&admin, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Admin not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    admin,
	})
}

// CreateAuthMstAdmin creates a new AuthMstAdmin record
func CreateAuthMstAdmin(c *gin.Context) {
	var admin models.AuthMstAdmin

	// Bind JSON input
	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid request body",
			"details": err.Error(),
		})
		return
	}

	// Create the admin in the database
	result := models.DB.Create(&admin)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to create admin",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    admin,
	})
}

// UpdateAuthMstAdmin updates an existing AuthMstAdmin record
func UpdateAuthMstAdmin(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid ID format",
		})
		return
	}

	var admin models.AuthMstAdmin
	result := models.DB.First(&admin, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Admin not found",
		})
		return
	}

	// Bind JSON input
	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid request body",
			"details": err.Error(),
		})
		return
	}

	// Update the admin in the database
	result = models.DB.Save(&admin)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to update admin",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    admin,
	})
}

// DeleteAuthMstAdmin deletes an AuthMstAdmin record
func DeleteAuthMstAdmin(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid ID format",
		})
		return
	}

	var admin models.AuthMstAdmin
	result := models.DB.First(&admin, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Admin not found",
		})
		return
	}

	// Delete the admin from the database
	result = models.DB.Delete(&admin)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to delete admin",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Admin deleted successfully",
	})
}
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":    "healthy",
		"service":   "gin-app",
		"timestamp": time.Now().Unix(),
	})
}
