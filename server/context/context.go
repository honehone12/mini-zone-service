package context

import (
	"github.com/labstack/echo/v4"
)

type ZoneList struct {
	Address   string
	HomePort  string
	Zone1Port string
	Zone2Port string
}

type Metadata struct {
	Name    string
	Version string
}

type Context struct {
	echo.Context
	Metadata
	ZoneList
}
