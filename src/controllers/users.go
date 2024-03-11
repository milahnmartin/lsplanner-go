package controllers

import (
	"context"
	"errors"
	"lsplanner-go/models"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"gorm.io/gorm"
)

var db *gorm.DB // This should be initialized appropriately, e.g., in your main function or through dependency injection

// Initialize sets the database connection for the controllers.
func Initialize(database *gorm.DB) {
	db = database
}

// CreateUser example
func CreateUser(ctx context.Context, c *app.RequestContext) {
	var user models.User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	if result := db.Create(&user); result.Error != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, &user)
}
func GetAllUsers(ctx context.Context, c *app.RequestContext) {
	var users []models.User
	if result := db.Find(&users); result.Error != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, &users)
}

func getAllUsers(ctx context.Context, c *app.RequestContext) {
	var users []models.User
	if result := db.Find(&users); result.Error != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, &users)
}

// GetUser example
func GetUser(ctx context.Context, c *app.RequestContext) {
	id := c.Param("id")
	var user models.User
	if result := db.First(&user, "id = ?", id); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, map[string]string{"error": "user not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, map[string]string{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, &user)
}

// UpdateUser example
func UpdateUser(ctx context.Context, c *app.RequestContext) {
	id := c.Param("id")
	var user models.User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	if result := db.First(&user, "id = ?", id); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, map[string]string{"error": "user not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, map[string]string{"error": result.Error.Error()})
		return
	}
	if result := db.Save(&user); result.Error != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, &user)
}

// DeleteUser example
func DeleteUser(ctx context.Context, c *app.RequestContext) {
	id := c.Param("id")
	var user models.User
	if result := db.First(&user, "id = ?", id); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, map[string]string{"error": "user not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, map[string]string{"error": result.Error.Error()})
		return
	}
	if result := db.Delete(&user); result.Error != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]string{"message": "user deleted successfully"})
}
