package routes

import (
	"context"
	"lsplanner-go/controllers"
	"lsplanner-go/repositories"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/route"
)

func QuotaRoutes(group *route.RouterGroup, quotaRepo repositories.QuotaRepo) {
	group.GET("/:id", func(ctx context.Context, c *app.RequestContext) {
		controllers.GetByID(ctx, c, quotaRepo)
	})

	group.POST("/init/:id", func(ctx context.Context, c *app.RequestContext) {
		controllers.Init(ctx, c, quotaRepo)
	})
}
