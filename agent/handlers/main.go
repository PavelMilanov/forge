package handlers

import (
	"github.com/labstack/echo/v4"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) InitRouters() *echo.Echo {
	e := echo.New()
	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"https://app.uhvahta.ru"},
	// 	AllowMethods: []string{http.MethodGet, http.MethodHead},
	// }))
	// e.GET("/check", h.check)
	// e.Any("/*", h.route)
	deployGroup := e.Group("/deploy")
	deployGroup.POST("/", deployHandler)
	return e
}
