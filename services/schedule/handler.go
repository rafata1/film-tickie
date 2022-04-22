package schedule

import (
	"github.com/gin-gonic/gin"
	"github.com/templateOfService/services/auth"
	"net/http"
	"strconv"
	"strings"
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

// HoldSeats godoc
// @Summary      hold seats for 10 minutes before payment
// @Tags         Ticket service
// @Accept       json
// @Produce      json
// @Success      200
// @Router       /api/v1/ticket/hold [post]
// @param 		 Authorization header string true "Authorization"
// @Param        HoldSeatsRequest body HoldSeatsRequest true "hold seats"
func (h *Handler) HoldSeats(c *gin.Context) {
	payload, err := auth.GetGlobalJWTManager().
		VerifyToken(strings.ReplaceAll(c.GetHeader("authorization"), "Bearer ", ""))
	if err != nil {
		c.JSON(http.StatusUnauthorized, BaseRes{
			Message: "unauthorized",
		})
		return
	}

	var req HoldSeatsRequest
	err = c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, BaseRes{
			Message: "invalid request format",
		})
		return
	}

	err = h.service.HoldSeats(payload.Phone, req.ScheduleId, req.Seats)
	if err != nil {
		c.JSON(http.StatusInternalServerError, BaseRes{
			Message: "internal error",
			Data:    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, BaseRes{
		Message: "success",
	})
}

// ConfirmSeats godoc
// @Summary      confirm seats after payment
// @Tags         Ticket service
// @Accept       json
// @Produce      json
// @Success      200
// @Router       /api/v1/ticket/confirm [post]
// @param 		 Authorization header string true "Authorization"
// @Param        ConfirmSeatsRequest body ConfirmSeatsRequest true "confirm after payment"
func (h *Handler) ConfirmSeats(c *gin.Context) {
	payload, err := auth.GetGlobalJWTManager().
		VerifyToken(strings.ReplaceAll(c.GetHeader("authorization"), "Bearer ", ""))
	if err != nil {
		c.JSON(http.StatusUnauthorized, BaseRes{
			Message: "unauthorized",
		})
		return
	}

	var req ConfirmSeatsRequest
	err = c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, BaseRes{
			Message: "invalid request format",
		})
		return
	}

	err = h.service.ConfirmSeats(payload.Phone, req.ScheduleId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, BaseRes{
			Message: "internal error",
			Data:    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, BaseRes{
		Message: "success",
	})
}

// ListSeats godoc
// @Summary      get seats status of schedule
// @Tags         Ticket service
// @Accept       json
// @Produce      json
// @Success      200
// @Router       /api/v1/ticket/ [get]
// @param 		 Authorization header string true "Authorization"
// @Param        scheduleId query int true "schedule id"
func (h *Handler) ListSeats(c *gin.Context) {
	_, err := auth.GetGlobalJWTManager().
		VerifyToken(strings.ReplaceAll(c.GetHeader("authorization"), "Bearer ", ""))
	if err != nil {
		c.JSON(http.StatusUnauthorized, BaseRes{
			Message: "unauthorized",
		})
		return
	}

	scheduleId, err := ParseInt(c.Query("scheduleId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, BaseRes{Message: "invalid schedule id"})
		return
	}

	holdingSeats, orderedSeats, err := h.service.ListSeats(scheduleId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, BaseRes{
			Message: "internal error",
			Data:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, BaseRes{
		Message: "success",
		Data: ListSeatsResponse{
			HoldingSeats: holdingSeats,
			OrderedSeats: orderedSeats,
		},
	})
}

// CancelSeats godoc
// @Summary      user cancel payment then call this api
// @Tags         Ticket service
// @Accept       json
// @Produce      json
// @Success      200
// @Router       /api/v1/ticket/cancel [post]
// @param 		 Authorization header string true "Authorization"
// @Param        CancelSeatsRequest body CancelSeatsRequest true "cancel seats"
func (h *Handler) CancelSeats(c *gin.Context) {
	payload, err := auth.GetGlobalJWTManager().
		VerifyToken(strings.ReplaceAll(c.GetHeader("authorization"), "Bearer ", ""))
	if err != nil {
		c.JSON(http.StatusUnauthorized, BaseRes{
			Message: "unauthorized",
		})
		return
	}

	var req CancelSeatsRequest
	err = c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, BaseRes{
			Message: "invalid request format",
		})
		return
	}

	err = h.service.CancelSeats(payload.Phone, req.ScheduleId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, BaseRes{
			Message: "internal error",
			Data:    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, BaseRes{
		Message: "success",
	})
}
