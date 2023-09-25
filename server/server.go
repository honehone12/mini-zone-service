package server

import (
	"mini-zone-service/server/context"
	"mini-zone-service/server/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func Run(
	metadata context.Metadata,
	listenAt string,
	zones context.ZoneList,
) {
	e := echo.New()
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := &context.Context{
				Context:  c,
				Metadata: metadata,
				ZoneList: zones,
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
