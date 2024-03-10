package controllers

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// PingTest is a controller method to handle ping requests
func PingTest(c context.Context, ctx *app.RequestContext) {
	ctx.JSON(consts.StatusOK, utils.H{"message": "pong"})
}

// HandleGetUsers is a controller method to handle GET requests for users
func HandleGetUsers(c context.Context, ctx *app.RequestContext) {
	ctx.JSON(consts.StatusOK, utils.H{"message": "pong"})
}
