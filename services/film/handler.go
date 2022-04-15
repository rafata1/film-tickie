package film

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

type Handler struct {
    service *Service
}

func NewHandler() *Handler {
    return &Handler{
        service: NewService(),
    }
}

func ParseInt(s string) (int, error) {
    if len(s) == 0 {
        return 0, nil
    }
    number, err := strconv.Atoi(s)
    return number, err
}

// GetFilmById godoc
// @Summary      GetFilmById
// @Tags         Film service
// @Accept       json
// @Produce      json
// @Success      200
// @Router       /api/v1/film/{id} [get]
// @Param id path integer true "id"
func (h *Handler) GetFilmById(c *gin.Context) {
    filmId, err := ParseInt(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, BaseRes{
            Message: "invalid film id",
        })
        return
    }
    film, err := h.service.GetFilmById(filmId)
    if err != nil {
        c.JSON(http.StatusOK, BaseRes{
            Message: "internal error",
            Data:    err.Error(),
        })
        return
    }
    res := BaseRes{
        Message: "success",
        Data:    film,
    }
    c.JSON(http.StatusOK, res)
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
