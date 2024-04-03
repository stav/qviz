package middleware

import (
	"fmt"
	"net/http"
	"sort"

	"github.com/labstack/echo/v4"
)

// ServerHeader middleware adds a custom header to the response.
func serverHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Custom-Header", "blah!!!")
		return next(c)
	}
}

func printHeaders(header http.Header) {
	keys := make([]string, 0, len(header))
	for k := range header {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Printf("%-20s: %#v\n", k, header.Get(k))
	}
}

func Sentry(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("\n____________________")
		fmt.Println("Middleware: Headers")
		fmt.Println()

		request := c.Request()
		headers := request.Header

		printHeaders(headers)

		cookie, err := c.Cookie("token")
		if err != nil {
			return c.Redirect(http.StatusSeeOther, "/login")
		}
		// "token=eyJhbGciOiJI---.---" 877 chars
		token := cookie.Value
		fmt.Println("\nToken: ", token)

		// Now perform a security check on the token

		// For invalid credentials
		// return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")

		// For valid credentials call next
		return next(c)
	}
}
