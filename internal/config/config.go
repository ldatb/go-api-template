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

package config

import (
	"log"

	"github.com/spf13/viper"
)

// Config represents the application configuration.
type Config struct {
	// Server is the Shieldwall API configuration
	Server ServerConfig `mapstructure:"General"`

	// Database is the database's configuration
	Database DatabaseConfig `mapstructure:"Database"`
}

// ServerConfig represents the server's configuration
type ServerConfig struct {
	Address  string `mapstructure:"api_address"`
	Port     int    `mapstructure:"api_port"`
	LogFile  string `mapstructure:"log_file"`
	LogLevel string `mapstructure:"log_level"`
}

// DatabaseConfig represents the config of the server's database
type DatabaseConfig struct {
	DBName       string `mapstructure:"db_name" `
	User         string `mapstructure:"db_user"`
	Password     string `mapstructure:"db_password"`
	Address      string `mapstructure:"db_address"`
	Port         int    `mapstructure:"db_port"`
	MaxIdleConns int    `mapstructure:"db_max_idle_conns"`
	MaxOpenConns int    `mapstructure:"db_max_open_conns"`
}

// LoadConfig loads the configuration from a file.
func LoadConfig(configPath string) *Config {
	// Set Viper configs
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(configPath)
	viper.AutomaticEnv()

	// Try to read file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	// Try to parse config
	config := &Config{}
	if err := viper.Unmarshal(config); err != nil {
		log.Fatalf("Unable to decode config file into struct: %v", err)
	}

	// Validate if all configs are set
	err := config.validateConfig()
	if err != nil {
		log.Fatalf("Error in configuration file: %v", err)
	}

	return config
}
