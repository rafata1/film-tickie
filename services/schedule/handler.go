package schedule

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

// ListSchedules godoc
// @Summary      ListSchedules
// @Tags         Schedule service
// @Accept       json
// @Produce      json
// @Success      200
// @Router       /api/v1/schedules [get]
// @Param cinemaId query integer false "cinemaId"
// @Param filmId query integer false "filmId"
func (h *Handler) ListSchedules(c *gin.Context) {
    cinemaId, err := ParseInt(c.Query("cinemaId"))
    if err != nil {
        c.JSON(http.StatusBadRequest, BaseRes{
            Message: "invalid cinema id",
        })
        return
    }
    filmId, err := ParseInt(c.Query("filmId"))
    if err != nil {
        c.JSON(http.StatusBadRequest, BaseRes{
            Message: "invalid film id",
        })
        return
    }

    if filmId == 0 && cinemaId == 0 {
        c.JSON(http.StatusBadRequest, BaseRes{
            Message: "request must include one of film id or cinema id",
        })
        return
    }

    schedules, err := h.service.ListSchedules(cinemaId, filmId)
    if err != nil {
        c.JSON(http.StatusOK, BaseRes{
            Message: "internal error",
            Data:    err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, BaseRes{
        Message: "success",
        Data:    schedules,
    })
    return
}
