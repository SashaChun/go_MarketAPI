package config

import "os"

var Port = os.Getenv("PORT")

var Host = os.Getenv("HOST")

var DatabaseURL = os.Getenv("DATABASE_URL")

var User = os.Getenv("DB_USER")

var Password = os.Getenv("PASSWORD")
