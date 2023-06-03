package db

import (
	"fmt"
	"github.com/MicBun/protobuf-golang-todo/internal/infra/model"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var gormInstance *gorm.DB

func NewGormDB() (*gorm.DB, error) {
	if gormInstance == nil {
		host := os.Getenv("PG_HOST")
		port := os.Getenv("PG_PORT")
		user := os.Getenv("PG_USER")
		password := os.Getenv("PG_PASSWORD")
		database := os.Getenv("PG_DATABASE")
		dsn := fmt.Sprintf(
			"host=%s port=%v user=%s password=%s dbname=%s",
			host,
			port,
			user,
			password,
			database,
		)

		gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			return nil, err
		}

		gormInstance = gormDB
	}

	if err := gormInstance.AutoMigrate(&model.Todo{}); err != nil {
		return nil, errors.WithStack(err)
	}

	return gormInstance, nil
}

type PostgresTransactionManager struct {
	DB *gorm.DB
}

func NewPostgresTransactionManager(db *gorm.DB) *PostgresTransactionManager {
	return &PostgresTransactionManager{DB: db}
}

func (m *PostgresTransactionManager) Run(callback func(tx any) error) error {
	return m.DB.Transaction(func(tx *gorm.DB) error {
		return callback(tx)
	})
}
