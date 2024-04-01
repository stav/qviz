package routes

import (
	"context"
	"fmt"
	"os"

	"github.com/labstack/echo/v4"

	supa "github.com/nedpals/supabase-go"
)

var SUPABASE_URL string = os.Getenv("SUPABASE_URL")
var SUPABASE_KEY string = os.Getenv("SUPABASE_KEY")
var supabase = supa.CreateClient(SUPABASE_URL, SUPABASE_KEY)

type Result struct {
	Error error
	Id    string
	Email string
	Token string
	Message string
}

func IndexHandler(c echo.Context) error {
	return c.Render(200, "index.html", nil)
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
		result.Id = auth.User.ID
		result.Email = auth.User.Email
		result.Token = auth.AccessToken
		result.Message = "User has been authenticated"
	}

	return c.Render(status, "result", result)
}

func GetRegisterHandler(c echo.Context) error {
	return c.Render(200, "register", nil)
}

func PostRegisterHandler(c echo.Context) error {
	ctx := context.Background()

	creds := supa.UserCredentials{
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
	}
	user, err := supabase.Auth.SignUp(ctx, creds)

	status := 200
	if err != nil {
		status = 422
	}

	result := Result{
		Error: err,
	}

	if user != nil {
		result.Id = user.ID
		result.Email = user.Email
		result.Message = fmt.Sprintf("User %s has been created", user.Email)
	}

	return c.Render(status, "result", result)
}
