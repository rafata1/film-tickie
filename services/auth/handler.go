package auth

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

type Handler struct{}

func NewHandler() *Handler {
    return &Handler{}
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
