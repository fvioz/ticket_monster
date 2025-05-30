package server

import (
	"api/handlers"
	"api/libs"
	"net/http"
	"time"

	_ "api/docs"

	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.uber.org/zap"
)

//	@title			Events API
//	@version		1.0
//	@description	Events API server.

// @host		localhost:8080
// @BasePath	/v1
func (s *Server) routes() {
	logger := libs.LoggerInstance()
	planHandler := handlers.NewPlansHandler()

	// @Summary Events Plans
	// @Description get plans within a time range
	// @Tags Plans
	// @Accept json
	// @Produce xml
	// @Param starts_at query string true "Start time in format 2006-01-02T15:04:05"
	// @Param ends_at query string true "End time in format 2006-01-02T15:04:05"
	// @Success 200 {array} handlers.Plan
	// @Failure 400 {object} echo.HTTPError "Invalid parameters"
	// @Failure 500 {object} echo.HTTPError "Internal server error"
	// @Router /v1/events/plans [get]
	s.echo.GET("v1/events/plans", func(c echo.Context) error {
		startsAt, err := s.ParseTime(c.QueryParam("starts_at"))
		if err != nil {
			logger.Error("Failed to parse start time", zap.Error(err))
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid start time format, start_time parameter is required")
		}

		endsAt, err := s.ParseTime(c.QueryParam("ends_at"))
		if err != nil {
			logger.Error("Failed to parse end time", zap.Error(err))
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid end time format, end_time parameter is required")
		}

		plans, err := planHandler.GetPlansV1(startsAt, endsAt)
		if err != nil {
			logger.Error("Failed to fetch plans", zap.Error(err))

			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}

		return c.XMLPretty(http.StatusOK, &plans, "  ")
	})

	s.echo.GET("/swagger/*", echoSwagger.WrapHandler)
	s.echo.GET("/metrics", echoprometheus.NewHandler())
	s.echo.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
}

func (s *Server) ParseTime(t string) (time.Time, error) {
	return time.Parse("2006-01-02T15:04:05", t)
}
