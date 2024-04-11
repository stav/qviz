package middleware

import (
	"fmt"
	"net/http"
	"sort"

	"github.com/labstack/echo/v4"
)

// printHeaders is a helper function to print the headers of an HTTP request.
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

// AddScriptHeader adds a custom header to the response.
func AddScriptHeader(next echo.HandlerFunc) echo.HandlerFunc {
	fmt.Println("AddScriptHeader middleware")
	return func(c echo.Context) error {
		fmt.Println("Setting Content-Type to application/javascript")
		c.Response().Header().Set("Content-Type", "application/javascript")
		return next(c)
	}
}

// Sentry performs a security check on the token in the request cookie.
// If the token is valid, it calls the next handler in the chain.
// If the token is invalid or missing, it redirects the user to the login page.
func Sentry(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("\n____________________")
		fmt.Println("Middleware")

		// request := c.Request()
		// headers := request.Header
		// printHeaders(headers)

		cookie, err := c.Cookie("token")
		if err != nil {
			return c.Redirect(http.StatusSeeOther, "/login")
		}
		// "token=eyJhbGciOiJI---.---" 877 chars
		token := cookie.Value
		fmt.Println("Token: ", len(token))

		// Now perform a security check on the token

		// For invalid credentials
		// return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")

		// For valid credentials call next
		return next(c)
	}
}
