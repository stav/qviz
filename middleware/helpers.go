package middleware

import (
	"fmt"
	"net/http"
	"sort"

	"github.com/labstack/echo/v4"
)

// printHeaders is a helper function to print the headers of an HTTP request.
func printHeaders(headers http.Header) {
	keys := make([]string, 0, len(headers))
	for k := range headers {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Printf("%-20s: %#v\n", k, headers.Get(k))
	}
}

// AddScriptHeader adds a custom header to the response.
func AddScriptHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("Setting Content-Type to application/javascript")
		c.Response().Header().Set("Content-Type", "application/javascript")
		return next(c)
	}
}
