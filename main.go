package main

import (
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Config struct {
	ServerPort string `mapstructure:"SERVER_PORT"`
	JWTSecret  string `mapstructure:"JWT_SECRET"`
	JWTIssuer  string `mapstructure:"JWT_ISSUER"`
	Database   struct {
		Driver string `mapstructure:"DB_DRIVER"`
		DSN    string `mapstructure:"DB_DSN"`
	} `mapstructure:"DATABASE"`
}

var config Config

func initConfig() {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&config); err != nil {
		panic(err)
	}
}

func main() {
	initConfig()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/auth", authHandler)
	e.GET("/me", meHandler, jwtMiddleware)
	e.GET("/manga/:id", getMangaHandler)
	e.GET("/manga", listMangaHandler)
	e.POST("/manga", createMangaHandler)
	e.PUT("/manga/:id", updateMangaHandler)

	e.Logger.Fatal(e.Start(":" + config.ServerPort))
}

func authHandler(c echo.Context) error {
	// Implement user registration and login
	return nil
}

func meHandler(c echo.Context) error {
	// Implement retrieving authenticated user's information
	return nil
}

func getMangaHandler(c echo.Context) error {
	// Implement retrieving a specific manga by its ID
	return nil
}

func listMangaHandler(c echo.Context) error {
	// Implement retrieving a list of manga
	return nil
}

func createMangaHandler(c echo.Context) error {
	// Implement creating a new manga
	return nil
}

func updateMangaHandler(c echo.Context) error {
	// Implement updating an existing manga
	return nil
}

func jwtMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "missing or malformed jwt")
		}

		tokenString := authHeader[len("Bearer "):]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, echo.NewHTTPError(http.StatusUnauthorized, "unexpected signing method")
			}
			return []byte(config.JWTSecret), nil
		})

		if err != nil || !token.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid or expired jwt")
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid or expired jwt")
		}

		c.Set("user", claims["user"])
		return next(c)
	}
}

func initDB() *gorm.DB {
	var db *gorm.DB
	var err error

	switch config.Database.Driver {
	case "mysql":
		db, err = gorm.Open(mysql.Open(config.Database.DSN), &gorm.Config{})
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(config.Database.DSN), &gorm.Config{})
	default:
		panic("unsupported database driver")
	}

	if err != nil {
		panic("failed to connect database")
	}

	return db
}
