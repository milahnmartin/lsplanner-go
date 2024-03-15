package controllers

import (
	"context"
	"lsplanner-go/models"
	"lsplanner-go/repositories"
	"lsplanner-go/services"
	"net/http"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
)

func CreateUser(ctx context.Context, c *app.RequestContext, userRepo repositories.UserRepo) {
	var user models.User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	if user, err := services.GetUserByID(ctx, int(user.ID), userRepo); user != nil || err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": "Someting went wrong, user possibly already exists or error fetching user by ID"})
		return
	}

	if err := services.CreateUser(ctx, &user, userRepo); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, &user)
}

func GetAllUsers(ctx context.Context, c *app.RequestContext, userRepo repositories.UserRepo) {
	users, err := services.GetAllUsers(ctx, userRepo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func GetUser(ctx context.Context, c *app.RequestContext, userRepo repositories.UserRepo) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid user ID"})
		return
	}

	user, err := services.GetUserByID(ctx, id, userRepo)
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

func UpdateUser(ctx context.Context, c *app.RequestContext, userRepo repositories.UserRepo) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid user ID"})
		return
	}

	var user models.User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	if err := services.UpdateUserByID(ctx, id, &user, userRepo); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, &user)
}

func DeleteUser(ctx context.Context, c *app.RequestContext, userRepo repositories.UserRepo) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid user ID"})
		return
	}

	if err := services.DeleteUserByID(ctx, id, userRepo); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]string{"message": "user deleted successfully"})
}

func GetUserAreaID(ctx context.Context, c *app.RequestContext, userRepo repositories.UserRepo) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid user ID"})
		return
	}

	areaID, err := userRepo.GetUserAreaID(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]uint{"area_id": areaID})
}
