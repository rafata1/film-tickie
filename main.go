package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/templateOfService/connectors/mysql"
	_ "github.com/templateOfService/docs"
	"github.com/templateOfService/jobs"
	"github.com/templateOfService/services/auth"
	"github.com/templateOfService/services/cinema"
	"github.com/templateOfService/services/film"
	"github.com/templateOfService/services/schedule"
	"log"
	"os"
	"sync"
)

func initRouter() *gin.Engine {
	gin.SetMode(os.Getenv("GIN_MODE"))
	router := gin.Default()

	authHandler := auth.NewHandler()
	router.POST("/api/v1/auth/check_otp", authHandler.CheckOTP)
	router.POST("/api/v1/user", authHandler.UpdateUserInfo)

	filmHandler := film.NewHandler()
	router.GET("/api/v1/films", filmHandler.ListFilms)
	router.GET("/api/v1/film/:id", filmHandler.GetFilmById)
	router.GET("/api/v1/films/:category", filmHandler.ListFilmsByCategory)

	cinemaHandler := cinema.NewHandler()
	router.GET("/api/v1/cinemas", cinemaHandler.ListCinemas)
	router.GET("/api/v1/cinema/:id", cinemaHandler.GetCinema)

	scheduleHandler := schedule.NewHandler()
	router.GET("/api/v1/schedules", scheduleHandler.ListSchedules)
	router.GET("/api/v1/ticket", scheduleHandler.ListSeats)
	router.POST("/api/v1/ticket/hold", scheduleHandler.HoldSeats)
	router.POST("/api/v1/ticket/confirm", scheduleHandler.ConfirmSeats)
	router.POST("/api/v1/ticket/cancel", scheduleHandler.CancelSeats)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}

// @title User API documentation
// @BasePath /
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading env: %s", err.Error())
	}

	err = mysql.Connect()
	if err != nil {
		log.Fatalf("Error connecting to mysql: %s", err.Error())
	}

	router := initRouter()
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		err = router.Run()
		if err != nil {
			log.Fatalf("Error starting server: %s", err.Error())
		}
	}()

	go func() {
		jobs.CleanExpiredHoldingSeats()
	}()
	wg.Wait()
}
