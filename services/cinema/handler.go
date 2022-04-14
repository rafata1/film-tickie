package cinema

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

// ListCinemas godoc
// @Summary      ListAllCinemas
// @Tags         Cinema service
// @Accept       json
// @Produce      json
// @Success      200
// @Router       /api/v1/cinemas [get]
func (h *Handler) ListCinemas(c *gin.Context) {
    cinemas, err := h.service.ListCinemas()
    if err != nil {
        c.JSON(http.StatusOK, BaseRes{
            Message: "internal error",
            Data:    err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, BaseRes{
        Message: "success",
        Data:    cinemas,
    })
    return
}

func ParseInt(s string) (int, error) {
    if len(s) == 0 {
        return 0, nil
    }
    number, err := strconv.Atoi(s)
    return number, err
}

// GetCinema godoc
// @Summary      GetCinemaById
// @Tags         Cinema service
// @Accept       json
// @Produce      json
// @Success      200
// @Router       /api/v1/cinema/{id} [get]
// @Param id path integer true "cinema id"
func (h *Handler) GetCinema(c *gin.Context) {
    cinemaId, err := ParseInt(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, BaseRes{
            Message: "invalid cinema id",
        })
        return
    }

    cinema, err := h.service.GetCinemaById(cinemaId)
    if err != nil {
        c.JSON(http.StatusOK, BaseRes{
            Message: "internal error",
            Data:    err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, BaseRes{
        Message: "success",
        Data:    cinema,
    })
    return
}
