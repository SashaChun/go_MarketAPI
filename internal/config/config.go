package config

import "os"

func GetAppPort() string {
	return os.Getenv("APP_PORT")
}

func GetDBPort() string {
	return os.Getenv("DB_PORT")
}

func GetDBHost() string {
	return os.Getenv("DB_HOST")
}

func GetDBName() string {
	return os.Getenv("DB_NAME")
}

func GetDBUser() string {
	return os.Getenv("DB_USER")
}

func GetDBPassword() string {
	return os.Getenv("DB_PASSWORD")
}

func GetJWTSecret() string {
	return os.Getenv("JWT_SECRET")
}
