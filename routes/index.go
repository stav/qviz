package routes

import (
	"os"

	"github.com/labstack/echo/v4"

	supa "github.com/nedpals/supabase-go"

	_ "github.com/joho/godotenv/autoload"
)

var SUPABASE_URL string = os.Getenv("SUPABASE_URL")
var SUPABASE_KEY string = os.Getenv("SUPABASE_KEY")
var supabase = supa.CreateClient(SUPABASE_URL, SUPABASE_KEY)

func IndexHandler(c echo.Context) error {
	return c.Render(200, "index", "Hello, Qviz!")
}
