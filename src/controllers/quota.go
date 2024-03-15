package controllers

import (
	"context"
	"lsplanner-go/repositories"
	"lsplanner-go/services"
	"net/http"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
)

func GetByID(ctx context.Context, c *app.RequestContext, quotaRepo repositories.QuotaRepo) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid user ID"})
		return
	}

	user, err := services.GetByID(ctx, id, quotaRepo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	if user == nil {
		c.JSON(http.StatusNotFound, map[string]string{"error": "user not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func Init(ctx context.Context, c *app.RequestContext, quotaRepo repositories.QuotaRepo) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid user ID"})
		return
	}
	_, err = services.InitQuota(ctx, id, 50, quotaRepo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}
