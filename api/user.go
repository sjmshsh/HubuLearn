package api

import (
	"HubuLearn/service"
	"HubuLearn/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRegisterHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req types.UserServiceReq
		if err := context.ShouldBind(&req); err != nil {
			l := service.GetUserSrv()
			resp, err := l.Register(context.Request.Context(), &req)
			if err != nil {
				context.JSON(http.StatusInternalServerError, "")
				return
			}
			context.JSON(http.StatusOK, resp)
		}
	}
}

func UserLoginHandler() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}
