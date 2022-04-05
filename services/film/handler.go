package film

import (
    "github.com/gin-gonic/gin"
    "log"
    "net/http"
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
// @Success      200
// @Router       /api/v1/films [get]
func (h *Handler) ListFilms(c *gin.Context) {
    log.Printf("here")

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
