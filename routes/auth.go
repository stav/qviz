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

type User struct {
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

	user := User{
		Error: err,
	}

	if auth != nil {
		user.Id = auth.User.ID
		user.Email = auth.User.Email
		user.Token = auth.AccessToken
		user.Message = "User has been authenticated"
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

	return c.Render(status, "__user.html", user)
}

func PostLogoutHandler(c echo.Context) error {
	ctx := context.Background()

	cookie, err := c.Cookie("token")
	if err != nil {
		return c.Render(http.StatusAlreadyReported, "__logout.html", "Tried to logout but no token found")
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
	return c.Render(http.StatusAccepted, "__logout.html", "User has been logged out")
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
	auth, err := supabase.Auth.SignUp(ctx, creds)

	status := 200
	if err != nil {
		status = 422
	}

	user := User{
		Error: err,
	}

	if auth != nil {
		user.Id = auth.ID
		user.Email = auth.Email
		user.Message = fmt.Sprintf("User %s has been created", auth.Email)
	}

	return c.Render(status, "__user.html", user)
}
