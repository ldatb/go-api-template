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

package main

import (
	"flag"
	"fmt"

	"github.com/labstack/echo/v4"
	v1 "github.com/ldatb/go-api-template/api/v1"
	"github.com/ldatb/go-api-template/internal/config"
	"github.com/ldatb/go-api-template/internal/database"
	log "github.com/ldatb/go-api-template/internal/logger"
	"github.com/ziflex/lecho/v3"
)

// Main is the entrypoint for the application
func main() {
	// Load application parameters and flags
	configPath := flag.String("config-dir", ".", "The directory of the config.yaml file")
	flag.Parse()

	// Load configuration
	config := config.LoadConfig(string(*configPath))

	// Initialize logger instance
	logger := log.InitializeMainLogger(config.Server.LogFile, config.Server.LogLevel)

	// Connect and make database migrations
	db, err := database.ConnectToDB(&config.Database)
	if err != nil {
		logger.Fatal("Failed to connect to database: %v", err)
	}
	err = database.MakeInitialMigrations(db)
	if err != nil {
		logger.Fatal("Failed to make database migrations: %v", err)
	}

	// Create a new Echo instance
	router := echo.New()
	router.HideBanner = true
	router.Logger = lecho.From(logger.Log)

	// Register routes
	v1.RegisterRoutes(router)

	// Start the server
	apiFullAddr := fmt.Sprintf("%s:%v", config.Server.Address, config.Server.Port)
	logger.Info("Starting server on %s", apiFullAddr)
	if err := router.Start(apiFullAddr); err != nil {
		logger.Fatal("Failed to start server: %v", err)
	}
}
