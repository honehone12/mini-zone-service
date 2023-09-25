package server

import (
	"mini-zone-service/context"
	"mini-zone-service/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func Run(
	name string,
	version string,
	listenAt string,
	home string,
	zone1 string,
	zone2 string,
) {
	e := echo.New()
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := &context.Context{
				Context: c,
				Metadata: context.Metadata{
					Name:    name,
					Version: version,
				},
				ZoneList: context.ZoneList{
					Home:  home,
					Zone1: zone1,
					Zone2: zone2,
				},
			}
			return next(ctx)
		}
	})
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	e.GET("/", handlers.Root)
	e.GET("/home", handlers.Home)
	e.GET("/zone1", handlers.Zone1)
	e.GET("/zone2", handlers.Zone2)

	e.Logger.SetLevel(log.WARN)
	e.Logger.Fatal(e.Start(listenAt))
}
