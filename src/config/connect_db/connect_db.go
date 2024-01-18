package connect_db

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ConnectionProvider interface {
	NewConnectionDB() (*gorm.DB, error)
}

type defaultConnectionProvider struct {
	config Config
}

type Config struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

func (p *defaultConnectionProvider) NewConnectionDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		p.config.User,
		p.config.Password,
		p.config.Host,
		p.config.Port,
		p.config.Database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func NewConnectionProviderFromEnv(
	db_user,
	db_password,
	db_host,
	db_port,
	db_database string) ConnectionProvider {

	cfg := Config{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Database: os.Getenv("DB_DATABASE"),
	}

	return &defaultConnectionProvider{config: cfg}
}
