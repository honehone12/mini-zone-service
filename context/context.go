package context

import (
	"github.com/labstack/echo/v4"
)

type ZoneList struct {
	Home  string
	Zone1 string
	Zone2 string
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
