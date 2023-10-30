package main

import (
	"context"
	"fmt"
	"my-auth-service/config"
	"my-auth-service/server"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

// Config represents server specific config
type DefaultConfig struct {
	Stage        string
	Port         int
	ReadTimeout  int
	WriteTimeout int
	Debug        bool
	AllowOrigins []string
}

func main() {
	e := prepare()
	run(e)
}

func prepare() *echo.Echo {
	cfg, err := config.Load()
	if err != nil {
		fmt.Println("Error when load config:", err)
	}

	fmt.Println(cfg)

	e := echo.New()
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{LogLevel: log.ERROR}),
		Headers(), CORS(cfg))
	e.GET("/", healthCheck)

	e.Validator = server.NewValidator()
	e.HTTPErrorHandler = server.NewErrorHandler(e).Handle
	e.Binder = server.NewBinder()
	e.Debug = cfg.Debug
	e.Server.Addr = fmt.Sprintf(":%d", cfg.Port)
	e.Server.ReadTimeout = time.Duration(cfg.ReadTimeout) * time.Minute
	e.Server.WriteTimeout = time.Duration(cfg.WriteTimeout) * time.Minute

	return e
}

func healthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
		// "version":    version,
		// "build_time": buildTime,
	})
}

func run(e *echo.Echo) {
	go func() {
		if err := e.StartServer(e.Server); err != nil {
			if err == http.ErrServerClosed {
				e.Logger.Info("shutting down the server")
			} else {
				e.Logger.Errorf("error shutting down the server: ", err)
			}
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

// CORS adds Cross-Origin Resource Sharing support
func CORS(cfg *config.Configuration) echo.MiddlewareFunc {
	allowOrigins := []string{"*"}
	if cfg != nil && cfg.AllowOrigins != nil {
		allowOrigins = cfg.AllowOrigins
	}

	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     allowOrigins,
		AllowMethods:     []string{"POST", "GET", "PUT", "DELETE", "PATCH", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Content-Length"},
		MaxAge:           86400,
	})
}

// Headers adds general security headers for basic security measures
func Headers() echo.MiddlewareFunc {
	return middleware.SecureWithConfig(middleware.SecureConfig{
		XSSProtection:         "1; mode=block",
		ContentTypeNosniff:    "nosniff",
		XFrameOptions:         "DENY",
		HSTSMaxAge:            31536000,
		HSTSExcludeSubdomains: true,
		// ContentSecurityPolicy: "default-src 'self'",
	})
}
