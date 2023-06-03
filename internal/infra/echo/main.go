package echo

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

var ProviderSet = wire.NewSet(New)

func New() *echo.Echo {
	return echo.New()
}
