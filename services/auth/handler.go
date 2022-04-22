package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type Handler struct {
	service *Service
}

func NewHandler() *Handler {
	return &Handler{
		service: NewService(),
	}
}

// CheckOTP godoc
// @Summary      Check otp with phone number, return jwt token
// @Tags         Auth service
// @Accept       json
// @Produce      json
// @Param        checkOTPRequest body CheckOTPRequest true "CheckOTP"
// @Success      200
// @Router       /api/v1/auth/check_otp [post]
func (h *Handler) CheckOTP(c *gin.Context) {
	var req CheckOTPRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, BaseRes{
			Message: "request format is invalid",
			Data:    err,
		})
		return
	}

	if req.OTP != "123456" {
		c.JSON(http.StatusUnauthorized, BaseRes{
			Message: "otp is incorrect",
		})
		return
	}

	jwtManager := GetGlobalJWTManager()
	token := jwtManager.GenerateToken(req.Phone)
	res := BaseRes{
		Message: "success",
		Data: CheckOTPResponseData{
			Token: token,
		},
	}
	c.JSON(http.StatusOK, res)
}

// UpdateUserInfo godoc
// @Summary      update user info
// @Tags         Profile service
// @Accept       json
// @Produce      json
// @Success      200
// @Router       /api/v1/user [post]
// @param 		 Authorization header string true "Authorization"
// @Param        UpdateUserInfoRequest body UpdateUserInfoRequest true "update user info request"
func (h *Handler) UpdateUserInfo(c *gin.Context) {
	payload, err := GetGlobalJWTManager().
		VerifyToken(strings.ReplaceAll(c.GetHeader("authorization"), "Bearer ", ""))
	if err != nil {
		c.JSON(http.StatusUnauthorized, BaseRes{
			Message: "unauthorized",
		})
		return
	}

	var req UpdateUserInfoRequest
	err = c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, BaseRes{
			Message: "request format is invalid",
			Data:    err,
		})
		return
	}

	err = h.service.UpdateUserInfo(payload.Phone, req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, BaseRes{
			Message: "failed",
		})
		return
	}
	c.JSON(http.StatusOK, BaseRes{
		Message: "success",
	})
}
