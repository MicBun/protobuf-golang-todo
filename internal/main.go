package internal

import (
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"net"
	"os"
)

type App struct {
	grpc   *grpc.Server
	mainDB *gorm.DB
}

func NewApp(
	grpc *grpc.Server,
	mainDB *gorm.DB,
) *App {
	app := &App{
		grpc,
		mainDB,
	}

	return app
}

func (a *App) Serve() error {
	lis, err := net.Listen("tcp", os.Getenv("GRPC_PORT"))
	if err != nil {
		return err
	}

	return a.grpc.Serve(lis)
}

func (a *App) Clean() error {
	a.grpc.Stop()

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
