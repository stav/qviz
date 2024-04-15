package middleware

import (
	echo_middleware "github.com/labstack/echo/v4/middleware"
)

type LoggerConfig = echo_middleware.LoggerConfig

var LoggerWithConfig = echo_middleware.LoggerWithConfig
