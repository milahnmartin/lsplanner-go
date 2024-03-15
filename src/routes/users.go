package routes

import (
	"context"
	"lsplanner-go/controllers"
	"lsplanner-go/repositories"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/route"
)

func UserRoutes(group *route.RouterGroup, userRepo repositories.UserRepo) {
	group.GET("/:id", func(ctx context.Context, c *app.RequestContext) {
		controllers.GetUser(ctx, c, userRepo)
	})
	group.GET("/", func(ctx context.Context, c *app.RequestContext) {
		controllers.GetAllUsers(ctx, c, userRepo)
	})
	group.POST("/", func(ctx context.Context, c *app.RequestContext) {
		controllers.CreateUser(ctx, c, userRepo)
	})
	group.PUT("/:id", func(ctx context.Context, c *app.RequestContext) {
		controllers.UpdateUser(ctx, c, userRepo)
	})
	group.DELETE("/:id", func(ctx context.Context, c *app.RequestContext) {
		controllers.DeleteUser(ctx, c, userRepo)
	})
	group.GET("/area/:id", func(ctx context.Context, c *app.RequestContext) {
		controllers.GetUserAreaID(ctx, c, userRepo)
	})
}
