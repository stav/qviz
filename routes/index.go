package routes

import (
	"context"
	"os"

	"github.com/labstack/echo/v4"

	supa "github.com/nedpals/supabase-go"
)

var SUPABASE_URL string = os.Getenv("SUPABASE_URL")
var SUPABASE_KEY string = os.Getenv("SUPABASE_KEY")
var supabase = supa.CreateClient(SUPABASE_URL, SUPABASE_KEY)

type Result struct {
	Error error
	Auth  supa.AuthenticatedDetails
}

func IndexHandler(c echo.Context) error {
	return c.Render(200, "index", nil)
}

func GetLoginHandler(c echo.Context) error {
	return c.Render(200, "login", nil)
}

func PostLoginHandler(c echo.Context) error {
	ctx := context.Background()

	creds := supa.UserCredentials{
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
	}
	auth, err := supabase.Auth.SignIn(ctx, creds)

	status := 200
	if err != nil {
		status = 422
	}

	result := Result{
		Error: err,
	}

	if auth != nil {
		result.Auth = *auth
	}

	return c.Render(status, "result", result)
}
