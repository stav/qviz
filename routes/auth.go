package routes

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	supa "github.com/nedpals/supabase-go"

	_ "github.com/joho/godotenv/autoload"
)

type Result struct {
	Error error
	Id    string
	Email string
	Token string
	Message string
}

func GetLoginHandler(c echo.Context) error {
	return c.Render(200, "user.html", "login")
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
		c.SetCookie(&http.Cookie{
			Path: "/",
			Name: "token",
			Value: auth.AccessToken,
			Expires: time.Now().Add(24 * time.Hour),
			SameSite: http.SameSiteLaxMode,
			HttpOnly: true,
			Secure: true,
		})
	}

	return c.Render(status, "result", result)
}

func PostLogoutHandler(c echo.Context) error {
	ctx := context.Background()

	cookie, err := c.Cookie("token")
	if err != nil {
		return c.Render(http.StatusAlreadyReported, "logout", "Tried to logout but no token found")
	}
	supabase.Auth.SignOut(ctx, cookie.Value)
	c.SetCookie(&http.Cookie{
		Path: "/",
		Name: "token",
		Value: "-",
		MaxAge: -1,
		SameSite: http.SameSiteLaxMode,
		HttpOnly: true,
		Secure: true,
	})
	return c.Render(http.StatusAccepted, "logout", "User has been logged out")
}

func GetRegisterHandler(c echo.Context) error {
	return c.Render(200, "user.html", "register")
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
