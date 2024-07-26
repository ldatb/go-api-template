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
	"fmt"
)

// ValidateConfig makes sure all configuration fields are valid
func (c *Config) validateConfig() error {
	// Check if port is between 1000 and 9999
	if c.Server.Port < 1000 || c.Server.Port > 9999 {
		return fmt.Errorf("defined port has to be between 1000 and 9999")
	}

	// Make sure log level is valid
	validLogLevels := []string{"debug", "info", "warning", "error", "fatal"}
	logLevelValid := false
	for _, level := range validLogLevels {
		if c.Server.LogLevel == level {
			logLevelValid = true
		}
	}
	if !logLevelValid {
		return fmt.Errorf("requested log level %s is not valid", c.Server.LogLevel)
	}

	// Make sure all database string fields are filled
	databaseStringValues := []string{c.Database.DBName, c.Database.User,
		c.Database.Password, c.Database.Address}
	databaseFieldNames := []string{"db_name", "user", "password", "address", "port", "ssl_mode"}
	for i, value := range databaseStringValues {
		if value == "" {
			return fmt.Errorf("required database field %s is empty", databaseFieldNames[i])
		}
	}

	// Make sure all database numerical fields are filled
	if c.Database.MaxIdleConns == 0 {
		c.Database.MaxIdleConns = 10
	}

	if c.Database.MaxOpenConns == 0 {
		c.Database.MaxOpenConns = 100
	}

	// All is good
	return nil
}
