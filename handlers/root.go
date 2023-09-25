package handlers

import (
	"mini-zone-service/context"
	"net/http"

	"github.com/labstack/echo/v4"
)

type RootResponse struct {
	Name    string
	Version string
}

type ZoneResponse struct {
	Url string
}

func Root(c echo.Context) error {
	ctx := c.(*context.Context)
	return c.JSON(http.StatusOK, RootResponse{
		Name:    ctx.Name,
		Version: ctx.Version,
	})
}

func Home(c echo.Context) error {
	ctx := c.(*context.Context)
	return c.JSON(http.StatusOK, ZoneResponse{
		Url: ctx.Home,
	})
}

func Zone1(c echo.Context) error {
	ctx := c.(*context.Context)
	return c.JSON(http.StatusOK, ZoneResponse{
		Url: ctx.Zone1,
	})
}

func Zone2(c echo.Context) error {
	ctx := c.(*context.Context)
	return c.JSON(http.StatusOK, ZoneResponse{
		Url: ctx.Zone2,
	})
}
