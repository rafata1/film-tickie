package film

import (
    "github.com/gin-gonic/gin"
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

// ListFilmsByCategory godoc
// @Summary      ListFilmsByCategory
// @Tags         Film service
// @Accept       json
// @Produce      json
// @Success      200
// @Router       /api/v1/films/{category} [get]
// @Param        category path string true "action"
func (h *Handler) ListFilmsByCategory(c *gin.Context) {
    category := c.Param("category")
    films, err := h.service.ListFilmsByCategory(category)
    if err != nil {
        c.JSON(http.StatusOK, BaseRes{
            Message: "internal error",
            Data:    err.Error(),
        })
        return
    }

    if len(films) == 0 {
        c.JSON(http.StatusNotFound, BaseRes{
            Message: "no films found",
        })
        return
    }

    res := BaseRes{
        Message: "success",
        Data:    films,
    }
    c.JSON(http.StatusOK, res)
}
