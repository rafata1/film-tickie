package film

import (
    "github.com/gin-gonic/gin"
    "github.com/templateOfService/services/auth"
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

// ListFilms godoc
// @Summary      ListAllFilms
// @Tags         Film service
// @Accept       json
// @Produce      json
// @param Authorization header string true "Authorization"
// @Success      200
// @Router       /api/v1/films [get]
func (h *Handler) ListFilms(c *gin.Context) {
    _, err := auth.GetGlobalJWTManager().VerifyToken(
        strings.ReplaceAll(c.GetHeader("authorization"), "Bearer ", ""))
    if err != nil {
        c.JSON(http.StatusUnauthorized, BaseRes{
            Message: "unauthorized",
        })
        return
    }

    films, err := h.service.ListFilms()
    if err != nil {
        c.JSON(http.StatusOK, BaseRes{
            Message: "internal error",
            Data:    err.Error(),
        })
        return
    }
    res := BaseRes{
        Message: "success",
        Data:    films,
    }
    c.JSON(http.StatusOK, res)
}
