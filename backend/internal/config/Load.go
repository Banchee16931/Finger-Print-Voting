package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type RuntimeConfig struct {
	DB             DBConfig
	PasswordSecret string
}

type DBConfig struct {
	Database string
	Username string
	Password string
	Port     string
	Server   string
}

// loads the config from the .env file
func Load() RuntimeConfig {
	codebaseLoc, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fileLoc := fmt.Sprintf("%s\\internal\\config\\.env", codebaseLoc)
	fmt.Println("file location: ", fileLoc)
	file, err := os.Open(fileLoc)
	if err != nil {
		panic(fmt.Sprintf("failed to open config file: %s", err.Error()))
	}

	fileScanner := bufio.NewScanner(file)

	var fileLines []string

	for fileScanner.Scan() {
		fmt.Sprintln("scan line")
		fileLines = append(fileLines, fileScanner.Text())
	}

	file.Close()

	config := RuntimeConfig{}

	fmt.Println("file line count: ", len(fileLines))

	for _, line := range fileLines {
		if len(line) != 0 {
			values := strings.Split(line, "=")
			fmt.Println("value pair: ", values)
			if len(values) == 2 {
				switch values[0] {
				case "DB":
					config.DB.Database = values[1]
				case "DB_SERVER":
					config.DB.Server = values[1]
				case "DB_PORT":
					config.DB.Port = values[1]
				case "DB_USERNAME":
					config.DB.Username = values[1]
				case "DB_PASSWORD":
					config.DB.Password = values[1]
				case "PASS_SECRET":
					config.PasswordSecret = values[1]
				}
			}
		}
	}

	if config.DB.Database == "" || config.DB.Server == "" || config.DB.Port == "" || config.DB.Username == "" || config.DB.Password == "" {
		panic("missing database config")
	}

	if config.PasswordSecret == "" {
		panic("missing password secret")
	}

	return config
}
