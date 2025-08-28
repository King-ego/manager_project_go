package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
}

func LoadDbConfig() DbConfig {
	return DbConfig{
		Host:     getEnv("DB_HOST", "manager_db"),
		Port:     getEnvAsInt("DB_PORT", 5432),
		User:     getEnv("POSTGRES_USER", "test"),
		Password: getEnv("POSTGRES_PASSWORD", "test"),
		Dbname:   getEnv("POSTGRES_DB", "test_db"),
	}
}

func ConnectDb() (*gorm.DB, error) {
	config := LoadDbConfig()
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Dbname,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})

	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	sqlDB, err := db.DB()

	if err != nil {
		return nil, fmt.Errorf("failed to get generic database object: %w", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	} else {
		return defaultValue
	}
}

func getEnvAsInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
