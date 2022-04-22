package schedule

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
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

func ParseStringToDate(date string) (*time.Time, error) {
	if len(date) == 0 {
		return nil, nil
	}

	parsedDate, err := time.Parse("2006-01-02", date)
	return &parsedDate, err
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
// @Param onDate query string false "filter by date, ex: 2018-01-20"
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

	onDate, err := ParseStringToDate(c.Query("onDate"))
	if err != nil {
		c.JSON(http.StatusBadRequest, BaseRes{
			Message: "invalid on date",
		})
		return
	}

	schedules, err := h.service.ListSchedules(cinemaId, filmId, onDate)
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
