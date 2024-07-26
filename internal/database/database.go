/*
Copyright 2024 Lucas de Ataides

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package database

import (
	"fmt"
	"time"

	"github.com/ldatb/go-api-template/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// ConnectToDB creates a connection with a MySQL database
func ConnectToDB(dbConfig *config.DatabaseConfig) (*gorm.DB, error) {
	// dsn is how to connect to the database. It stands for data source name
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.User, dbConfig.Password,
		dbConfig.Address, dbConfig.Port,
		dbConfig.DBName,
	)

	// Open a gorm connection
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}

	// Configure connections
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConns)
	sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Minute * 10)

	return db, nil
}

// MakeInitialMigrations migrates all models to the database
func MakeInitialMigrations(db *gorm.DB) error {
	return db.AutoMigrate()
}
