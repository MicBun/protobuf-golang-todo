package internal

import (
	"context"
	"fmt"
	"github.com/MicBun/protobuf-golang-todo/internal/infra/route"
	souin_echo "github.com/darkweak/souin/plugins/echo"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
	"net/http"
	"os"
)

type App struct {
	server    *echo.Echo
	mainDB    *gorm.DB
	httpRoute *route.HTTP
}

func NewApp(
	server *echo.Echo,
	mainDB *gorm.DB,
	httpRoute *route.HTTP,
) (*App, error) {
	app := &App{
		server,
		mainDB,
		httpRoute,
	}

	s := souin_echo.NewMiddleware(souin_echo.DevDefaultConfiguration)
	router := app.server
	router.Use(s.Process)
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())
	router.Validator = &CustomValidator{validator: validator.New()}
	router.HTTPErrorHandler = customHTTPErrorHandler
	routerGroup := router.Group(os.Getenv("APP_ROUTE_PREFIX"))
	httpRoute.Load(routerGroup)

	return app, nil
}

func (a *App) Serve() error {
	return a.server.Start(os.Getenv("APP_ADDRESS"))
}

func (a *App) Clean() error {
	shutdownAppErr := a.server.Shutdown(context.Background())
	if shutdownAppErr != nil {
		return shutdownAppErr
	}

	sqlDB, getDBErr := a.mainDB.DB()
	if getDBErr != nil {
		return getDBErr
	}

	closeDBErr := sqlDB.Close()
	if closeDBErr != nil {
		return closeDBErr
	}

	return nil
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return nil
}

func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}

	c.Logger().Error(err)
	errJson := c.JSON(code, echo.Map{
		"message": fmt.Sprintf("%v", err),
	})
	if errJson != nil {
		return
	}
}
