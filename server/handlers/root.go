package handlers

import (
	"mini-zone-service/server/context"
	"net/http"

	"github.com/labstack/echo/v4"
)

type RootResponse struct {
	Name    string
	Version string
}

type ZoneResponse struct {
	Address string
	Port    string
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
		Address: ctx.Address,
		Port:    ctx.HomePort,
	})
}

func Zone1(c echo.Context) error {
	ctx := c.(*context.Context)
	return c.JSON(http.StatusOK, ZoneResponse{
		Address: ctx.Address,
		Port:    ctx.Zone1Port,
	})
}

func Zone2(c echo.Context) error {
	ctx := c.(*context.Context)
	return c.JSON(http.StatusOK, ZoneResponse{
		Address: ctx.Address,
		Port:    ctx.Zone2Port,
	})
}
