package controllers

import (
	"context"
	"lsplanner-go/config"
	"lsplanner-go/models"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
)

func setupQuota(ctx context.Context, c *app.RequestContext) {
	var db, err = config.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": "database not found"})
		return
	}
	var quota models.Quota
	if err := c.Bind(&quota); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	if result := db.Create(&quota); result.Error != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, &quota)
}

func getQuota(ctx context.Context, c *app.RequestContext) {
	var db, err = config.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": "database not found"})
		return
	}
	var quota []models.Quota
	if result := db.Find(&quota); result.Error != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, &quota)
}
